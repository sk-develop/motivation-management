package repository

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/value"
)

type TaskRepository interface {
	ReadAll(userId value.UserID) (*model.Tasks, error)
	Create(userID value.UserID, name value.Name, dueDate value.DueDate) (*model.Task, error)
}
