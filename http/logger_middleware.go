package http

import (
	"github.com/gofiber/fiber/v2"
	slogfiber "github.com/samber/slog-fiber"
	"log/slog"
)

/*
LoggerMiddleware sets up a logging middleware for Gin using the `slog-gin` package.
It configures logging behavior, including request and response details, trace IDs,
and error levels. The middleware captures and logs the following:
- Request and response bodies (up to 64KB each).
- Request and response headers.
- Trace ID and Request ID for correlation.
- User agent information.
- Error levels for client and server errors.

The middleware is configured with the following options:
- Default log level: Info.
- Client error log level: Info.
- Server error log level: Error.
- Maximum size for request and response bodies: 64KB.

Returns:
- A `gin.HandlerFunc` that can be used as middleware in a Gin router.
*/
func LoggerMiddleware() fiber.Handler {
	slogfiber.TraceIDKey = TraceIDCtxKey
	slogfiber.RequestIDKey = RequestIDCtxKey
	slogfiber.RequestIDHeaderKey = RequestIDHeader
	slogfiber.RequestBodyMaxSize = 64 * 1024  // 64KB
	slogfiber.ResponseBodyMaxSize = 64 * 1024 // 64KB
	return slogfiber.NewWithConfig(slog.Default(), slogfiber.Config{
		DefaultLevel:       slog.LevelInfo,
		ClientErrorLevel:   slog.LevelInfo,
		ServerErrorLevel:   slog.LevelError,
		WithUserAgent:      true,
		WithRequestID:      true,
		WithRequestBody:    true,
		WithRequestHeader:  true,
		WithResponseBody:   true,
		WithResponseHeader: true,
		WithTraceID:        true,
	})
}
