package usecase

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/repository"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/value"
)

type TaskUsecase interface {
	ReadAll(userId value.UserID) (*model.Tasks, error)
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
