package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/*
ContextMiddleware is a Gin middleware function that sets various context values
for each incoming HTTP request. These values include request ID, trace ID, tenant ID,
host, request IP, and language. It ensures that these values are available in the
request context for downstream handlers.

The middleware performs the following steps:
1. Retrieves the `X-Request-Id` header. If not present, generates a new UUID.
2. Retrieves the `X-Trace-Id` header. If not present, uses the request ID as the trace ID.
3. Retrieves the `X-Tenant-Id` header and sets it if present.
4. Sets the host, client IP, and `Accept-Language` header values in the context.
5. Proceeds to the next middleware or handler in the chain.
*/
func ContextMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get(RequestIDHeader)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Locals(RequestIDCtxKey, requestID)

		traceID := c.Get(TraceIDHeader)
		if traceID == "" {
			traceID = requestID
		}
		c.Locals(TraceIDCtxKey, traceID)

		tenantID := c.Get(TenantIDHeader)
		if tenantID != "" {
			c.Locals(TenantIDCtxKey, tenantID)
		}

		c.Locals(HostCtxKey, c.Hostname())
		c.Locals(RequestIPCtxKey, c.IP())
		c.Locals(LangCtxKey, c.Get(LangHeader))

		return c.Next()
	}
}
