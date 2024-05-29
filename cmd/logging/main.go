package main

import (
	"flag"
	"github.com/krls256/dsd2024/api"
	facadeDI "github.com/krls256/dsd2024/facade/di"
	"github.com/krls256/dsd2024/logging/di"
	"github.com/krls256/dsd2024/logging/handlers"
	pkgDI "github.com/krls256/dsd2024/pkg/di"
	"github.com/krls256/dsd2024/pkg/transport/grpc"
	"github.com/krls256/dsd2024/utils"
	"log/slog"
	"time"
)

var port = flag.Uint("p", 1234, "set port")

func init() {
	flag.Parse()
}

func main() {
	now := time.Now()

	defs := di.Defs()
	defs = append(defs, facadeDI.Defs()...)

	ctn, err := pkgDI.Build(defs...)
	if err != nil {
		panic(err)
	}

	h := ctn.Get(di.LoggingHandlerName).(*handlers.LoggingHandler)

	s := grpc.NewServer[api.LoggingServiceServer](
		grpc.Config{Host: "0.0.0.0", Port: uint16(*port)},
		h, api.RegisterLoggingServiceServer)

	s.RunAsync()

	slog.Info("running server", "port", *port)

	<-utils.WaitTermSignal()

	s.Shutdown()

	slog.Info("shutdown", "server was running for", time.Since(now))
}
