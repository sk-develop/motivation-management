package request

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/task"
	pb "github.com/sk-develop/motivation-management/shared/proto/task"
)

type GetTaskReq struct {
	userID task.UserID
}

func NewGetTask(pbReq *pb.GetTasksReq) (*GetTaskReq, error) {
	validUserID, err := task.NewUserID(pbReq.UserID)
	if err != nil {
		return nil, err
	}

	return &GetTaskReq{userID: *validUserID}, nil
}

func (gtr GetTaskReq) UserID() task.UserID {
	return gtr.userID
}
