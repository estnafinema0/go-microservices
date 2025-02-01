package data

import "testing"

func TestChackValidation(t *testing.T) {
	p := &Product{
		Name:  "kat",
		Price: 2.9,
		SKU:   "agf-k-k",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
