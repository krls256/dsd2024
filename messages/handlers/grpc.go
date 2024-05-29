package handlers

import (
	"context"
	"errors"
	"github.com/krls256/dsd2024/api"
	"github.com/krls256/dsd2024/messages/services"
)

func NewMessagesHandler(messageService *services.MessageService) *MessagesHandler {
	return &MessagesHandler{
		messageService: messageService,
	}
}

type MessagesHandler struct {
	messageService *services.MessageService
}

func (h *MessagesHandler) SendMessage(ctx context.Context, message *api.Message) (*api.MessagesStatusResponse, error) {
	return nil, errors.New("method is not implemented")
}

func (h *MessagesHandler) AllMessages(ctx context.Context, request *api.MessagesZeroRequest) (*api.Messages, error) {
	return &api.Messages{Text: h.messageService.All()}, nil
}
