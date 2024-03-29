package main

import (
	"github.com/krls256/dsd2024/api"
	"github.com/krls256/dsd2024/logging/handlers"
	"github.com/krls256/dsd2024/logging/services"
	"github.com/krls256/dsd2024/pkg/transport/grpc"
	"github.com/krls256/dsd2024/utils"
	"log/slog"
	"time"
)

func main() {
	now := time.Now()

	loggingService := services.NewLoggingService()
	h := handlers.NewLoggingHandler(loggingService)

	s := grpc.NewServer[api.LoggingServiceServer](
		grpc.Config{Host: "0.0.0.0", Port: 1234},
		h, api.RegisterLoggingServiceServer)

	s.RunAsync()

	<-utils.WaitTermSignal()

	s.Shutdown()

	slog.Info("shutdown", "server was running for", time.Since(now))
}
