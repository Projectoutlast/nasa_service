package interceptor

import (
	"fmt"
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

type wrappedStream struct {
	grpc.ServerStream
	log *slog.Logger
}

func (w *wrappedStream) RecvMsg(m any) error {
	info := fmt.Sprintf("Receive a message (Type: %T) at %s", m, time.Now().Format(time.RFC3339))
	w.log.Info(info)
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m any) error {
	info := fmt.Sprintf("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	w.log.Info(info)
	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream, log *slog.Logger) grpc.ServerStream {
	return &wrappedStream{s, log}
}

func NewStreamInterceptor(log *slog.Logger) grpc.StreamServerInterceptor {
	return func(
		srv any,
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		err := handler(srv, newWrappedStream(ss, log))

		return err
	}
}
