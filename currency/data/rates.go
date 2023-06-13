package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/go-hclog"
)

// Model for exchanging rates
type ExchangeRates struct {
	log hclog.Logger
	rates map[string]float64
}

// creates a new rates exchanger
func NewExchangeRates(log hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{log: log, rates: map[string]float64{}}
	err := er.getRates()
	return er, err
}


func (e *ExchangeRates) getRates() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected status code is 200 but got %d", resp.StatusCode)
	}
	resp.Body.Close()
	md := &Cubes{}
	xml.NewDecoder(resp.Body).Decode(md)
	for _, cube := range md.CubeData {
		r, err := strconv.ParseFloat(cube.Rate, 64)
		if err != nil {
			return err
		}

		e.rates[cube.Currency] = r
	}
	return nil
}



type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}


type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate string `xml:"rate,attr"`
}


