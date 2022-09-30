package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func MetricsWebServer() {
	fmt.Println("Starting things up")

	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(os.Stdout),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		// stdouttrace.WithoutTimestamps(),
	)
	if err != nil {
		panic(err)
	}

	resource, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("fib"),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	if err != nil {
		panic(err)
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource),
	)
	defer func() {
		err = traceProvider.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	otel.SetTracerProvider(traceProvider)

	ctx, span := otel.Tracer("name").Start(context.Background(), "First")

	ctx, span2 := otel.Tracer("name").Start(ctx, "Second")

	time.Sleep(time.Second * time.Duration(2))

	span2.SetStatus(codes.Error, "A error")
	span2.End()

	span.End()

	requestNumber := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "request_number",
			Help: "the number of requests",
		},
		[]string{"path"},
	)

	requestCount := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "request_counter_duration",
			Help: "hello",
		},
		[]string{"path"},
	)

	prometheus.MustRegister(requestNumber)
	prometheus.MustRegister(requestCount)

	requestNumber.WithLabelValues("/hello").Add(2)

	http.HandleFunc("/add", func(response http.ResponseWriter, request *http.Request) {

		n, err := strconv.Atoi(request.URL.Query().Get("amount"))
		if err != nil {
			panic("Amount is not sent")
		}

		requestNumber.WithLabelValues("path").Add(float64(n))
		requestCount.WithLabelValues("/add").Observe(float64(n))

		response.WriteHeader(http.StatusOK)
	})

	http.Handle("/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))

	err = http.ListenAndServe("0.0.0.0:9000", nil)
	if err != nil {
		panic(err)
	}

}
