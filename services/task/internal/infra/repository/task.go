package repository

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/repository"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/value"
	"github.com/sk-develop/motivation-management/shared/logger"

	"gorm.io/gorm"
)

type TaskRepository struct {
	conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &TaskRepository{conn: conn}
}

func (tr *TaskRepository) ReadAll(userId value.UserID) (*model.Tasks, error) {
	var foundTask model.Tasks
	if err := tr.conn.Where("user_id = ?", userId).Find(&foundTask).Error; err != nil {
		logger.Warn(err)
		return nil, err
	}

	return &foundTask, nil
}

func (tr *TaskRepository) Create(userID value.UserID, name value.Name, dueDate value.DueDate) (*model.Task, error) {
	createdTask := model.Task{UserID: userID, Name: name, DueDate: dueDate}

	if err := tr.conn.Create(&createdTask).Error; err != nil {
		return nil, err
	}

	return &createdTask, nil
}
