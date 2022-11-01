package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusMetricsService struct {
	requestDuration *prometheus.HistogramVec
}

func NewPrometheusMetricsService(requestDuration *prometheus.HistogramVec) *PrometheusMetricsService {
	return &PrometheusMetricsService{requestDuration}
}

func (this *PrometheusMetricsService) ObserveRequest(r *http.Request, statusCode int, requestDuration float64) {
	statusCodeLabelValue := strconv.Itoa(statusCode)
	endpointLabelValue := fmt.Sprintf("%s %s", r.Method, r.URL.Path)

	this.requestDuration.WithLabelValues(endpointLabelValue, statusCodeLabelValue).Observe(requestDuration)
}
