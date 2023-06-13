package server

import (
	"context"

	protos "github.com/zakisk/microservice/currency/protos/currency"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	l hclog.Logger
	protos.UnimplementedCurrencyServer
}

// creates new Currency	
func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l: l}
}


func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.l.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
} 


func (c *Currency) mustEmbedUnimplementedCurrencyServer() { }

