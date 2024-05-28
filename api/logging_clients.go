package api

import (
	"context"
	"google.golang.org/grpc"
	"math/rand/v2"
)

func NewLoggingClients(loggingClient ...LoggingServiceClient) LoggingClients {
	return LoggingClients{loggingClient: loggingClient}
}

type LoggingClients struct {
	loggingClient []LoggingServiceClient
}

func (l LoggingClients) randClient() LoggingServiceClient {
	return l.loggingClient[rand.N[int](len(l.loggingClient))]
}

func (l LoggingClients) Log(ctx context.Context, in *LoggingMessage, opts ...grpc.CallOption) (*LoggingStatusResponse, error) {
	return l.randClient().Log(ctx, in, opts...)
}

func (l LoggingClients) All(ctx context.Context, in *LoggingZeroRequest, opts ...grpc.CallOption) (*AllText, error) {
	return l.randClient().All(ctx, in, opts...)
}
