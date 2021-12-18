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

func (ts TaskServer) UpdateTask(ctx context.Context, req *pb.UpdateTaskReq) (*pb.UpdateTaskRes, error) {
	return ts.taskController.Update(req)
}

func (ts TaskServer) DeleteTasks(ctx context.Context, req *pb.DeleteTasksReq) (*pb.DeleteTasksRes, error) {
	return ts.taskController.Delete(req)
}

func (ts TaskServer) SwitchCompleted(ctx context.Context, req *pb.SwitchCompletedReq) (*pb.SwitchCompletedRes, error) {
	return ts.taskController.SwitchCompleted(req)
}
