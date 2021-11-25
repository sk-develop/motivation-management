package adapter

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/value"
	"github.com/sk-develop/motivation-management/services/task/internal/usecase"
	pb "github.com/sk-develop/motivation-management/shared/proto/task"
)

type Controller interface {
	Get(userID string) ([]*pb.Task, error)
	// Post(context.Context) error
	// Patch(context.Context) error
	// Delete(context.Context) error
}

type TaskController struct {
	taskUsecase usecase.TaskUsecase
	taskValue   value.TaskValue
	taskModel   model.TaskModel
}

func NewTaskController(
	taskUsecase usecase.TaskUsecase,
	taskValue value.TaskValue,
	taskModel model.TaskModel,
) Controller {
	return &TaskController{taskUsecase: taskUsecase, taskValue: taskValue, taskModel: taskModel}
}

func (taskController *TaskController) Get(userID string) ([]*pb.Task, error) {
	validUserID, err := taskController.taskValue.NewUserID(userID)
	if err != nil {
		return nil, err
	}

	foundTasks, err := taskController.taskUsecase.ReadAll(*validUserID)
	if err != nil {
		return nil, err
	}

	tasksMessage := taskController.taskModel.TasksToGrpcMessage(foundTasks)

	return tasksMessage, nil
}
