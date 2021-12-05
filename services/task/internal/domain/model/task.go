package model

import (
	"time"

	"github.com/sk-develop/motivation-management/services/task/internal/domain/value"
	"github.com/sk-develop/motivation-management/shared/grpc"
	pb "github.com/sk-develop/motivation-management/shared/proto/task"
)

type Task struct {
	ID        value.ID
	UserID    value.UserID
	Name      value.Name
	Completed value.Completed
	DueDate   value.DueDate
	CreatedAt value.CreatedAt
	UpdatedAt value.UpdatedAt
}

type Tasks []Task

type TaskModel interface {
	TasksToGrpcMessage(tasks *Tasks) []*pb.Task
	TaskToGrpcMessage(task *Task) *pb.Task
}

type taskModel struct {
	taskValue value.TaskValue
	grpcType  grpc.GrpcType
}

func (tm *taskModel) TasksToGrpcMessage(tasks *Tasks) []*pb.Task {
	var grpcMessage []*pb.Task

	for _, t := range *tasks {
		task := &pb.Task{
			ID:        int64(t.ID),
			UserID:    string(t.UserID),
			Name:      string(t.Name),
			Completed: bool(t.Completed),
			DueDate:   tm.grpcType.Timestamp(time.Time(t.DueDate)),
			CreatedAt: tm.grpcType.Timestamp(time.Time(t.DueDate)),
			UpdatedAt: tm.grpcType.Timestamp(time.Time(t.DueDate)),
		}

		grpcMessage = append(grpcMessage, task)
	}

	return grpcMessage
}

func (tm *taskModel) TaskToGrpcMessage(t *Task) *pb.Task {
	grpcMessage := &pb.Task{
		ID:        int64(t.ID),
		UserID:    string(t.UserID),
		Name:      string(t.Name),
		Completed: bool(t.Completed),
		DueDate:   tm.grpcType.Timestamp(time.Time(t.DueDate)),
		CreatedAt: tm.grpcType.Timestamp(time.Time(t.DueDate)),
		UpdatedAt: tm.grpcType.Timestamp(time.Time(t.DueDate)),
	}

	return grpcMessage
}

func NewTaskModel(taskValue value.TaskValue, grpcType grpc.GrpcType) TaskModel {
	return &taskModel{taskValue: taskValue, grpcType: grpcType}
}
