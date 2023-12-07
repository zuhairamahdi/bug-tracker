package repos

import (
	"bugtracker/models"

	"gorm.io/gorm"
)

type commentRepo struct {
	storage *gorm.DB
}

func newCommentRepo(storage *gorm.DB) *commentRepo {
	return &commentRepo{
		storage: storage,
	}
}

func (r *commentRepo) GetAll(taskId uint) ([]models.Comment, error) {
	comments := []models.Comment{}
	if err := r.storage.Find(&comments).Where("task_id =?", taskId).Error; err != nil {
		return comments, err
	}
	return comments, nil
}

func (r *commentRepo) GetById(commentId uint) (models.Comment, error) {
	comment := models.Comment{}
	if err := r.storage.Find(&comment).Where("id =?", commentId).Error; err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *commentRepo) Create(comment models.Comment) error {
	if err := r.storage.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (r *commentRepo) Update(comment models.Comment) error {
	if err := r.storage.Save(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (r *commentRepo) Delete(commentId uint) error {
	comment := models.Comment{}
	if err := r.storage.Find(&comment).Where("id =?", commentId).Error; err != nil {
		return err
	}
	if err := r.storage.Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (r *commentRepo) GetUserComments(commentId uint, userId string) ([]models.Comment, error) {
	comments := []models.Comment{}
	if err := r.storage.Find(&comments).Where("task_id =? and user_id =?", commentId, userId).Error; err != nil {
		return comments, err
	}
	return comments, nil
}
