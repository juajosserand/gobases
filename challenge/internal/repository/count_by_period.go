package repository

import (
	"strconv"
	"strings"

	"github.com/bootcamp-go/desafio-go-bases/internal/ticket"
)

func (r *repository) GetCountByPeriod(time string) (total int, err error) {
	var start, end int

	switch {
	case time == ticket.Dawning:
		start, end = 0, 6
	case time == ticket.Morning:
		start, end = 7, 12
	case time == ticket.Afternoon:
		start, end = 13, 19
	case time == ticket.Night:
		start, end = 20, 23
	default:
		return 0, errInvalidDayTimePeriod
	}

	for _, t := range r.Tickets {
		hourStr := strings.Split(t.DepartureTime, ":")[0]
		hour, err := strconv.Atoi(hourStr)
		if err != nil {
			return 0, err
		}

		if start <= hour && hour <= end {
			total++
		}
	}

	return total, nil
}
