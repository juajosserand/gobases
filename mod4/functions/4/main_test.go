package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	Name      string
	Operation func(...int) float64
	Input     []int
	Expected  float64
}

func TestOperation(t *testing.T) {
	tests := []Test{
		// min
		{"TestMin", min, []int{2, 3, 3, 4, 10, 2, 4, 5}, 2},
		{"TestMinEmpty", min, []int{}, 0},
		{"TestMinNil", min, nil, 0},
		// avg
		{"TestAvg", avg, []int{2, 3, 3, 4, 1, 2, 4, 5}, 3},
		{"TestAvgEmpty", avg, []int{}, 0},
		{"TestAvgNil", avg, nil, 0},
		// max
		{"TestMax", max, []int{2, 3, 3, 4, 1, 2, 4, 5}, 5},
		{"TestMaxEmpty", max, []int{}, 0},
		{"TestMaxNil", max, nil, 0},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := test.Operation(test.Input...)
			assert.Equal(t, test.Expected, r)
		})
	}
}
