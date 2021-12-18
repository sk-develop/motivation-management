package repository

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/task"
)

type TaskRepository interface {
	ReadAll(userId task.UserID) (*task.Tasks, error)
	Create(userID task.UserID, name task.Name, dueDate task.DueDate) (*task.Task, error)
	Update(userID task.ID, name task.Name, dueDate task.DueDate) (*task.Task, error)
	Delete(ids []task.ID) error
	SwitchCompleted(id task.ID) (*task.Task, error)
}
