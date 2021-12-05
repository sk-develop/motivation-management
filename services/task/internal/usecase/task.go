package usecase

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/repository"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/value"
)

type TaskUsecase interface {
	ReadAll(userId value.UserID) (*model.Tasks, error)
	Create(userID *value.UserID, name *value.Name, dueDate *value.DueDate) (*model.Task, error)
}

type taskUsecase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUsecase(taskRepository repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepository: taskRepository}
}

func (tu *taskUsecase) ReadAll(userId value.UserID) (*model.Tasks, error) {
	return tu.taskRepository.ReadAll(userId)
}

func (tu *taskUsecase) Create(userID *value.UserID, name *value.Name, dueDate *value.DueDate) (*model.Task, error) {
	return tu.taskRepository.Create(userID, name, dueDate)
}
