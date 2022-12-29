package main

import (
	"fmt"

	"github.com/juajosserand/gobases/challenge/internal/repository"
	"github.com/juajosserand/gobases/challenge/internal/ticket"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("errors were detected at runtime")
			fmt.Println("error:", err)
		}
	}()

	tr, err := repository.NewCSVRepository("./tickets.csv")
	if err != nil {
		panic(err)
	}

	total, err := tr.GetTotalTickets("Brazil")
	if err != nil {
		panic(err)
	}
	fmt.Println("total tickets for Brazil:", total)

	total, err = tr.GetCountByPeriod(ticket.Morning)
	if err != nil {
		panic(err)
	}
	fmt.Printf("count for period %s is: %d\n", ticket.Morning, total)

	avg, err := tr.AverageDestination("Brazil")
	if err != nil {
		panic(err)
	}
	fmt.Println("average destination for Brazil:", avg)
}
