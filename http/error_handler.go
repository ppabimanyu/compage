package http

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

func ErrorHandler(dataType ...string) fiber.ErrorHandler {
	handler := Handler{}
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		if code == fiber.StatusInternalServerError {
			slog.Error("ErrorHandler: Failed to handle error", "error", err)
			return handler.InternalServerError(c, "Internal Server Error", err)
		} else {
			return handler.BadRequest(c, err.Error(), err)
		}
	}
}
