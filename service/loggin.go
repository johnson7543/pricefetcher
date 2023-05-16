package service

import (
	"context"
	"time"

	"github.com/johnson7543/pricefetcher/types"
	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next PriceService
}

func NewLoggingService(next PriceService) PriceService {
	return &loggingService{
		next: next,
	}
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: false,
		DisableColors:    false,
		DisableSorting:   false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05.000",
	})
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		const uuidKey types.ContextKey = "uuid"
		// duration := fmt.Sprintf("%.6f/sec", time.Since(begin).Seconds())
		duration := time.Since(begin)

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"duration": duration,
				"ticker":   ticker,
				"error":    err.Error(),
			}).Errorf("[%s] FetchPrice", ctx.Value(uuidKey))
		} else {
			logrus.WithFields(logrus.Fields{
				"duration": duration,
				"ticker":   ticker,
				"price":    price,
			}).Infof("[%s] FetchPrice", ctx.Value(uuidKey))
		}

	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}
