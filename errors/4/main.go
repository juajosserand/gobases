package main

import (
	"fmt"
)

func main() {
	var salary float64 = 10000

	err := validateSalary(salary)
	if err != nil {
		panic(err)
	}

	fmt.Println("must pay tax")
}

func validateSalary(s float64) error {
	if s < 150000 {
		return fmt.Errorf(
			"taxable minimum is 150.000 and given salary is %.2f",
			s,
		)
	}

	return nil
}
