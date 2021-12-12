package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func InvalidArgument(message error) error {
	return status.Error(codes.InvalidArgument, message.Error())
}

func Internal(message error) error {
	return status.Error(codes.Internal, message.Error())
}
