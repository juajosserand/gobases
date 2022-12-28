package main

import (
	"errors"
	"fmt"
)

func main() {
	var salary float64 = 10000

	err := validateSalary(salary)

	if errors.Is(err, ErrMinSalary{}) {
		fmt.Println(err)
	}
}

type ErrMinSalary struct{}

func (s ErrMinSalary) Error() string {
	return "salary less than 10.000"
}

func validateSalary(s float64) error {
	if s <= 10000 {
		return ErrMinSalary{}
	}

	return nil
}
