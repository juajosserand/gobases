package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	Name               string
	InputCategory      string
	InputWorkedMinutes float64
	Expected           float64
}

func TestSalary(t *testing.T) {
	tests := []Test{
		{"CategoryC", "C", 60, 1000},
		{"CategoryB", "B", 60, 1800},
		{"CategoryA", "A", 60, 4500},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			s, err := salary(test.InputWorkedMinutes, test.InputCategory)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, test.Expected, s)
		})
	}
}

func TestSalaryErrorInvalidWorkedMinutes(t *testing.T) {
	var workedMin float64 = -1
	var category string = "C"
	var expected float64 = 0

	s, err := salary(workedMin, category)
	assert.NotNil(t, err)
	assert.Equal(t, expected, s)
}

func TestSalaryErrorInvalidCategory(t *testing.T) {
	var workedMin float64 = 60
	var category string = "D"
	var expected float64 = 0

	s, err := salary(workedMin, category)
	assert.NotNil(t, err)
	assert.Equal(t, expected, s)
}
