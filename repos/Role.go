package repos

import (
	"bugtracker/models"
	"bugtracker/structs"
	"errors"
	"time"

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
func (r *roleRepo) FindUsersByRoleName(name string) ([]structs.UserWithRoles, error) {
	users := []models.User{}

	if err := r.storage.Preload("Roles").Find(&users).Where("role_name =?", name).Error; err != nil {
		return nil, err
	}
	//Convert []models.User to []UserWithRoles
	usersWithRoles := make([]structs.UserWithRoles, len(users))
	for i, user := range users {
		// Convert models.User to structs.User
		usersWithRoles[i].User = structs.User{
			Id:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Active:    user.Active,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}

		usersWithRoles[i].Roles = make([]structs.Role, len(user.Roles))
		for j, role := range user.Roles {
			usersWithRoles[i].Roles[j].Id = role.ID
			usersWithRoles[i].Roles[j].Name = role.Name
		}
	}
	return usersWithRoles, nil
}

// unassign user from role
func (r *roleRepo) UnassignUserFromRole(userId uint, roleId uint) error {

	user := models.User{}
	if err := r.storage.Find(&user).Where("id =?", userId).Error; err != nil {
		return err
	}
	//remove role from roles array in user model
	for i, role := range user.Roles {
		if role.ID == roleId {
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

func (r *roleRepo) GetByName(name string) models.Role {
	role := models.Role{}
	r.storage.Find(&role).Where("name =?", name)
	return role
}

func (r *roleRepo) MigrateInitialRoles() error {
	if r.storage.Migrator().HasTable(&models.Role{}) {
		if err := r.storage.First(&models.Role{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			roles := []models.Role{
				{
					Name:      "admin",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					Name:      "viewer",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					Name:      "project_owner",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					Name:      "contributor",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			}
			for _, role := range roles {
				if err := r.Create(role); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
