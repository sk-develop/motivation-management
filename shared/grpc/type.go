package grpc

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type grpcType struct{}

type GrpcType interface {
	Timestamp(t time.Time) *timestamppb.Timestamp
}

func NewGrpcType() GrpcType {
	return &grpcType{}
}

func (grpcType) Timestamp(t time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
