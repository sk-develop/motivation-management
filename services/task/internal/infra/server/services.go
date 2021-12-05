package server

import (
	"context"

	pb "github.com/sk-develop/motivation-management/shared/proto/task"
)

func (ts TaskServer) GetTasks(ctx context.Context, req *pb.GetTasksReq) (*pb.GetTasksRes, error) {
	tasks, err := ts.taskController.Get(req.UserID)
	if err != nil {
		return nil, err
	}

	return &pb.GetTasksRes{Tasks: tasks}, nil
}

func (ts TaskServer) CreateTask(ctx context.Context, req *pb.CreateTaskReq) (*pb.CreateTaskRes, error) {
	task, err := ts.taskController.Create(req.UserID, req.Name, req.DueDate)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTaskRes{Task: task}, nil
}
