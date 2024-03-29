package main

import (
	"github.com/krls256/dsd2024/api"
	"github.com/krls256/dsd2024/messages/handlers"

	"github.com/krls256/dsd2024/pkg/transport/grpc"
	"github.com/krls256/dsd2024/utils"
	"log/slog"
	"time"
)

func main() {
	now := time.Now()

	h := handlers.NewMessagesHandler()

	s := grpc.NewServer[api.MessagesServiceServer](
		grpc.Config{Host: "0.0.0.0", Port: 1235},
		h, api.RegisterMessagesServiceServer)

	s.RunAsync()

	<-utils.WaitTermSignal()

	s.Shutdown()

	slog.Info("shutdown", "server was running for", time.Since(now))
}
