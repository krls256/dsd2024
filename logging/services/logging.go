package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"sync"
)

type Message struct {
	ID   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}

func NewLoggingService() *LoggingService {
	return &LoggingService{
		storage: make(map[uuid.UUID]string),
	}
}

type LoggingService struct {
	storage   map[uuid.UUID]string
	storageMu sync.RWMutex
}

func (s *LoggingService) Log(ctx context.Context, msg Message) error {
	s.storageMu.Lock()
	defer s.storageMu.Unlock()

	s.storage[msg.ID] = msg.Text

	return nil
}

func (s *LoggingService) AllLog(ctx context.Context) []string {
	s.storageMu.RLock()
	defer s.storageMu.RUnlock()

	return lo.Values(s.storage)
}
