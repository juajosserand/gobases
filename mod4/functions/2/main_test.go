package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	Name     string
	Input    []int
	Expected float64
}

func TestAvg(t *testing.T) {
	tests := []Test{
		{"TestAvgOk", []int{2, 3, 3, 4, 1, 2, 4, 5}, 3},
		{"TestAvgNilScores", nil, 0},
		{"TestAvgEmptyScores", []int{}, 0},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			avg, err := avg(test.Input...)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, test.Expected, avg)
		})
	}
}

func TestAvgErrorNegative(t *testing.T) {
	input := []int{-2, 3, 3, 4, 1, 2, 4, 5}
	var expected float64 = 0

	avg, err := avg(input...)
	assert.NotNil(t, err)
	assert.Equal(t, expected, avg)
}
