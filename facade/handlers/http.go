package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/krls256/dsd2024/facade/services"
	"github.com/krls256/dsd2024/pkg/transport/http"
)

func NewFacadeHandler(facadeService *services.FacadeService) *FacadeHandler {
	return &FacadeHandler{facadeService: facadeService}
}

type FacadeHandler struct {
	facadeService *services.FacadeService
}

func (h *FacadeHandler) Register(router fiber.Router) {
	messages := router.Group("messages")

	messages.Get("", func(ctx fiber.Ctx) error {
		resp, err := h.facadeService.Info(ctx.Context())
		if err != nil {
			return http.BadRequest(ctx, nil, err)
		}

		return http.OK(ctx, nil, resp)
	})

	messages.Post("", func(ctx fiber.Ctx) error {
		err := h.facadeService.Message(ctx.Context(), string(ctx.Body()))
		if err != nil {
			return http.BadRequest(ctx, nil, err)
		}

		return http.OK(ctx, nil, nil)
	})
}
