package main

import (
	"errors"
	"fmt"
)

var (
	ErrMinSalary = errors.New("salary less than 10.000")
)

func main() {
	var salary float64 = 10000

	err := validateSalary(salary)

	if errors.Is(err, ErrMinSalary) {
		fmt.Println(err)
	}
}

func validateSalary(s float64) error {
	if s <= 10000 {
		return ErrMinSalary
	}

	return nil
}
