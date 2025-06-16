package logger

import (
	"context"
	"github.com/ppabimanyu/compage/logger/prettyslog"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
)

type LogLevel string

const (
	DebugLevel LogLevel = "DEBUG"
	InfoLevel  LogLevel = "INFO"
	WarnLevel  LogLevel = "WARN"
	ErrorLevel LogLevel = "ERROR"
)

func (l LogLevel) ToSlogLevel() slog.Level {
	switch l {
	case DebugLevel:
		return slog.LevelDebug
	case InfoLevel:
		return slog.LevelInfo
	case WarnLevel:
		return slog.LevelWarn
	case ErrorLevel:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

type ContextHandler struct {
	slog.Handler
	contextKeys []string
}

// Handle overrides the default Handle method to add context values.
func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, key := range h.contextKeys {
		if value := ctx.Value(key); value != nil {
			r.AddAttrs(slog.Any(key, value))
		}
	}
	return h.Handler.Handle(ctx, r)
}

type Config struct {
	// LogLevel is the level of logging to be used.
	// It can be one of the following:
	// DEBUG, INFO, WARN, ERROR
	LogLevel LogLevel

	// PrettyPrint is a boolean value that indicates whether to use pretty print or not.
	PrettyPrint bool

	// LogToFile is a boolean value that indicates whether to log to a file or not.
	LogToFile bool

	// FilePath is the path to the log file.
	FilePath string

	// FileMaxSize is the maximum size of the log file in MB.
	FileMaxSize int

	// FileMaxAge is the maximum age of the log file in days.
	FileMaxAge int

	// FileMaxBackups is the maximum number of backup files to keep.
	FileMaxBackups int

	// FileCompress is a boolean value that indicates whether to compress the log file or not.
	FileCompress bool

	// contextKeys is a list of keys to extract from the context
	// and add to the log record.
	// e.g. "request_id", "trace_id", "user_id", etc.
	ContextKeys []string
}

func SetupLogger(config *Config) *slog.Logger {
	if config == nil {
		config = &Config{}
	}
	if config.LogToFile {
		if config.FilePath == "" {
			config.FilePath = "./logs"
		}
		if config.FileMaxSize < 1 {
			config.FileMaxSize = 1 // in MB
		}
	}

	var handler slog.Handler
	var writer io.Writer
	var handlerOptions = slog.HandlerOptions{
		//AddSource: true,
		Level: config.LogLevel.ToSlogLevel(),
	}

	if config.LogToFile {
		writer = &lumberjack.Logger{
			Filename:   config.FilePath + "/service.log",
			MaxSize:    config.FileMaxSize,
			MaxBackups: config.FileMaxBackups,
			MaxAge:     config.FileMaxAge,
			Compress:   config.FileCompress,
		}
	} else {
		writer = os.Stdout
	}

	if config.PrettyPrint {
		//handler = slog.NewTextHandler(writer, &handlerOptions)
		handler = prettyslog.NewHandler(writer, &handlerOptions)
	} else {
		handler = slog.NewJSONHandler(writer, &handlerOptions)
	}

	logger := slog.New(&ContextHandler{handler, config.ContextKeys})
	slog.SetDefault(logger)

	return logger
}
