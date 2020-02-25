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
	listenAddr  = flag.String("web.listen-port", "9001", "An port to listen on for web interface and telemetry.")
	metricsPath = flag.String("web.telemetry-path", "/metrics", "A path under which to expose metrics.")
	metricsName = flag.String("metric.name", "product_demo", "Prometheus metrics name")
)


func main() {
	flag.Parse()

	metrics := collector.NewMetrics(*metricsName)
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)

	http.Handle(*metricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>A Prometheus Exporter</title></head>
			<body>
			<h1>A Prometheus Exporter</h1>
			<p><a href='/metrics'>Metrics</a></p>
			</body>
			</html>`))
	})

	log.Printf("Starting Server at http://localhost:%s%s", *listenAddr, *metricsPath)
	log.Fatal(http.ListenAndServe(":"+*listenAddr, nil))
}

