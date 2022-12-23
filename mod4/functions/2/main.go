package main

import (
	"errors"
)

func avg(scores ...int) (result float64, err error) {
	for _, s := range scores {
		if s < 0 {
			return 0, errors.New("invalid negative score")
		}

		result += float64(s) / float64(len(scores))
	}

	return result, nil
}
