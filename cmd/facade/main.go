package main

import (
	"context"
	"flag"
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/krls256/dsd2024/api"
	facadeDI "github.com/krls256/dsd2024/facade/di"
	"github.com/krls256/dsd2024/facade/handlers"
	"github.com/krls256/dsd2024/facade/services"
	"github.com/krls256/dsd2024/pkg/consul"
	pkgDI "github.com/krls256/dsd2024/pkg/di"
	pkgHazelcast "github.com/krls256/dsd2024/pkg/hazelcast"
	"github.com/krls256/dsd2024/pkg/transport/http"
	"github.com/krls256/dsd2024/utils"
	"log/slog"
	"os"
	"time"
)

var port = flag.Uint("p", 1250, "set port")

func init() {
	flag.Parse()
}

func main() {
	now := time.Now()
	defs := facadeDI.Defs()

	ctn, err := pkgDI.Build(defs...)
	if err != nil {
		panic(err)
	}

	hc := ctn.Get(facadeDI.HazelcastClientName).(*hazelcast.Client)
	cfg := ctn.Get(facadeDI.HazelcastConfigName).(pkgHazelcast.Config)

	loggingClient := api.NewLoggingClients("0.0.0.0:8500", "logging")
	messagesClient := api.NewMessageClients("0.0.0.0:8500", "messages")

	facadeService := services.NewFacadeService(loggingClient, messagesClient, hc, cfg)
	facadeHandler := handlers.NewFacadeHandler(facadeService)

	s := http.NewServer(context.Background(), "facade", slog.New(slog.NewTextHandler(os.Stdout, nil)),
		http.Config{
			Host:         "127.0.0.1",
			Port:         uint16(*port),
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
			IdleTimeout:  time.Second * 10,
		},
		[]http.Handler{facadeHandler})

	s.AsyncRun()

	cancel, err := consul.Register("0.0.0.0:8500", "facade", "127.0.0.1", uint16(*port))
	if err != nil {
		slog.Error("can't register'", "err", err)
		s.Shutdown(context.Background())

		return
	}

	<-utils.WaitTermSignal()

	cancel()
	s.Shutdown(context.Background())

	slog.Info("shutdown", "server was running for", time.Since(now))
}
