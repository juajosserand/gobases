package main

import (
	"fmt"
)

func main() {
	var salary float64

	if salary < 150000 {
		fmt.Println(ErrTaxMin{})
	} else {
		fmt.Println("must pay tax")
	}
}

type ErrTaxMin struct{}

func (s ErrTaxMin) Error() string {
	return "salary doesn't reach the taxable minimum"
}
