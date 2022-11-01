package container

import (
	"context"
	"net/http"
)

type Config struct {
	RedisEndpoint  string
	RedisUser      string
	RedisPassword  string
	JaegerEndpoint string
	AppPort        string
	TraceExporter  string
}

type Logger interface {
	Log(context.Context, string)
	LogStartup(string)
	LogShutdown(string)
	LogError(context.Context, string, error)
	Shutdown()
}

type CounterService interface {
	SetCounter(context.Context, int) error
	IncrementCounter(context.Context) (int, error)
}

type MetricsService interface {
	ObserveRequest(*http.Request, int, float64)
}
