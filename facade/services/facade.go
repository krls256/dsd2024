package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/krls256/dsd2024/api"
	pkgHazelcast "github.com/krls256/dsd2024/pkg/hazelcast"
	"strings"
)

func NewFacadeService(loggingClient api.LoggingServiceClient, messagesClient api.MessagesServiceClient, hazelcastClient *hazelcast.Client, hazelcastCfg pkgHazelcast.Config) *FacadeService {
	return &FacadeService{
		loggingClient:  loggingClient,
		messagesClient: messagesClient,

		hazelcastClient: hazelcastClient,
		hazelcastCfg:    hazelcastCfg,
	}
}

type FacadeService struct {
	loggingClient  api.LoggingServiceClient
	messagesClient api.MessagesServiceClient

	hazelcastClient *hazelcast.Client
	hazelcastCfg    pkgHazelcast.Config
}

func (s *FacadeService) Info(ctx context.Context) (string, error) {
	logs, err := s.loggingClient.All(ctx, nil)
	if err != nil {
		return "", err
	}

	messages, err := s.messagesClient.AllMessages(ctx, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("[%v]: %v", strings.Join(logs.Text, ","), messages.Text), nil
}

func (s *FacadeService) Message(ctx context.Context, text string) error {
	id := uuid.New()

	logRes, err := s.loggingClient.Log(ctx, &api.LoggingMessage{
		Id:   id[:],
		Text: text,
	})

	if err != nil {
		return fmt.Errorf(logRes.ErrorMessage)
	}

	if err = s.SendMessage(ctx, Message{
		ID:   id,
		Text: text,
	}); err != nil {
		return err
	}

	return nil
}

func (s *FacadeService) SendMessage(ctx context.Context, message Message) error {
	bts, err := json.Marshal(message)
	if err != nil {
		return err
	}

	q, err := s.hazelcastClient.GetQueue(ctx, s.hazelcastCfg.QueueName)
	if err != nil {
		return err
	}

	return q.Put(ctx, string(bts))
}

type Message struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}
