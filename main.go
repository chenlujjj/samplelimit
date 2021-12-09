package main

import (
	"flag"
	"fmt"
	"net/http"
	"sort"

	io_prometheus_client "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

type MetricStat struct {
	Name  string
	Count int
}

func main() {
	url := flag.String("url", "http://localhost:9090/metrics", "Prometheus metrics URL")
	flag.Parse()
	resp, err := http.Get(*url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	p := expfmt.TextParser{}
	metricFamilies, err := p.TextToMetricFamilies(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var total int
	stats := make([]MetricStat, 0, len(metricFamilies))
	for _, mf := range metricFamilies {
		var count int
		switch mf.GetType() {
		case io_prometheus_client.MetricType_COUNTER, io_prometheus_client.MetricType_GAUGE, io_prometheus_client.MetricType_UNTYPED:
			count = len(mf.Metric)
		case io_prometheus_client.MetricType_HISTOGRAM:
			// histogram类型的指标有sum，count还有buckets
			buckets := len(mf.Metric[0].Histogram.Bucket)
			count = len(mf.Metric) * (2 + buckets)
		case io_prometheus_client.MetricType_SUMMARY:
			// summary类型的指标有sum, count还有quantiles
			quantiles := len(mf.Metric[0].Summary.Quantile)
			count = len(mf.Metric) * (2 + quantiles)
		}
		stats = append(stats, MetricStat{Name: mf.GetName(), Count: count})
		total += count
	}
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Count < stats[j].Count
	})

	fmt.Println("Total samples:", total)
	fmt.Println("==============================")
	for _, stat := range stats {
		fmt.Println(stat.Name, stat.Count)
	}
}
