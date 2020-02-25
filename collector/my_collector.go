package collector

import (
	"sync"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
)

// 指标结构体
type Metrics struct {
	metrics map[string]*prometheus.Desc
	mutex   sync.Mutex
}


func newGlobalMetric(namespace string, metricName string, docString string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(namespace+"_"+metricName, docString, labels, nil)
}



func NewMetrics(namespace string) *Metrics {
	return &Metrics{
		metrics: map[string]*prometheus.Desc{
			"my_counter_metric": newGlobalMetric(namespace, "my_counter_metric", "The description of my_counter_metric", []string{"host"}),
			"my_gauge_metric": newGlobalMetric(namespace, "my_gauge_metric","The description of my_gauge_metric", []string{"host"}),
		},
	}
}

func (c *Metrics) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		ch <- m
	}
}


func (c *Metrics) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()  // 加锁
	defer c.mutex.Unlock()

	mockCounterMetricData, mockGaugeMetricData := c.GenerateMockData()
	for host, currentValue := range mockCounterMetricData {
		ch <-prometheus.MustNewConstMetric(c.metrics["my_counter_metric"], prometheus.CounterValue, float64(currentValue), host)
	}
	for host, currentValue := range mockGaugeMetricData {
		ch <-prometheus.MustNewConstMetric(c.metrics["my_gauge_metric"], prometheus.GaugeValue, float64(currentValue), host)
	}
}


 func (c *Metrics) GenerateMockData() (mockCounterMetricData map[string]int, mockGaugeMetricData map[string]int) {
 	mockCounterMetricData = map[string]int{
		"yahoo.com": int(rand.Int31n(1000)),
		"google.com": int(rand.Int31n(1000)),
	}
	mockGaugeMetricData = map[string]int{
		"yahoo.com": int(rand.Int31n(10)),
		"google.com": int(rand.Int31n(10)),
	}
	return
 }

