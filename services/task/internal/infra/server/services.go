package server

import (
	"context"

	pb "github.com/sk-develop/motivation-management/shared/proto/task"
)

func (ts TaskServer) GetTasks(ctx context.Context, req *pb.GetTasksReq) (*pb.GetTasksRes, error) {
	return ts.taskController.Get(req)
}

func (ts TaskServer) CreateTask(ctx context.Context, req *pb.CreateTaskReq) (*pb.CreateTaskRes, error) {
	return ts.taskController.Create(req)
}
