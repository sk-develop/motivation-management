package model

import (
	"time"

	"github.com/sk-develop/motivation-management/services/task/internal/domain/value"
	pb "github.com/sk-develop/motivation-management/shared/proto/task"
	"google.golang.org/protobuf/types/known/timestamppb"
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
}

func (tm *taskModel) TasksToGrpcMessage(tasks *Tasks) []*pb.Task {
	var grpcMessage []*pb.Task

	for _, t := range *tasks {
		task := &pb.Task{
			ID:        int64(t.ID),
			UserID:    string(t.UserID),
			Name:      string(t.Name),
			Completed: bool(t.Completed),
			DueDate:   Timestamp(time.Time(t.DueDate)),
			CreatedAt: Timestamp(time.Time(t.DueDate)),
			UpdatedAt: Timestamp(time.Time(t.DueDate)),
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
		DueDate:   Timestamp(time.Time(t.DueDate)),
		CreatedAt: Timestamp(time.Time(t.DueDate)),
		UpdatedAt: Timestamp(time.Time(t.DueDate)),
	}

	return grpcMessage
}

func Timestamp(t time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

func NewTaskModel(taskValue value.TaskValue) TaskModel {
	return &taskModel{taskValue: taskValue}
}
