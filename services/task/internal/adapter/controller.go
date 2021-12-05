package adapter

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/value"
	"github.com/sk-develop/motivation-management/services/task/internal/usecase"
	pb "github.com/sk-develop/motivation-management/shared/proto/task"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Controller interface {
	Get(userID string) ([]*pb.Task, error)
	Create(userID, name string, dueDates *timestamppb.Timestamp) (*pb.Task, error)
}

type TaskController struct {
	taskUsecase usecase.TaskUsecase
	taskValue   value.TaskValue
	taskModel   model.TaskModel
}

func NewTaskController(tu usecase.TaskUsecase, tv value.TaskValue, tm model.TaskModel) Controller {
	return &TaskController{taskUsecase: tu, taskValue: tv, taskModel: tm}
}

func (tc *TaskController) Get(userID string) ([]*pb.Task, error) {
	validUserID, err := tc.taskValue.NewUserID(userID)
	if err != nil {
		return nil, err
	}

	foundTasks, err := tc.taskUsecase.ReadAll(*validUserID)
	if err != nil {
		return nil, err
	}

	tasksMessage := tc.taskModel.TasksToGrpcMessage(foundTasks)

	return tasksMessage, nil
}

func (tc *TaskController) Create(userID, name string, dueDate *timestamppb.Timestamp) (*pb.Task, error) {
	validUserID, err := tc.taskValue.NewUserID(userID)
	if err != nil {
		return nil, err
	}
	validName, err := tc.taskValue.NewName(name)
	if err != nil {
		return nil, err
	}
	validDueDate, err := tc.taskValue.NewDueDate(dueDate.AsTime())
	if err != nil {
		return nil, err
	}

	foundTasks, err := tc.taskUsecase.Create(validUserID, validName, validDueDate)
	if err != nil {
		return nil, err
	}

	tasksMessage := tc.taskModel.TaskToGrpcMessage(foundTasks)

	return tasksMessage, nil
}
