package test

import (
	"context"
	"net/http"
)

type NopLogger struct{}

func (this *NopLogger) Log(ctx context.Context, message string)                 {}
func (this *NopLogger) LogStartup(message string)                               {}
func (this *NopLogger) LogShutdown(message string)                              {}
func (this *NopLogger) LogError(ctx context.Context, message string, err error) {}
func (this *NopLogger) Shutdown()                                               {}

type NopMetricsService struct{}

func (this *NopMetricsService) ObserveRequest(r *http.Request, statusCode int, requestDuration float64) {
}
