package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	Name     string
	Input    float64
	Expected float64
}

func TestTax(t *testing.T) {
	tests := []Test{
		{"TestTaxUnder50000", 49999, 0},
		{"TestTaxOver50000", 50001, 0.17},
		{"TestTaxOver150000", 150001, 0.27},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			tax, err := tax(test.Input)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, test.Expected, tax)
		})
	}
}

func TestTaxError(t *testing.T) {
	tests := []Test{
		{"TestTaxError", -1, 0},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			tax, err := tax(test.Input)
			assert.NotNil(t, err)
			assert.Equal(t, test.Expected, tax)
		})
	}
}
