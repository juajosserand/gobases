package repository

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/bootcamp-go/desafio-go-bases/internal/ticket"
)

var (
	errEmptyDestination     = errors.New("invalid empty destination")
	errInvalidDayTimePeriod = errors.New("invalid day time period")
)

type TicketRepository interface {
	GetTotalTickets(string) (int, error)
	GetCountByPeriod(string) (int, error)
	AverageDestination(string) (float64, error)
}

type repository struct {
	filename string
	Tickets  []ticket.Ticket
}

func NewCSVRepository(fn string) (TicketRepository, error) {
	tr := &repository{
		filename: fn,
	}

	err := tr.readData(tr.filename)
	if err != nil {
		return tr, err
	}

	return tr, nil
}

func (r *repository) readData(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, row := range rows {
		id, err := strconv.ParseUint(row[0], 10, 64)
		if err != nil {
			return err
		}

		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return err
		}

		name, email, dc, dt := row[1], row[2], row[3], row[4]
		if name == "" || email == "" || dc == "" || dt == "" {
			log.Printf("skipping record %d with invalid fields\n", id)
			continue
		}

		// check valid email

		// check valid destination country

		// check valid time in departure time

		t := ticket.Ticket{
			ID:                 id,
			Name:               name,
			Email:              email,
			DestinationCountry: dc,
			DepartureTime:      dt,
			Price:              price,
		}

		r.Tickets = append(r.Tickets, t)
	}

	return nil
}
