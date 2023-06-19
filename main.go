package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	// demo metric
	cpuUsage = promauto.NewGauge(prometheus.GaugeOpts{
		// metric name
		Name: "cpu_usage",
		// provides information about this metri
		Help: "Current usage of the CPU.",
	})
)

func recordMetrics() {
	go func() {
		for {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			cpuUsage.Set(r.Float64())
			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	recordMetrics()
	// NOTE: /metrics is fixed for prometheus pull, keep it fixed
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("====================== start metrics =================")
	// fix port to 9999
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal(err)
	}
}
