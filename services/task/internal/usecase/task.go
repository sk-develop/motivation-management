package usecase

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/repository"
)

type TaskUsecase interface {
	ReadAll(req *model.GetTaskReq) (*model.Tasks, error)
	Create(req *model.CreateTaskReq) (*model.Task, error)
}

type taskUsecase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUsecase(taskRepository repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepository: taskRepository}
}

func (tu *taskUsecase) ReadAll(req *model.GetTaskReq) (*model.Tasks, error) {
	return tu.taskRepository.ReadAll(req.UserID)
}

func (tu *taskUsecase) Create(req *model.CreateTaskReq) (*model.Task, error) {
	return tu.taskRepository.Create(req.UserID, req.Name, req.DueDate)
}
