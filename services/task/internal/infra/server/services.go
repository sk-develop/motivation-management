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

// func (ts TaskServer) GetCompletedTasksTotal(
// 	ctx context.Context, request *pb.GetTasksTotalReq,
// ) (
// 	*pb.GetCompletedTasksTotalRes, error,
// ) {
// 	mockTotal := 777777
// 	return &pb.GetCompletedTasksTotalRes{CompletedTasksTotal: uint32(mockTotal)}, nil
// }

// func (ts TaskServer) GetCompletedTasksDailyTotal(
// 	ctx context.Context, request *pb.GetTasksTotalReq,
// ) (
// 	*pb.CompletedTasksDailyTotalRes, error,
// ) {
// 	var totals []*pb.CompletedTasksDailyTotal
// 	return &pb.CompletedTasksDailyTotalRes{CompletedTasksDailyTotals: totals}, nil
// }
