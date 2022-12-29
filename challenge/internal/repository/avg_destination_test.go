package repository

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverageDestination(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"Brazil", "Brazil", 0.045},
	}

	r, _ := NewCSVRepository("../../tickets.csv")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			avg, err := r.AverageDestination(test.input)
			assert.Nil(t, err)
			fmt.Println(avg)
			assert.Equal(t, test.expected, avg)
		})
	}
}

func TestAverageDestination_Err(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
		err      error
	}{
		{"InvalidDestination", "", 0, errEmptyDestination},
	}

	r, _ := NewCSVRepository("../../tickets.csv")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			avg, err := r.AverageDestination(test.input)
			assert.NotNil(t, err)
			assert.ErrorIs(t, err, test.err)
			assert.Equal(t, avg, test.expected)
		})
	}
}
