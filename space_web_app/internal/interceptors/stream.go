package interceptors

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

type wrappedStream struct {
	grpc.ClientStream
	log *slog.Logger
}

func (w *wrappedStream) RecvMsg(m any) error {
	info := fmt.Sprintf("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	w.log.Info(info)
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m any) error {
	info := fmt.Sprintf("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	w.log.Info(info)
	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ClientStream, log *slog.Logger) grpc.ClientStream {
	return &wrappedStream{s, log}
}

func StreamLoggingInterceptor(log *slog.Logger) grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		start := time.Now()
		s, err := streamer(ctx, desc, cc, method, opts...)
		log.Info(fmt.Sprintf("sent request to %s", method), "method", method, "duration", time.Since(start))
		return newWrappedStream(s, log), err
	}
}
