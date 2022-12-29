package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCSVRepository(t *testing.T) {
	test := struct{ input string }{"../../tickets.csv"}

	_, err := NewCSVRepository(test.input)
	assert.Nil(t, err)
}

func TestNewCSVRepository_Err(t *testing.T) {
	test := struct{ input string }{""}

	_, err := NewCSVRepository(test.input)
	assert.NotNil(t, err)
}
