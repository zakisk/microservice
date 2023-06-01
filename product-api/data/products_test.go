package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "Zaki",
		Price: 10,
		SKU: "dsf-dfkf-sdfk",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}