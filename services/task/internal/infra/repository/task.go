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

func (taskRepository *TaskRepository) ReadAll(userId value.UserID) (*model.Tasks, error) {
	var foundTask model.Tasks
	if err := taskRepository.conn.Where("user_id = ?", userId).Find(&foundTask).Error; err != nil {
		logger.Warn(err)
		return nil, err
	}

	return &foundTask, nil
}
