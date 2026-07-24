package app

import (
	stdlog "log"
	"time"
)

func Configure() {
	stdlog.SetFlags(stdlog.Ldate | stdlog.Ltime | stdlog.Lmicroseconds | stdlog.LUTC)
}

func Started()              { stdlog.Printf("application: starting") }
func Stopped()              { stdlog.Printf("application: stopped") }
func ConfigurationLoading() { stdlog.Printf("application: loading configuration") }
func ConfigurationFailed(err error) {
	stdlog.Printf("application: configuration failed error=%q", err)
}
func ConfigurationLoaded(grpcAddress, postgresAddress, redisAddress string, kafkaEnabled bool, kafkaBrokers []string) {
	stdlog.Printf(
		"application: configuration loaded grpc_address=%s postgres_address=%s redis_address=%s kafka_enabled=%t kafka_brokers=%v",
		grpcAddress, postgresAddress, redisAddress, kafkaEnabled, kafkaBrokers,
	)
}
func DependencyConnecting(name, address string) {
	stdlog.Printf("application: connecting dependency=%s address=%s", name, address)
}
func DependencyConnected(name, address string) {
	stdlog.Printf("application: dependency connected dependency=%s address=%s", name, address)
}
func DependencyFailed(name, address string, err error) {
	stdlog.Printf("application: dependency connection failed dependency=%s address=%s error=%q", name, address, err)
}
func ComponentsInitializing() {
	stdlog.Printf("application: initializing repositories, caches, services, and gRPC handlers")
}
func ComponentsInitialized() {
	stdlog.Printf("application: repositories, caches, services, and gRPC handlers initialized")
}
func GRPCServerStarting(address string) {
	stdlog.Printf("application: gRPC server starting address=%s", address)
}
func GRPCServerFailed(err error) {
	stdlog.Printf("application: gRPC server failed error=%q", err)
}
func GRPCServerStopped()      { stdlog.Printf("application: gRPC server stopped") }
func ShutdownSignalReceived() { stdlog.Printf("application: shutdown signal received") }
func ShutdownForced(timeout time.Duration) {
	stdlog.Printf("application: graceful shutdown timed out timeout=%s", timeout)
}
