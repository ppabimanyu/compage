package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"os"
)

type Config struct {
	Host         string
	Port         int
	Password     string
	DB           int
	RESPProtocol int
}

func NewConnection(config *Config) *redis.Client {
	if config == nil {
		slog.Error("Redis: config cannot be nil")
		os.Exit(1)
	}
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == 0 {
		config.Port = 6379
	}
	if config.RESPProtocol == 0 {
		config.RESPProtocol = 3
	}
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
		Protocol: config.RESPProtocol,
	})
}

func NewConnectionWithURL(url string) *redis.Client {
	opts, err := redis.ParseURL(url)
	if err != nil {
		slog.Error("Redis: Failed to parse redis url", "error", err.Error())
		os.Exit(1)
	}
	return redis.NewClient(opts)
}
