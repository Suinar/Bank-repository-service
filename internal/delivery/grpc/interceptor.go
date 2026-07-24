package grpc

import (
	"context"

	grpclog "github.com/kVinsom/Bank-repository-service/internal/logging/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// LoggingUnaryInterceptor logs every unary request without serializing payloads.
func LoggingUnaryInterceptor(
	ctx context.Context,
	request any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	startedAt := grpclog.RequestStarted(info.FullMethod)
	response, err := handler(ctx, request)
	code := status.Code(err).String()
	if err != nil {
		grpclog.RequestFailed(info.FullMethod, code, startedAt, err)
		return response, err
	}
	grpclog.RequestCompleted(info.FullMethod, code, startedAt)
	return response, nil
}
