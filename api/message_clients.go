package api

import (
	"context"
	"google.golang.org/grpc"
	"math/rand/v2"
)

func NewMessageClients(messageClient ...MessagesServiceClient) MessageClients {
	return MessageClients{messageClient: messageClient}
}

type MessageClients struct {
	messageClient []MessagesServiceClient
}

func (m MessageClients) SendMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*MessagesStatusResponse, error) {
	return m.randClient().SendMessage(ctx, in, opts...)
}

func (m MessageClients) AllMessages(ctx context.Context, in *MessagesZeroRequest, opts ...grpc.CallOption) (*Messages, error) {
	return m.randClient().AllMessages(ctx, in, opts...)
}

func (m MessageClients) randClient() MessagesServiceClient {
	return m.messageClient[rand.N[int](len(m.messageClient))]
}
