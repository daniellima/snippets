package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"github.com/daniellima/counting-api/app"
	"github.com/daniellima/counting-api/app/base/container"
	"github.com/daniellima/counting-api/app/services"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func bootstrap() *sdktrace.TracerProvider {
	var err error

	container.SetProvider("config", func(c *container.Container) interface{} {
		config := container.Config{}

		config.AppPort = os.Getenv("APP_PORT")
		if config.AppPort == "" {
			log.Fatalf("The APP_PORT environment variable must be defined and be equal to a valid unix port number")
		}
		config.RedisUser = os.Getenv("REDIS_USER")
		if config.RedisUser == "" {
			log.Fatalf("The variable REDIS_USER must be defined")
		}
		config.RedisPassword = os.Getenv("REDIS_PASSWORD")
		if config.RedisPassword == "" {
			log.Fatalf("The variable REDIS_PASSWORD must be defined")
		}
		config.RedisEndpoint = os.Getenv("REDIS_ENDPOINT")
		if config.RedisEndpoint == "" {
			log.Fatalf("The variable REDIS_ENDPOINT must be defined")
		}
		config.JaegerEndpoint = os.Getenv("JAEGER_ENDPOINT")
		if config.JaegerEndpoint == "" {
			log.Fatalf("The variable JAEGER_ENDPOINT must be defined and be a valid endpoint")
		}
		config.TraceExporter = os.Getenv("TRACE_EXPORTER")
		if config.TraceExporter == "" {
			log.Fatalf("The variable TRACE_EXPORTER must be defined and be 'stdout' or 'jaeger'")
		}

		return config
	})

	container.SetProvider("logger", func(c *container.Container) interface{} {
		logger, err := zap.NewProduction()
		if err != nil {
			panic(fmt.Sprintf("Could not build zap logger %v", err))
		}
		return services.NewZapLoggerService(logger)
	})

	container.SetProvider("counterService", func(c *container.Container) interface{} {
		config := container.GetConfig()

		return services.NewRedisService(context.Background(), fmt.Sprintf("redis://%s:%s@%s", config.RedisUser, config.RedisPassword, config.RedisEndpoint))
	})

	container.SetProvider("metricsService", func(c *container.Container) interface{} {
		requestDuration := prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "The duration of a request",
				Buckets: []float64{0.1, 0.3, 0.5, 0.7, 1.0, 1.5, 2.0},
			},
			[]string{"endpoint", "status_code"},
		)
		prometheus.MustRegister(requestDuration)

		return services.NewPrometheusMetricsService(requestDuration)
	})

	var (
		traceProvider *sdktrace.TracerProvider
		exporter      sdktrace.SpanExporter
	)

	config := container.GetConfig()
	if config.TraceExporter == "stdout" {
		exporter, err = stdouttrace.New(
			stdouttrace.WithWriter(os.Stdout),
			stdouttrace.WithPrettyPrint(),
			stdouttrace.WithoutTimestamps(),
		)
		if err != nil {
			log.Fatalf("Cannot initialize stdouttrace for OpenTelemetry lib: %v", err)
		}
	} else {
		exporter, err = jaeger.New(jaeger.WithAgentEndpoint(jaeger.WithAgentHost(config.JaegerEndpoint)))
		if err != nil {
			log.Fatalf("Cannot initialize Jaeger exporter for OpenTelemetry lib: %v", err)
		}
	}

	resource, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("onboarding-counter-api"),
			semconv.ServiceVersionKey.String("0.0.0"),
			attribute.String("type", "onboarding"),
		),
	)
	if err != nil {
		log.Fatalf("Cannot initialize stdouttrace for OpenTelemetry lib: %v", err)
	}

	traceProvider = sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource),
	)

	otel.SetTracerProvider(traceProvider)

	return traceProvider
}

func main() {
	var err error

	traceProvider := bootstrap()
	defer func() {
		err = traceProvider.Shutdown(context.Background())
		if err != nil {
			log.Fatalf("Cannot shutdown trace provider: %v", err)
		}
	}()

	logger := container.GetLogger()
	defer logger.Shutdown()

	config := container.GetConfig()

	logger.LogStartup("🚀 Starting app...")

	logger.LogStartup(fmt.Sprintf("📝 Started with the following env variables: %v", map[string]string{
		"APP_PORT":        config.AppPort,
		"REDIS_ENDPOINT":  config.RedisEndpoint,
		"REDIS_USER":      config.RedisUser,
		"REDIS_PASSWORD":  "***",
		"JAEGER_ENDPOINT": config.JaegerEndpoint,
		"TRACE_EXPORTER":  config.TraceExporter,
	}))

	serverMux := http.DefaultServeMux

	app.ConfigureRoutes(serverMux)

	logger.LogStartup(fmt.Sprintf("🎧 Listening on 0.0.0.0:%s...", config.AppPort))
	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", config.AppPort), serverMux)

	logger.LogShutdown(fmt.Sprintf("👋 Exiting app... Reason %v", err))
}
