package server

import (
	"context"

	"github.com/zakisk/microservice/currency/data"
	protos "github.com/zakisk/microservice/currency/protos/currency"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	rates *data.ExchangeRates
	l hclog.Logger
	protos.UnimplementedCurrencyServer
}

// creates new Currency	
func NewCurrency(rates *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{rates: rates, l: l}
}


func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.l.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Rate: rate}, nil
} 


func (c *Currency) mustEmbedUnimplementedCurrencyServer() { }

