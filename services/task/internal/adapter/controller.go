package adapter

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/value"
	"github.com/sk-develop/motivation-management/services/task/internal/usecase"
	pb "github.com/sk-develop/motivation-management/shared/proto/task"
)

type TaskController struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskController(tu usecase.TaskUsecase, tv value.TaskValue, tm model.TaskModel) Controller {
	return &TaskController{taskUsecase: tu, taskValue: tv, taskModel: tm}
}

func (tc *TaskController) Get(pbReq *pb.GetTasksReq) ([]*pb.Task, error) {
	req, err := tc.taskModel.NewGetTaskReq(pbReq)
	if err != nil {
		return nil, err
	}

	foundTasks, err := tc.taskUsecase.ReadAll(req)
	if err != nil {
		return nil, err
	}

	res := tc.taskModel.TasksToGrpcMessage(foundTasks)

	return res, nil
}

func (tc *TaskController) Create(pbReq *pb.CreateTaskReq) (*pb.Task, error) {
	req, err := tc.taskModel.NewCreateTaskReq(pbReq)
	if err != nil {
		return nil, err
	}

	foundTasks, err := tc.taskUsecase.Create(req)
	if err != nil {
		return nil, err
	}

	tasksMessage := tc.taskModel.TaskToGrpcMessage(foundTasks)

	return tasksMessage, nil
}
