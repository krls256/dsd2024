package handlers

import (
	"context"
	"github.com/krls256/dsd2024/api"
)

func NewMessagesHandler() *MessagesHandler {
	return &MessagesHandler{}
}

type MessagesHandler struct {
}

func (h *MessagesHandler) SendMessage(ctx context.Context, message *api.Message) (*api.MessagesStatusResponse, error) {
	return &api.MessagesStatusResponse{Success: true}, nil
}

func (h *MessagesHandler) AllMessages(ctx context.Context, request *api.MessagesZeroRequest) (*api.Messages, error) {
	return &api.Messages{Text: "method is not implemented"}, nil
}
