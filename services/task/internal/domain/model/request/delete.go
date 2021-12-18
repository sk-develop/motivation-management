package request

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/task"
	pb "github.com/sk-develop/motivation-management/shared/proto/task"
)

type DeleteTasksReq struct {
	ids []task.ID
}

func NewDeleteTasks(pbReq *pb.DeleteTasksReq) (*DeleteTasksReq, error) {
	var validIDs []task.ID

	for _, v := range pbReq.IDs {
		validID, err := task.NewID(int(v))
		if err != nil {
			return nil, err
		}
		validIDs = append(validIDs, *validID)
	}

	return &DeleteTasksReq{ids: validIDs}, nil
}

func (dtr DeleteTasksReq) IDs() []task.ID {
	return dtr.ids
}
