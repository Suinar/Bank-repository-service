package grpc

import (
	stdlog "log"
	"time"
)

func RequestStarted(method string) time.Time {
	stdlog.Printf("gRPC server: request started method=%s", method)
	return time.Now()
}
func RequestCompleted(method, code string, startedAt time.Time) {
	stdlog.Printf(
		"gRPC server: request completed method=%s code=%s duration=%s",
		method, code, time.Since(startedAt),
	)
}
func RequestFailed(method, code string, startedAt time.Time, err error) {
	stdlog.Printf(
		"gRPC server: request failed method=%s code=%s duration=%s error=%q",
		method, code, time.Since(startedAt), err,
	)
}
