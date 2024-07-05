package interceptor

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

func UnaryInterceptor(logger *slog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		start := time.Now()
		logger.Info("request: "+info.FullMethod, info.Server)

		m, err := handler(ctx, req)

		logger.Info("request: "+info.FullMethod, info.Server, slog.Duration("duration", time.Since(start)))

		return m, err
	}
}