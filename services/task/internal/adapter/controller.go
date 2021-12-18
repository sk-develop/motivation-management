package adapter

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/request"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/response"
	"github.com/sk-develop/motivation-management/services/task/internal/usecase"
	"github.com/sk-develop/motivation-management/shared/errors"
	pb "github.com/sk-develop/motivation-management/shared/proto/task"
)

type TaskController struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskController(tu usecase.TaskUsecase) *TaskController {
	return &TaskController{taskUsecase: tu}
}

func (tc *TaskController) Get(pbReq *pb.GetTasksReq) (*pb.GetTasksRes, error) {
	req, err := request.NewGetTask(pbReq)
	if err != nil {
		return nil, errors.InvalidArgument(err)
	}

	foundTasks, err := tc.taskUsecase.ReadAll(req)
	if err != nil {
		return nil, errors.Internal(err)
	}

	return response.GetTasksGrpcRes(foundTasks), nil
}

func (tc *TaskController) Create(pbReq *pb.CreateTaskReq) (*pb.CreateTaskRes, error) {
	req, err := request.NewCreateTask(pbReq)
	if err != nil {
		return nil, errors.InvalidArgument(err)
	}

	foundTask, err := tc.taskUsecase.Create(req)
	if err != nil {
		return nil, errors.Internal(err)
	}

	return response.CreateTaskGrpcRes(foundTask), nil
}

func (tc *TaskController) Update(pbReq *pb.UpdateTaskReq) (*pb.UpdateTaskRes, error) {
	req, err := request.NewUpdateTask(pbReq)
	if err != nil {
		return nil, errors.InvalidArgument(err)
	}

	foundTask, err := tc.taskUsecase.Update(req)
	if err != nil {
		return nil, errors.Internal(err)
	}

	return response.UpdateTaskGrpcRes(foundTask), nil
}

func (tc *TaskController) Delete(pbReq *pb.DeleteTasksReq) (*pb.DeleteTasksRes, error) {
	req, err := request.NewDeleteTasks(pbReq)
	if err != nil {
		return nil, errors.InvalidArgument(err)
	}

	if err := tc.taskUsecase.Delete(req); err != nil {
		return nil, errors.Internal(err)
	}

	return response.DeleteTasksGrpcRes(), nil
}

func (tc *TaskController) SwitchCompleted(pbReq *pb.SwitchCompletedReq) (*pb.SwitchCompletedRes, error) {
	req, err := request.NewSwitchCompleted(pbReq)
	if err != nil {
		return nil, errors.InvalidArgument(err)
	}

	foundTask, err := tc.taskUsecase.SwitchCompleted(req)
	if err != nil {
		return nil, errors.Internal(err)
	}

	return response.SwitchCompletedGrpcRes(foundTask), nil
}
