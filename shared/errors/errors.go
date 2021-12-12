package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidationError(message string) error {
	return status.Error(codes.InvalidArgument, message)
}
