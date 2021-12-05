package model

import (
	"time"

	"github.com/sk-develop/motivation-management/services/task/internal/domain/value"
	pb "github.com/sk-develop/motivation-management/shared/proto/task"
)

type GetTaskReq struct {
	UserID value.UserID
}

type CreateTaskReq struct {
	UserID  value.UserID
	Name    value.Name
	DueDate value.DueDate
}

type UpdateTaskReq struct {
	UserID  value.UserID
	Name    string
	DueDate time.Time
}

type DeleteTaskReq struct {
	UserID value.UserID
	IDs    []value.ID
}

func (tm *taskModel) NewGetTaskReq(pbReq *pb.GetTasksReq) (*GetTaskReq, error) {
	validUserID, err := tm.taskValue.NewUserID(pbReq.UserID)
	if err != nil {
		return nil, err
	}

	return &GetTaskReq{UserID: *validUserID}, nil
}

func (tm *taskModel) NewCreateTaskReq(pbReq *pb.CreateTaskReq) (*CreateTaskReq, error) {
	validUserID, err := tm.taskValue.NewUserID(pbReq.UserID)
	if err != nil {
		return nil, err
	}
	validName, err := tm.taskValue.NewName(pbReq.Name)
	if err != nil {
		return nil, err
	}
	validDueDate, err := tm.taskValue.NewDueDate(pbReq.DueDate.AsTime())
	if err != nil {
		return nil, err
	}

	return &CreateTaskReq{UserID: *validUserID, Name: *validName, DueDate: *validDueDate}, nil
}
