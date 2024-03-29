package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/samber/lo"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Meta    interface{} `json:"meta"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func NewResponse(status int, meta interface{}, data interface{}) *Response {
	success := false
	if status >= 200 && status <= 299 {
		success = true
	}

	response := &Response{
		Status:  status,
		Success: success,
		Meta:    meta,
	}

	if response.Success {
		response.Data = StatusTextIfEmpty(data, status)
	} else {
		response.Errors = StatusTextIfEmpty(data, status)
	}

	return response
}

func NewFormattedError(status int, meta interface{}, err []error) *Response {
	err = lo.Filter(err, func(item error, index int) bool {
		return item != nil
	})

	return NewResponse(status, meta, lo.Map(err, func(item error, index int) string {
		return item.Error()
	}))
}

func StatusTextIfEmpty(data interface{}, status int) interface{} {
	if data == nil {
		return http.StatusText(status)
	}

	return data
}

func OK(ctx fiber.Ctx, meta interface{}, data interface{}) error {
	r := NewResponse(http.StatusOK, meta, data)

	return ctx.Status(r.Status).JSON(r)
}

func BadRequest(ctx fiber.Ctx, meta interface{}, err ...error) error {
	r := NewFormattedError(http.StatusBadRequest, meta, err)

	return ctx.Status(r.Status).JSON(r)
}

func Unauthorized(ctx fiber.Ctx, meta interface{}, err ...error) error {
	r := NewFormattedError(http.StatusUnauthorized, meta, err)

	return ctx.Status(r.Status).JSON(r)
}

func PaymentRequired(ctx fiber.Ctx, meta interface{}, err ...error) error {
	r := NewFormattedError(http.StatusPaymentRequired, meta, err)

	return ctx.Status(r.Status).JSON(r)
}

func Forbidden(ctx fiber.Ctx, meta interface{}, err ...error) error {
	r := NewFormattedError(http.StatusForbidden, meta, err)

	return ctx.Status(r.Status).JSON(r)
}

func NotFound(ctx fiber.Ctx, meta interface{}, err ...error) error {
	r := NewFormattedError(http.StatusNotFound, meta, err)

	return ctx.Status(r.Status).JSON(r)
}

func Conflict(ctx fiber.Ctx, meta interface{}, err ...error) error {
	r := NewFormattedError(http.StatusConflict, meta, err)

	return ctx.Status(r.Status).JSON(r)
}

func Teapot(ctx fiber.Ctx, meta interface{}, err ...error) error {
	r := NewFormattedError(http.StatusTeapot, meta, err)

	return ctx.Status(r.Status).JSON(r)
}

func ValidationFailed(ctx fiber.Ctx, meta interface{}, err ...error) error {
	r := NewFormattedError(http.StatusUnprocessableEntity, meta, err)

	return ctx.Status(r.Status).JSON(r)
}

func TooEarly(ctx fiber.Ctx, meta interface{}, err ...error) error {
	r := NewFormattedError(http.StatusTooEarly, meta, err)

	return ctx.Status(r.Status).JSON(r)
}

func TooManyRequests(ctx fiber.Ctx, meta interface{}, err ...error) error {
	r := NewFormattedError(http.StatusTooManyRequests, meta, err)

	return ctx.Status(r.Status).JSON(r)
}

func ServerError(ctx fiber.Ctx, meta interface{}, err ...error) error {
	r := NewFormattedError(http.StatusInternalServerError, meta, err)

	return ctx.Status(r.Status).JSON(r)
}
