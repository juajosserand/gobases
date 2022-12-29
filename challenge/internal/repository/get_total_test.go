package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalTickets(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Brazil", "Brazil", 45},
	}

	r, err := NewCSVRepository("../../tickets.csv")
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			total, err := r.GetTotalTickets(test.input)
			assert.Nil(t, err)
			assert.Equal(t, total, test.expected)
		})
	}
}

func TestGetTotalTickets_Err(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"EmptyDestination", "", 0},
	}

	r, _ := NewCSVRepository("../../tickets.csv")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			total, err := r.GetTotalTickets(test.input)
			assert.NotNil(t, err)
			assert.Equal(t, total, test.expected)
		})
	}
}
