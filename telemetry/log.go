package telemetry

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcLoggerExporter(ctx context.Context, host string, port int) (*otlploggrpc.Exporter, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	exp, err := otlploggrpc.New(ctx, otlploggrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, err
	}
	return exp, nil
}

func NewLoggerProvider(exporter *otlploggrpc.Exporter, r *resource.Resource) (*log.LoggerProvider, error) {
	return log.NewLoggerProvider(
		log.WithResource(r),
		log.WithProcessor(log.NewBatchProcessor(exporter)),
	), nil
}
