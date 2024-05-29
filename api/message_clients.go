package api

import (
	"context"
	"github.com/krls256/dsd2024/pkg/consul"
	transportGRPC "github.com/krls256/dsd2024/pkg/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewMessageClients(consulAddress string, serviceType string) MessageClients {
	return MessageClients{
		consulAddress: consulAddress,
		serviceType:   serviceType,
	}
}

type MessageClients struct {
	consulAddress string
	serviceType   string
}

func (m MessageClients) randClient() (MessagesServiceClient, error) {
	address, err := consul.DiscoverRandom(m.consulAddress, m.serviceType)
	if err != nil {
		return nil, err
	}

	cfg := transportGRPC.Config{Host: address.Host, Port: address.Port}
	conn, err := grpc.Dial(cfg.DNS(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return NewMessagesServiceClient(conn), nil
}

func (m MessageClients) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessagesStatusResponse, error) {
	c, err := m.randClient()
	if err != nil {
		return nil, err
	}

	return c.SendMessage(ctx, in, opts...)
}

func (m MessageClients) AllMessages(ctx context.Context, in *MessagesZeroRequest, opts ...grpc.CallOption) (*Messages, error) {
	c, err := m.randClient()
	if err != nil {
		return nil, err
	}

	return c.AllMessages(ctx, in, opts...)
}
