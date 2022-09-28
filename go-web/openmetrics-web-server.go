package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func OpenMetricsWebServer() {
	fmt.Println("Starting things up")

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

	err := http.ListenAndServe("0.0.0.0:9000", nil)
	if err != nil {
		panic(err)
	}

}
