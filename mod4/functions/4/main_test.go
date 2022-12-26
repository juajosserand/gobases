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
		{"TestMin", min, []int{3, 2, 3, 4, 10, 2, 4, 5}, 2},
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
			r, err := Operation(test.Operation, test.Input...)
			assert.Nil(t, err)
			assert.Equal(t, test.Expected, r)
		})
	}
}

func TestOperation_ErrNilOp(t *testing.T) {
	var expected float64 = 0
	r, err := Operation(nil, []int{}...)
	assert.NotNil(t, err)
	assert.Equal(t, expected, r)
}
