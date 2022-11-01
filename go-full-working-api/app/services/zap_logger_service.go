package services

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type ZapLoggerService struct {
	logger *zap.Logger
}

func NewZapLoggerService(logger *zap.Logger) *ZapLoggerService {
	return &ZapLoggerService{logger}
}

func (this *ZapLoggerService) Log(ctx context.Context, message string) {
	span := trace.SpanFromContext(ctx)
	this.logger.Info(message,
		zap.String("traceId", span.SpanContext().TraceID().String()),
	)
}

func (this *ZapLoggerService) Shutdown() {
	this.logger.Sync()
}

func (this *ZapLoggerService) LogStartup(message string) {
	this.logger.Info(message, zap.String("type", "startup"))
}

func (this *ZapLoggerService) LogShutdown(message string) {
	this.logger.Info(message, zap.String("type", "shutdown"))
}

func (this *ZapLoggerService) LogError(ctx context.Context, message string, err error) {
	span := trace.SpanFromContext(ctx)
	this.logger.Info(message,
		zap.String("type", "error"),
		zap.Error(err),
		zap.String("traceId", span.SpanContext().TraceID().String()),
	)
}
