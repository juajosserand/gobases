package main

import (
	"errors"
)

func animal(a string) (func(uint) float64, error) {
	switch {
	case a == "dog":
		return dog, nil
	case a == "cat":
		return cat, nil
	case a == "hamster":
		return hamster, nil
	case a == "tarantula":
		return tarantula, nil
	default:
		return nil, errors.New("invalid animal identifier")
	}
}

func dog(n uint) float64 {
	return 10 * float64(n)
}

func cat(n uint) float64 {
	return 5 * float64(n)
}

func hamster(n uint) float64 {
	return 250 * float64(n)
}

func tarantula(n uint) float64 {
	return 150 * float64(n)
}
