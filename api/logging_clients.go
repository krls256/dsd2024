package api

import (
	"context"
	"github.com/krls256/dsd2024/pkg/consul"
	transportGRPC "github.com/krls256/dsd2024/pkg/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewLoggingClients(consulAddress, serviceType string) LoggingClients {
	return LoggingClients{
		consulAddress: consulAddress,
		serviceType:   serviceType,
	}
}

type LoggingClients struct {
	consulAddress string
	serviceType   string
}

func (l LoggingClients) randClient() (LoggingServiceClient, error) {
	address, err := consul.DiscoverRandom(l.consulAddress, l.serviceType)
	if err != nil {
		return nil, err
	}

	cfg := transportGRPC.Config{Host: address.Host, Port: address.Port}
	conn, err := grpc.Dial(cfg.DNS(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return NewLoggingServiceClient(conn), nil
}

func (l LoggingClients) Log(ctx context.Context, in *LoggingMessage, opts ...grpc.CallOption) (*LoggingStatusResponse, error) {
	c, err := l.randClient()
	if err != nil {
		return nil, err
	}

	return c.Log(ctx, in, opts...)
}

func (l LoggingClients) All(ctx context.Context, in *LoggingZeroRequest, opts ...grpc.CallOption) (*AllText, error) {
	c, err := l.randClient()
	if err != nil {
		return nil, err
	}

	return c.All(ctx, in, opts...)
}
