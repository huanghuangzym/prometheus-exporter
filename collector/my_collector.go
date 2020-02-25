package collector

import (
	"fmt"
	"sync"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
)

// 指标结构体
type Metrics struct {
	metricName string
	metricAppName string
	metricGuageName string
	metrics map[string]*prometheus.Desc
	mutex   sync.Mutex
}

const  AppLabelKey = "APP"
const  AppVersionKey = "VERSION"
const  MetricKey = "Metric"



func newGlobalMetric(metricName string, docString string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(metricName, docString, labels, nil)
}



func NewMetrics(name, appname, guagename string) *Metrics {
	return &Metrics{
		metrics: map[string]*prometheus.Desc{
			name: newGlobalMetric( name,fmt.Sprintf("The description of %s",name), []string{AppLabelKey,AppVersionKey,MetricKey}),
		},
		metricName: name,
		metricAppName: appname,
		metricGuageName: guagename,
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
		ch <-prometheus.MustNewConstMetric(c.metrics[c.metricName], prometheus.GaugeValue, float64(currentValue), c.metricAppName,"v1",host)
	}
}


 func (c *Metrics) GenerateMockData() (mockGaugeMetricData map[string]int) {
	mockGaugeMetricData = map[string]int{
		fmt.Sprintf("total_req_%s",c.metricGuageName): int(rand.Int31n(10)),
		fmt.Sprintf("total_send_%s",c.metricGuageName): int(rand.Int31n(10)),
	}
	return
 }

