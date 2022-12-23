package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestAnimalFunc struct {
	Name   string
	Animal string
}

type TestAnimal struct {
	Input    uint
	Expected float64
}

func TestAnimalFuncs(t *testing.T) {
	tests := []TestAnimalFunc{
		{"TestDogFunc", "dog"},
		{"TestCatFunc", "cat"},
		{"TestHamsterFunc", "hamster"},
		{"TestTarantulaFunc", "tarantula"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			f, err := animal(test.Animal)
			if err != nil {
				t.Fatal(err)
			}

			assert.NotNil(t, f)
		})
	}
}

func TestAnimalFuncsErrorInvalid(t *testing.T) {
	test := TestAnimalFunc{"TestDogFunc", "dragon"}

	f, err := animal(test.Animal)
	assert.NotNil(t, err)
	assert.Nil(t, f)
}

func TestDog(t *testing.T) {
	test := TestAnimal{1, 10}

	r := dog(test.Input)

	assert.Equal(t, test.Expected, r)
}

func TestCat(t *testing.T) {
	test := TestAnimal{1, 5}

	r := cat(test.Input)

	assert.Equal(t, test.Expected, r)
}

func TestHamster(t *testing.T) {
	test := TestAnimal{1, 250}

	r := hamster(test.Input)

	assert.Equal(t, test.Expected, r)
}

func TestTarantula(t *testing.T) {
	test := TestAnimal{1, 150}

	r := tarantula(test.Input)

	assert.Equal(t, test.Expected, r)
}
