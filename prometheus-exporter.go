package main

import (
"log"
"flag"
"net/http"

"github.com/huanghuangzym/prometheus-exporter/collector"
"github.com/prometheus/client_golang/prometheus"
"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listenAddr  = flag.String("port", "9001", "An port to listen on for web interface and telemetry.")
	metricsName = flag.String("name", "product_metric_demo", "Prometheus metrics name")
	metricsAppName = flag.String("app", "asm", "Prometheus metrics name")
	metricsGuageName = flag.String("label", "istio", "Prometheus metrics name")

)


func main() {
	flag.Parse()

	metricPath := "/metrics"

	metrics := collector.NewMetrics(*metricsName, *metricsAppName, *metricsGuageName)
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)

	http.Handle(metricPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>A Prometheus Exporter</title></head>
			<body>
			<h1>A Prometheus Exporter</h1>
			<p><a href='/metrics'>Metrics</a></p>
			</body>
			</html>`))
	})

	log.Printf("Starting Server at http://localhost:%s%s", *listenAddr, metricPath)
	log.Fatal(http.ListenAndServe(":"+*listenAddr, nil))
}

