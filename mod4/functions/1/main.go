package main

import (
	"errors"
)

func tax(s float64) (float64, error) {
	if s < 0 {
		return 0, errors.New("invalid negative salary")
	}

	if s <= 50000 {
		return 0, nil
	} else if s > 50000 && s <= 150000 {
		return 0.17, nil
	} else {
		return 0.27, nil
	}
}
