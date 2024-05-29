package main

import (
	"flag"
	"github.com/krls256/dsd2024/api"
	facadeDI "github.com/krls256/dsd2024/facade/di"
	"github.com/krls256/dsd2024/messages/di"
	"github.com/krls256/dsd2024/messages/handlers"
	"github.com/krls256/dsd2024/pkg/consul"
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
		grpc.Config{Host: "127.0.0.1", Port: uint16(*port)},
		h, api.RegisterMessagesServiceServer)

	s.RunAsync()

	cancel, err := consul.Register("0.0.0.0:8500", "messages", "127.0.0.1", uint16(*port))
	if err != nil {
		slog.Error("can't register'", "err", err)
		s.Shutdown()

		return
	}

	<-utils.WaitTermSignal()

	cancel()
	s.Shutdown()

	slog.Info("shutdown", "server was running for", time.Since(now))
}
