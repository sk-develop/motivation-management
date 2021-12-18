package usecase

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/request"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/task"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/repository"
)

type TaskUsecase interface {
	ReadAll(req *request.GetTaskReq) (*task.Tasks, error)
	Create(req *request.CreateTaskReq) (*task.Task, error)
	Update(req *request.UpdateTaskReq) (*task.Task, error)
	Delete(req *request.DeleteTasksReq) error
}

type taskUsecase struct {
	taskRepository repository.TaskRepository
}

func NewTaskUsecase(taskRepository repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepository: taskRepository}
}

func (tu *taskUsecase) ReadAll(req *request.GetTaskReq) (*task.Tasks, error) {
	return tu.taskRepository.ReadAll(req.UserID())
}

func (tu *taskUsecase) Create(req *request.CreateTaskReq) (*task.Task, error) {
	return tu.taskRepository.Create(req.UserID(), req.Name(), req.DueDate())
}

func (tu *taskUsecase) Update(req *request.UpdateTaskReq) (*task.Task, error) {
	return tu.taskRepository.Update(req.ID(), req.Name(), req.DueDate())
}

func (tu *taskUsecase) Delete(req *request.DeleteTasksReq) error {
	return tu.taskRepository.Delete(req.IDs())
}
