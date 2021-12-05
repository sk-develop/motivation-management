package main

import (
	"github.com/sk-develop/motivation-management/services/task/internal/adapter"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/model"
	"github.com/sk-develop/motivation-management/services/task/internal/domain/value"
	"github.com/sk-develop/motivation-management/services/task/internal/infra/repository"
	"github.com/sk-develop/motivation-management/services/task/internal/infra/server"
	"github.com/sk-develop/motivation-management/services/task/internal/usecase"
	"github.com/sk-develop/motivation-management/shared/env"
	"github.com/sk-develop/motivation-management/shared/grpc"
	"github.com/sk-develop/motivation-management/shared/logger"
)

func main() {
	env.Load()
	logger := logger.SetUp()
	conn := repository.Conn()

	taskRepository := repository.NewTaskRepository(conn)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskValue := value.NewTaskValue()
	grpcType := grpc.NewGrpcType()
	taskModel := model.NewTaskModel(taskValue, grpcType)
	taskController := adapter.NewTaskController(taskUsecase, taskValue, taskModel)

	grpcServer := server.NewTaskGrpcServer(taskController, logger)
	server.Shutdown(grpcServer)
}
