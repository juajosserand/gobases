package repository

import (
	"strconv"
	"testing"

	"github.com/juajosserand/gobases/challenge/internal/ticket"
	"github.com/stretchr/testify/assert"
)

func TestGetCountByPeriod(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Dawning", ticket.Dawning, 304},
		{"Morning", ticket.Morning, 256},
		{"Afternoon", ticket.Afternoon, 289},
		{"Night", ticket.Night, 151},
	}

	r, _ := NewCSVRepository("../../tickets.csv")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			total, err := r.GetCountByPeriod(test.input)
			assert.Nil(t, err)
			assert.Equal(t, test.expected, total)
		})
	}
}

func TestGetCountByPeriod_ErrInvalidDayPeriod(t *testing.T) {
	test := struct {
		input    string
		expected int
		err      error
	}{
		"", 0, errInvalidDayTimePeriod,
	}

	r, _ := NewCSVRepository("../../tickets.csv")

	total, err := r.GetCountByPeriod(test.input)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, test.err)
	assert.Equal(t, test.expected, total)
}

func TestGetCountByPeriod_ErrHourConversion(t *testing.T) {
	test := struct {
		input    string
		expected int
		err      error
	}{
		ticket.Dawning, 0, strconv.ErrSyntax,
	}

	r := &repository{
		filename: "../../tickets.csv",
	}
	r.Tickets = append(r.Tickets, ticket.Ticket{
		DepartureTime: "abc:30", // forcing invalid ticket
	})

	total, err := r.GetCountByPeriod(test.input)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, test.err)
	assert.Equal(t, test.expected, total)
}
