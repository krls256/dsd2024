package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/krls256/dsd2024/api"
	"strings"
)

func NewFacadeService(loggingClient api.LoggingServiceClient, messagesClient api.MessagesServiceClient) *FacadeService {
	return &FacadeService{
		loggingClient:  loggingClient,
		messagesClient: messagesClient,
	}
}

type FacadeService struct {
	loggingClient  api.LoggingServiceClient
	messagesClient api.MessagesServiceClient
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

	messageRes, err := s.messagesClient.SendMessage(ctx, &api.Message{
		Id:   id[:],
		Text: text,
	})

	if err != nil {
		return fmt.Errorf(messageRes.ErrorMessage)
	}

	return nil
}
