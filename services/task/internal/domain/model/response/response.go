package response

import (
	"time"

	"github.com/sk-develop/motivation-management/services/task/internal/domain/model/task"
	"github.com/sk-develop/motivation-management/shared/grpc"
	pb "github.com/sk-develop/motivation-management/shared/proto/task"
)

func GetTasksGrpcRes(tasks *task.Tasks) *pb.GetTasksRes {
	return &pb.GetTasksRes{Tasks: tasksToGrpcMessage(tasks)}
}

func CreateTaskGrpcRes(task *task.Task) *pb.CreateTaskRes {
	return &pb.CreateTaskRes{Task: taskToGrpcMessage(task)}
}

func UpdateTaskGrpcRes(task *task.Task) *pb.UpdateTaskRes {
	return &pb.UpdateTaskRes{Task: taskToGrpcMessage(task)}
}

func tasksToGrpcMessage(tasks *task.Tasks) []*pb.Task {
	var grpcMessage []*pb.Task

	for _, t := range *tasks {
		task := &pb.Task{
			ID:        int64(t.ID),
			UserID:    string(t.UserID),
			Name:      string(t.Name),
			Completed: bool(t.Completed),
			DueDate:   grpc.Timestamp(time.Time(t.DueDate)),
			CreatedAt: grpc.Timestamp(time.Time(t.CreatedAt)),
			UpdatedAt: grpc.Timestamp(time.Time(t.UpdatedAt)),
		}

		grpcMessage = append(grpcMessage, task)
	}

	return grpcMessage
}

func taskToGrpcMessage(t *task.Task) *pb.Task {
	grpcMessage := &pb.Task{
		ID:        int64(t.ID),
		UserID:    string(t.UserID),
		Name:      string(t.Name),
		Completed: bool(t.Completed),
		DueDate:   grpc.Timestamp(time.Time(t.DueDate)),
		CreatedAt: grpc.Timestamp(time.Time(t.CreatedAt)),
		UpdatedAt: grpc.Timestamp(time.Time(t.UpdatedAt)),
	}

	return grpcMessage
}
