package server

import (
	"context"
	"io"
	"time"

	"github.com/zakisk/microservice/currency/data"
	protos "github.com/zakisk/microservice/currency/protos/currency"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	rates         *data.ExchangeRates
	l             hclog.Logger
	subscriptions map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest
	protos.UnimplementedCurrencyServer
}

// creates new Currency
func NewCurrency(rates *data.ExchangeRates, l hclog.Logger) *Currency {
	c := &Currency{
		rates:         rates,
		l:             l,
		subscriptions: make(map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest),
	}
	go c.handleUpdates()
	return c
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.l.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	if rr.Base.String() == rr.Destination.String() {
		err := status.Newf(
			codes.InvalidArgument,
			"Base currency %s can't be same as destination %s",
			rr.Base.String(),
			rr.Destination.String(),
		)

		err, wde := err.WithDetails(rr)
		if wde != nil {
			return nil, wde
		}

		return nil, err.Err()
	}
	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: rate}, nil
}

func (c *Currency) handleUpdates() {
	ru := c.rates.MonitorRates(5 * time.Second)
	for range ru {
		c.l.Info("Got Rate Update")
		// iterate over subscribed clients
		for k, v := range c.subscriptions {
			//iterate over subscribed rates by client `k`
			for _, rr := range v {
				rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
				if err != nil {
					c.l.Error(
						"Unable to get updated rate",
						"base",
						rr.GetBase().String(),
						"destination",
						rr.GetDestination().String(),
						"error",
						err,
					)
				} else {
					err = k.Send(
						&protos.StreamingRateResponse{
							Response: &protos.StreamingRateResponse_RateResponse{
								RateResponse: &protos.RateResponse{Base: rr.GetBase(), Destination: rr.GetDestination(), Rate: rate},
							},
						},
					)
					if err != nil {
						c.l.Error(
							"Unable to send updated rate",
							"base",
							rr.GetBase().String(),
							"destination",
							rr.GetDestination().String(),
							"error",
							err,
						)
					}
				}
			}
		}
	}
}

func (c *Currency) SubscribeRates(csr protos.Currency_SubscribeRatesServer) error {

	for {
		rr, err := csr.Recv()
		if err == io.EOF {
			c.l.Info("Client has close connection")
			break
		}

		if err != nil {
			c.l.Error("Unable to read from client", "error", err)
			return err
		}

		c.l.Info("Handle client request", "request", rr)

		rrs, ok := c.subscriptions[csr]
		if !ok {
			rrs = []*protos.RateRequest{}
		}

		var validationErr *status.Status
		for _, v := range rrs  {
			if v.Base == rr.Base && v.Destination == rr.Destination {
				// subscription is already exists
				validationErr := status.Newf(
					codes.AlreadyExists, 
					"Can't subscribe as subscription is already exist")
				validationErr, err := validationErr.WithDetails(rr)
				if err != nil {
					c.l.Error("Unable to add request struct to metadata", "error", err)
				}
				break
			}
		}

		if validationErr != nil {
			csr.Send(&protos.StreamingRateResponse{
				Response: &protos.StreamingRateResponse_Error{Error: validationErr.Proto()}},
			)
			continue
		}
		rrs = append(rrs, rr)
		c.subscriptions[csr] = rrs
	}

	return nil
}

func (c *Currency) mustEmbedUnimplementedCurrencyServer() {}

{
	"Base": "JPY",
	"Destination": "BGN"
}
