package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log/slog"
	"strings"
)

type Config struct {
	Port             int
	CORS             bool
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
}

type Server struct {
	config *Config
	server *fiber.App
}

func NewServer(config *Config) *Server {
	server := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		ErrorHandler:      ErrorHandler(),
	})
	server.Use(recover.New())
	server.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(config.AllowOrigins, ","),
		AllowMethods:     strings.Join(config.AllowMethods, ","),
		AllowHeaders:     strings.Join(config.AllowHeaders, ","),
		AllowCredentials: config.AllowCredentials,
	}))
	server.Use(requestid.New(requestid.Config{
		Header:     "X-Request-Id",
		ContextKey: "request_id",
	}))
	server.Use(ContextMiddleware())
	server.Use(LoggerMiddleware())
	return &Server{
		server: server,
		config: config,
	}
}

func (h *Server) Start() error {
	slog.Info("Server: Starting server", "port", h.config.Port)
	return h.server.Listen(fmt.Sprintf(":%d", h.config.Port))
}

func (h *Server) Shutdown() error {
	slog.Info("Server: Shutting down server")
	return h.server.Shutdown()
}

func (h *Server) App() *fiber.App {
	return h.server
}
