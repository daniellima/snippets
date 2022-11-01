package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/daniellima/counting-api/app/base/container"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type ResponseWriterWithStatus struct {
	http.ResponseWriter
	Status int
}

func (resp *ResponseWriterWithStatus) WriteHeader(status int) {
	if resp.Status != 0 {
		resp.Status = status
	}
	resp.ResponseWriter.WriteHeader(status)
}

func MeasureDuration(handler http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		newResponseWriter := &ResponseWriterWithStatus{w, http.StatusOK}
		handler.ServeHTTP(newResponseWriter, r)

		container.GetMetricsService().ObserveRequest(r, newResponseWriter.Status, time.Since(start).Seconds())
	}
}

func TraceHandler(handler http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		operation := fmt.Sprintf("%s %s", r.Method, r.URL.Path)

		ctx, span := otel.Tracer("onboarding-counter-api").Start(r.Context(), operation)
		defer span.End()

		newResponseWriter := &ResponseWriterWithStatus{w, http.StatusOK}
		handler.ServeHTTP(newResponseWriter, r.WithContext(ctx))

		statusCode := strconv.Itoa(newResponseWriter.Status)
		span.SetAttributes(
			attribute.String("status", statusCode),
		)

	}
}
