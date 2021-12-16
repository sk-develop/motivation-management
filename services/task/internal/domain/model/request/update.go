package request

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/task"
	pb "github.com/sk-develop/motivation-management/shared/proto/task"
)

type UpdateTaskReq struct {
	id      task.ID
	name    task.Name
	dueDate task.DueDate
}

func NewUpdateTask(pbReq *pb.UpdateTaskReq) (*UpdateTaskReq, error) {
	validID, err := task.NewID(int(pbReq.ID))
	if err != nil {
		return nil, err
	}
	validName, err := task.NewName(pbReq.Name)
	if err != nil {
		return nil, err
	}
	validDueDate, err := task.NewDueDate(pbReq.DueDate.AsTime())
	if err != nil {
		return nil, err
	}

	return &UpdateTaskReq{id: *validID, name: *validName, dueDate: *validDueDate}, nil
}

func (utr UpdateTaskReq) ID() task.ID {
	return utr.id
}

func (utr UpdateTaskReq) Name() task.Name {
	return utr.name
}

func (utr UpdateTaskReq) DueDate() task.DueDate {
	return utr.dueDate
}
