package services

import (
	"context"
	"github.com/krls256/dsd2024/logging/entities"
	"github.com/krls256/dsd2024/logging/repositories"
	"log/slog"
)

func NewLoggingService(repo *repositories.LoggingRepository) *LoggingService {
	return &LoggingService{
		repo: repo,
	}
}

type LoggingService struct {
	repo *repositories.LoggingRepository
}

func (s *LoggingService) Log(ctx context.Context, msg entities.Log) error {
	slog.Info("logging", "text", msg.Text)

	return s.repo.Save(ctx, msg)
}

func (s *LoggingService) AllLog(ctx context.Context) ([]string, error) {
	return s.repo.All(ctx)
}
