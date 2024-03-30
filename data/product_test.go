package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "roy",
		Price: 2.00,
		SKU:   "ajf-sdfsd-sdf",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
