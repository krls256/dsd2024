package main

import (
	"context"
	"github.com/krls256/dsd2024/api"
	"github.com/krls256/dsd2024/facade/handlers"
	"github.com/krls256/dsd2024/facade/services"
	transportGRPC "github.com/krls256/dsd2024/pkg/transport/grpc"
	"github.com/krls256/dsd2024/pkg/transport/http"
	"github.com/krls256/dsd2024/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"os"
	"time"
)

func main() {
	now := time.Now()

	messagesCfg := transportGRPC.Config{Host: "0.0.0.0", Port: 1235}

	messagesConn, err := grpc.Dial(messagesCfg.DNS(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error(err.Error())

		return
	}

	loggingClient, err := NewLoggingClients()
	if err != nil {
		slog.Error(err.Error())

		return
	}

	messagesClient := api.NewMessagesServiceClient(messagesConn)

	facadeService := services.NewFacadeService(loggingClient, messagesClient)
	facadeHandler := handlers.NewFacadeHandler(facadeService)

	s := http.NewServer(context.Background(), "facade", slog.New(slog.NewTextHandler(os.Stdout, nil)),
		http.Config{
			Host:         "0.0.0.0",
			Port:         1233,
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
			IdleTimeout:  time.Second * 10,
		},
		[]http.Handler{facadeHandler})

	s.AsyncRun()

	<-utils.WaitTermSignal()

	s.Shutdown(context.Background())

	slog.Info("shutdown", "server was running for", time.Since(now))
}

func NewLoggingClients() (api.LoggingClients, error) {
	ports := []uint16{1230, 1231, 1232}
	clients := []api.LoggingServiceClient{}

	for _, port := range ports {
		loggingCfg := transportGRPC.Config{Host: "0.0.0.0", Port: port}

		loggingConn, err := grpc.Dial(loggingCfg.DNS(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return api.LoggingClients{}, err
		}

		clients = append(clients, api.NewLoggingServiceClient(loggingConn))
	}

	return api.NewLoggingClients(clients...), nil
}
