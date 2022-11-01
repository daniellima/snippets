package app

import (
	"net/http"

	"github.com/daniellima/counting-api/app/base/middleware"
	"github.com/daniellima/counting-api/app/base/middleware/when"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func route(s *http.ServeMux, path string, handler http.Handler) {
	s.Handle(path, middleware.TraceHandler(middleware.MeasureDuration(handler)))
}

func ConfigureRoutes(s *http.ServeMux) {

	route(s, "/api/v1", when.Get(HelloWorldHandler))
	route(s, "/api/v1/count", when.Get(ShowCountHandler).And(when.Post(UpdateCountHandler)))
	route(s, "/readyz", when.Get(ReadyzHandler))
	route(s, "/livez", when.Get(ReadyzHandler))

	route(s, "/metrics", promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	))

}
