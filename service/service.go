package service

import (
	"context"
	"fmt"
)

// PriceService is an interface that can fetch a price.
type PriceService interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements the PriceFetcher interface.
type priceService struct{}

func NewPriceService() PriceService {
	return &priceService{}
}

func (s *priceService) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}

	return price, nil
}

var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"JJ":  100_000.0,
}
