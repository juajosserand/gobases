package main

import "errors"

func Operation(op func(...int) float64, values ...int) (float64, error) {
	if op != nil {
		return op(values...), nil
	}

	return 0, errors.New("invalid operation")
}

func min(values ...int) (min float64) {
	if len(values) > 0 {
		min = float64(values[0])
	}

	for _, v := range values {
		if float64(v) < min {
			min = float64(v)
		}
	}

	return min
}

func max(values ...int) (max float64) {
	if len(values) > 0 {
		max = float64(values[0])
	}

	for _, v := range values {
		if float64(v) > max {
			max = float64(v)
		}
	}

	return max
}

func avg(values ...int) (avg float64) {
	for _, v := range values {
		avg += float64(v) / float64(len(values))
	}

	return avg
}
