package main

import (
	"context"
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/krls256/dsd2024/api"
	facadeDI "github.com/krls256/dsd2024/facade/di"
	"github.com/krls256/dsd2024/facade/handlers"
	"github.com/krls256/dsd2024/facade/services"
	pkgDI "github.com/krls256/dsd2024/pkg/di"
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
	defs := facadeDI.Defs()

	ctn, err := pkgDI.Build(defs...)
	if err != nil {
		panic(err)
	}

	hc := ctn.Get(facadeDI.HazelcastClientName).(*hazelcast.Client)

	loggingClient, err := NewLoggingClients()
	if err != nil {
		slog.Error(err.Error())

		return
	}

	messagesClient, err := NewMessageClients()
	if err != nil {
		slog.Error(err.Error())

		return
	}

	facadeService := services.NewFacadeService(loggingClient, messagesClient, hc)
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

func NewMessageClients() (api.MessageClients, error) {
	ports := []uint16{1240, 1241}
	clients := []api.MessagesServiceClient{}

	for _, port := range ports {
		loggingCfg := transportGRPC.Config{Host: "0.0.0.0", Port: port}

		messageConn, err := grpc.Dial(loggingCfg.DNS(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return api.MessageClients{}, err
		}

		clients = append(clients, api.NewMessagesServiceClient(messageConn))
	}

	return api.NewMessageClients(clients...), nil
}
