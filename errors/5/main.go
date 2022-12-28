package main

import (
	"errors"
	"fmt"
)

func main() {
	s, err := salary(1000, 150)
	if err != nil {
		panic(err)
	}

	fmt.Println("salary:", s)
}

func salary(hours float64, value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("invalid negative hourly value")
	}

	if hours < 80 {
		return 0, errors.New("worker cannot have worked less than 80 hours/month")
	}

	if hours*value >= 150000 {
		return hours * value * 0.9, nil
	}

	return hours * value, nil
}
