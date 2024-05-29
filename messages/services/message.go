package services

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/krls256/dsd2024/messages/entities"
	"github.com/samber/lo"
	"log/slog"
)

const TopicName = "messages"

func NewMessageService(hazelcastClient *hazelcast.Client) (*MessageService, error) {
	s := &MessageService{
		hazelcastClient: hazelcastClient,
		messages:        map[uuid.UUID]string{},
	}

	return s, s.run()
}

type MessageService struct {
	hazelcastClient *hazelcast.Client
	messages        map[uuid.UUID]string
}

func (s *MessageService) All() []string {
	return lo.Values(s.messages)
}

func (s *MessageService) run() error {
	q, err := s.hazelcastClient.GetQueue(context.Background(), TopicName)
	if err != nil {
		return err
	}

	go func() {
		for {
			data, err := q.Poll(context.Background())
			if err != nil {
				slog.Error("cant poll queue", "err", err)

				continue
			}

			if data == nil {
				continue
			}

			str, ok := data.(string)
			if !ok {
				continue
			}

			var message entities.Message

			if err := json.Unmarshal([]byte(str), &message); err != nil {
				slog.Error("cant unmarshal message", "err", err)
				continue
			}

			slog.Info("logging", "text", message.Text)

			s.messages[message.ID] = message.Text
		}
	}()

	return nil
}
