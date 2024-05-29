package handlers

import (
	"context"
	"github.com/google/uuid"
	"github.com/krls256/dsd2024/api"
	"github.com/krls256/dsd2024/logging/entities"
	"github.com/krls256/dsd2024/logging/services"
)

func NewLoggingHandler(loggingService *services.LoggingService) *LoggingHandler {
	return &LoggingHandler{loggingService: loggingService}
}

type LoggingHandler struct {
	loggingService *services.LoggingService
}

func (h *LoggingHandler) Log(ctx context.Context, message *api.LoggingMessage) (*api.LoggingStatusResponse, error) {
	id, err := uuid.FromBytes(message.Id)
	if err != nil {
		return &api.LoggingStatusResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}

	err = h.loggingService.Log(ctx, entities.Log{
		ID:   id,
		Text: message.Text,
	})

	if err != nil {
		return &api.LoggingStatusResponse{
			Success:      false,
			ErrorMessage: err.Error(),
		}, err
	}

	return &api.LoggingStatusResponse{
		Success:      true,
		ErrorMessage: "",
	}, nil
}

func (h *LoggingHandler) All(ctx context.Context, request *api.LoggingZeroRequest) (*api.AllText, error) {
	logs, err := h.loggingService.AllLog(ctx)

	if err != nil {
		return nil, err
	}

	return &api.AllText{
		Text: logs,
	}, nil
}
