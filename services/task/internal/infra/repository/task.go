package repository

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/task"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/repository"

	"gorm.io/gorm"
)

type TaskRepository struct {
	conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &TaskRepository{conn: conn}
}

func (tr *TaskRepository) ReadAll(userId task.UserID) (*task.Tasks, error) {
	var foundTask task.Tasks
	if err := tr.conn.Where("user_id = ?", userId).Find(&foundTask).Error; err != nil {
		return nil, err
	}

	return &foundTask, nil
}

func (tr *TaskRepository) Create(userID task.UserID, name task.Name, dueDate task.DueDate) (*task.Task, error) {
	createdTask := task.New(userID, name, dueDate)

	if err := tr.conn.Create(createdTask).Error; err != nil {
		return nil, err
	}

	return createdTask, nil
}
