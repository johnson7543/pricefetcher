package main

import (
	"context"
	"time"

	"github.com/johnson7543/pricefetcher/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	fetchPriceCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "fetch_price_calls_total",
		Help: "Total number of calls to FetchPrice method.",
	}, []string{"status"})
	fetchPriceDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "fetch_price_duration_seconds",
		Help:    "Time taken to execute FetchPrice method.",
		Buckets: []float64{0.01, 0.05, 0.1, 0.5, 1, 2, 5, 10},
	})
)

type metricService struct {
	next service.PriceService
}

func NewMetricService(next service.PriceService) service.PriceService {
	return &metricService{
		next: next,
	}
}

func (s *metricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start).Seconds()
		status := "success"
		if err != nil {
			status = "error"
		}
		fetchPriceCounter.WithLabelValues(status).Inc()
		fetchPriceDuration.Observe(duration)
	}()

	// fmt.Println("pushing metrics to prometheus")
	// your metrics storage. Push to prometheus (gauge, counters)
	return s.next.FetchPrice(ctx, ticker)
}
