package request

import (
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/task"
	pb "github.com/sk-develop/motivation-management/shared/proto/task"
)

type CreateTaskReq struct {
	userID  task.UserID
	name    task.Name
	dueDate task.DueDate
}

func NewCreateTask(pbReq *pb.CreateTaskReq) (*CreateTaskReq, error) {
	validUserID, err := task.NewUserID(pbReq.UserID)
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

	return &CreateTaskReq{userID: *validUserID, name: *validName, dueDate: *validDueDate}, nil
}

func (ctr CreateTaskReq) UserID() task.UserID {
	return ctr.userID
}

func (ctr CreateTaskReq) Name() task.Name {
	return ctr.name
}

func (ctr CreateTaskReq) DueDate() task.DueDate {
	return ctr.dueDate
}
