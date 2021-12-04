package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/sk-develop/motivation-management/services/task/internal/adapter"
	"github.com/sk-develop/motivation-management/shared/logger"
	pb "github.com/sk-develop/motivation-management/shared/proto/task"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type TaskServer struct {
	pb.UnimplementedTaskServiceServer
	taskController adapter.Controller
}

func newTaskServer(taskController adapter.Controller) *TaskServer {
	return &TaskServer{taskController: taskController}
}

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

func NewTaskGrpcServer(taskController adapter.Controller, zapLogger *zap.Logger) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Minute,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Minute,
			Time:              gRPCTime * time.Minute,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zapLogger),
		),
		),
	)

	taskServer := newTaskServer(taskController)
	pb.RegisterTaskServiceServer(grpcServer, taskServer)
	grpc_prometheus.Register(grpcServer)

	if os.Getenv("ENV") == "DEV" {
		logger.Info("Development environment")
		reflection.Register(grpcServer)
	} else {
		logger.Info("production environment")
	}

	lis, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		panic(err)
	}

	go func() {
		logger.Infof("Task gRPC server is listening on port: %s", os.Getenv("GRPC_PORT"))
		log.Fatal(grpcServer.Serve(lis))
	}()

	return grpcServer
}

func Shutdown(server *grpc.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	sig := <-c

	logger.Infof("Got signal:", sig)
	logger.Infof("stopping grpc Server...")

	server.GracefulStop()
}
