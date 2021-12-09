package adapter

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/request"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/response"
	"github.com/sk-develop/motivation-management/services/task/internal/usecase"
	pb "github.com/sk-develop/motivation-management/shared/proto/task"
)

type TaskController struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskController(tu usecase.TaskUsecase) *TaskController {
	return &TaskController{taskUsecase: tu}
}

func (tc *TaskController) Get(pbReq *pb.GetTasksReq) (*pb.GetTasksRes, error) {
	req, err := request.NewGetTaskReq(pbReq)
	if err != nil {
		return nil, err
	}

	foundTasks, err := tc.taskUsecase.ReadAll(req)
	if err != nil {
		return nil, err
	}

	return response.GetTasksGrpcRes(foundTasks), nil
}

func (tc *TaskController) Create(pbReq *pb.CreateTaskReq) (*pb.CreateTaskRes, error) {
	req, err := request.NewCreateTask(pbReq)
	if err != nil {
		return nil, err
	}

	foundTasks, err := tc.taskUsecase.Create(req)
	if err != nil {
		return nil, err
	}

	return response.CreateTasksGrpcRes(foundTasks), nil
}
