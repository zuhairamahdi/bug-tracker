package repos

import (
	"bugtracker/models"

	"gorm.io/gorm"
)

type taskRepo struct {
	storage *gorm.DB
}

func newTaskRepo(storage *gorm.DB) *taskRepo {
	return &taskRepo{
		storage: storage,
	}
}

func (r *taskRepo) GetAll(columnId string) ([]models.Task, error) {
	tasks := []models.Task{}
	if err := r.storage.Find(&tasks).Where("column_id =?", columnId).Error; err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *taskRepo) GetAllUserTasks(userId string) ([]models.Task, error) {
	tasks := []models.Task{}
	if err := r.storage.Find(&tasks).Where("user_id =?", userId).Error; err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *taskRepo) GetAllBoardTasks(boardId string) ([]models.Task, error) {
	tasks := []models.Task{}
	if err := r.storage.Find(&tasks).Where("board_id =?", boardId).Error; err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (r *taskRepo) UpdateTitle(taskId uint, title string) error {
	taskToUpdate := models.Task{}
	if err := r.storage.Find(&taskToUpdate).Where("id =?", taskId).Error; err != nil {
		return err
	}
	taskToUpdate.Title = title
	if err := r.storage.Save(&taskToUpdate).Error; err != nil {
		return err
	}
	return nil
}

func (r *taskRepo) UpdateContent(taskId uint, content string) error {
	taskToUpdate := models.Task{}
	if err := r.storage.Find(&taskToUpdate).Where("id =?", taskId).Error; err != nil {
		return err
	}
	taskToUpdate.Content = content
	if err := r.storage.Save(&taskToUpdate).Error; err != nil {
		return err
	}
	return nil
}

func (r *taskRepo) UpdateStatus(taskId uint, status string) error {
	taskToUpdate := models.Task{}
	if err := r.storage.Find(&taskToUpdate).Where("id =?", taskId).Error; err != nil {
		return err
	}
	taskToUpdate.Status = status
	if err := r.storage.Save(&taskToUpdate).Error; err != nil {
		return err
	}
	return nil
}

func (r *taskRepo) Delete(taskId uint) error {
	task := models.Task{}
	if err := r.storage.Find(&task).Where("id =?", taskId).Error; err != nil {
		return err
	}
	if err := r.storage.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}

func (r *taskRepo) GetById(taskId uint) (models.Task, error) {
	task := models.Task{}
	if err := r.storage.Find(&task).Where("id =?", taskId).Error; err != nil {
		return task, err
	}
	return task, nil
}

func (r *taskRepo) Create(task models.Task) error {
	if err := r.storage.Create(&task).Error; err != nil {
		return err
	}
	return nil
}

func (r *taskRepo) UpdateUserToTask(taskId uint, userId uint) error {
	taskToUpdate := models.Task{}
	if err := r.storage.Find(&taskToUpdate).Where("id =?", taskId).Error; err != nil {
		return err
	}
	taskToUpdate.UserID = userId
	if err := r.storage.Save(&taskToUpdate).Error; err != nil {
		return err
	}
	return nil
}
