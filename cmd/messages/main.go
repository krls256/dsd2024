package main

import (
	"flag"
	"github.com/krls256/dsd2024/api"
	facadeDI "github.com/krls256/dsd2024/facade/di"
	"github.com/krls256/dsd2024/messages/di"
	"github.com/krls256/dsd2024/messages/handlers"
	pkgDI "github.com/krls256/dsd2024/pkg/di"

	"github.com/krls256/dsd2024/pkg/transport/grpc"
	"github.com/krls256/dsd2024/utils"
	"log/slog"
	"time"
)

var port = flag.Uint("p", 1240, "set port")

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

	h := ctn.Get(di.MessagesHandlerName).(*handlers.MessagesHandler)

	s := grpc.NewServer[api.MessagesServiceServer](
		grpc.Config{Host: "0.0.0.0", Port: uint16(*port)},
		h, api.RegisterMessagesServiceServer)

	s.RunAsync()

	slog.Info("running server", "port", *port)

	<-utils.WaitTermSignal()

	s.Shutdown()

	slog.Info("shutdown", "server was running for", time.Since(now))
}
