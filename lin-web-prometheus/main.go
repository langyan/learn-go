package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "go_requests_total",
	Help: "已处理请求总数",
})

func init() {
	prometheus.MustRegister(requestCount)
}

func handler(w http.ResponseWriter, r *http.Request) {
	requestCount.Inc()
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
