package collector

import (
	"sync"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
)

// 指标结构体
type Metrics struct {
	metricName string
	metrics map[string]*prometheus.Desc
	mutex   sync.Mutex
}


func newGlobalMetric(metricName string, docString string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(metricName, docString, labels, nil)
}



func NewMetrics(name string) *Metrics {
	return &Metrics{
		metrics: map[string]*prometheus.Desc{
			name: newGlobalMetric( name,"The description of my_gauge_metric", []string{"host"}),
		},
		metricName: name,
	}
}

func (c *Metrics) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		ch <- m
	}
}


func (c *Metrics) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	mockGaugeMetricData := c.GenerateMockData()
	for host, currentValue := range mockGaugeMetricData {
		ch <-prometheus.MustNewConstMetric(c.metrics[c.metricName], prometheus.GaugeValue, float64(currentValue), host)
	}
}


 func (c *Metrics) GenerateMockData() (mockGaugeMetricData map[string]int) {
	mockGaugeMetricData = map[string]int{
		"yahoo.com": int(rand.Int31n(10)),
		"google.com": int(rand.Int31n(10)),
	}
	return
 }

