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
