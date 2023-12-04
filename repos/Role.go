package repos

import (
	"bugtracker/models"

	"gorm.io/gorm"
)

type roleRepo struct {
	storage *gorm.DB
}

func newRoleRepo(storage *gorm.DB) *roleRepo {
	return &roleRepo{
		storage: storage,
	}
}

// Find role by name
func (r *roleRepo) FindByName(name string) (models.Role, error) {
	role := models.Role{}
	if err := r.storage.Find(&role).Where("name =?", name).Error; err != nil {
		return role, err
	}
	return role, nil
}

// Find role by id
func (r *roleRepo) FindById(id uint) (models.Role, error) {
	role := models.Role{}
	if err := r.storage.Find(&role).Where("id =?", id).Error; err != nil {
		return role, err
	}
	return role, nil
}

// Update role
func (r *roleRepo) Update(role models.Role) error {
	return r.storage.Save(&role).Error
}

// Update role name by id
func (r *roleRepo) UpdateNameById(id uint, name string) error {
	role := models.Role{}
	if err := r.storage.Find(&role).Where("id =?", id).Error; err != nil {
		return err
	}
	role.Name = name
	return r.storage.Save(&role).Error
}

// Delete role by id
func (r *roleRepo) DeleteById(id uint) error {
	role := models.Role{}
	if err := r.storage.Find(&role).Where("id =?", id).Error; err != nil {
		return err
	}
	return r.storage.Delete(&role).Error
}

// Find users by role id
func (r *roleRepo) FindUsersByRoleId(id uint) ([]models.User, error) {
	users := []models.User{}
	if err := r.storage.Find(&users).Where("role_id =?", id).Error; err != nil {
		return users, err
	}
	return users, nil
}

// Find users by role name
func (r *roleRepo) FindUsersByRoleName(name string) ([]models.User, error) {
	users := []models.User{}
	if err := r.storage.Find(&users).Where("role_name =?", name).Error; err != nil {
		return users, err
	}
	return users, nil
}

// unassign user from role
func (r *roleRepo) UnassignUserFromRole(userId uint, roleId uint) error {

	user := models.User{}
	if err := r.storage.Find(&user).Where("id =?", userId).Error; err != nil {
		return err
	}
	//remove role from roles array in user model
	for i, role := range user.Roles {
		if role.Id == roleId {
			user.Roles = append(user.Roles[:i], user.Roles[i+1:]...)
			break
		}
	}
	return r.storage.Save(&user).Error

}

// Assign user to role
func (r *roleRepo) AssignUserToRole(userId uint, roleId uint) error {
	user := models.User{}
	if err := r.storage.Find(&user).Where("id =?", userId).Error; err != nil {
		return err
	}
	role := models.Role{}
	if err := r.storage.Find(&role).Where("id =?", roleId).Error; err != nil {
		return err
	}
	user.Roles = append(user.Roles, role)
	return r.storage.Save(&user).Error
}

// Get all roles
func (r *roleRepo) GetAll() ([]models.Role, error) {
	roles := []models.Role{}
	if err := r.storage.Find(&roles).Error; err != nil {
		return roles, err
	}
	return roles, nil
}

// Get role by id
func (r *roleRepo) GetById(id uint) (models.Role, error) {
	role := models.Role{}
	if err := r.storage.Find(&role).Where("id =?", id).Error; err != nil {
		return role, err
	}
	return role, nil
}

func (r *roleRepo) Create(role models.Role) error {
	if err := r.storage.Create(&role).Error; err != nil {
		return err
	}
	return nil
}
