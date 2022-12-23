package main

import (
	"errors"
)

func salary(workedMin float64, category string) (float64, error) {
	if workedMin < 0 {
		return 0, errors.New("invalid negative worked minutes")
	}

	switch {
	case category == "C":
		return 1000 * (workedMin / 60), nil
	case category == "B":
		return 1500 * (workedMin / 60) * 1.2, nil
	case category == "A":
		return 3000 * (workedMin / 60) * 1.5, nil
	default:
		return 0, errors.New("invalid category")
	}
}
