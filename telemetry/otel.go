package telemetry

import (
	"context"
	"errors"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"log/slog"
)

type Config struct {
	ServiceName string
	Version     string
	GrpcHost    string
	GrpcPort    int
}

func SetupOtelSDK(ctx context.Context, cfg *Config) (func(context.Context) error, error) {
	if cfg == nil {
		err := errors.New("telemetry: Invalid configuration, must not be nil")
		slog.ErrorContext(ctx, err.Error())
		return nil, err
	}

	if cfg.ServiceName == "" {
		cfg.ServiceName = "go-service"
	}
	if cfg.Version == "" {
		cfg.Version = "1.0.0"
	}
	if cfg.GrpcHost == "" {
		cfg.GrpcHost = "localhost"
	}
	if cfg.GrpcPort == 0 {
		cfg.GrpcPort = 4317 // Default OTLP gRPC port
	}

	var shutdownFuncs []func(context.Context) error
	shutdown := func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(cfg.ServiceName),
			semconv.ServiceVersion(cfg.Version),
		),
	)
	if err != nil {
		slog.ErrorContext(ctx, "telemetry: Failed to create resource", "error", err.Error())
		return nil, err
	}

	matrixGrpcExporter, err := NewGrpcMetricExporter(ctx, cfg.GrpcHost, cfg.GrpcPort)
	if err != nil {
		slog.ErrorContext(ctx, "telemetry: Failed to create metric exporter", "error", err.Error())
		return nil, err
	}
	matrixProvider, err := NewMeterProvider(matrixGrpcExporter, r)
	if err != nil {
		slog.ErrorContext(ctx, "telemetry: Failed to create meter provider", "error", err.Error())
		return nil, err
	}
	shutdownFuncs = append(shutdownFuncs, matrixProvider.Shutdown)

	traceGrpcExporter, err := NewGrpcTraceExporter(ctx, cfg.GrpcHost, cfg.GrpcPort)
	if err != nil {
		slog.ErrorContext(ctx, "telemetry: Failed to create trace exporter", "error", err.Error())
		return nil, err
	}
	traceProvider, err := NewTracerProvider(traceGrpcExporter, r)
	if err != nil {
		slog.ErrorContext(ctx, "telemetry: Failed to create tracer provider", "error", err.Error())
		return nil, err
	}
	shutdownFuncs = append(shutdownFuncs, traceProvider.Shutdown)

	logGrpcExporter, err := NewGrpcLoggerExporter(ctx, cfg.GrpcHost, cfg.GrpcPort)
	if err != nil {
		slog.ErrorContext(ctx, "telemetry: Failed to create log exporter", "error", err.Error())
		return nil, err
	}
	logProvider, err := NewLoggerProvider(logGrpcExporter, r)
	if err != nil {
		slog.ErrorContext(ctx, "telemetry: Failed to create logger provider", "error", err.Error())
		return nil, err
	}
	shutdownFuncs = append(shutdownFuncs, logProvider.Shutdown)

	return shutdown, nil
}

func SlogBridge(provider *log.LoggerProvider) {
	logger := otelslog.NewLogger(
		"otel-slog-bridge",
		otelslog.WithLoggerProvider(provider),
	)
	slog.SetDefault(logger)
}
