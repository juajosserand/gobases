package model

import (
	"errors"
	"fmt"
)

var (
	errCustomerValidation = errors.New("customer has zero-value fields")
)

type Customer struct {
	ID          uint
	Name        string
	DNI         string
	PhoneNumber string
	Address     string
}

func (c *Customer) Validate() error {
	if ok := c.valid(); !ok {
		return fmt.Errorf("[Customer.Validate] error: %w", errCustomerValidation)
	}

	return nil
}

func (c Customer) valid() bool {
	return c.ID != 0 &&
		c.Name != "" &&
		c.DNI != "" &&
		c.PhoneNumber != "" &&
		c.Address != ""
}
