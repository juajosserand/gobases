package model

import (
	"errors"
	"fmt"
)

var (
	errCustomerValidation = errors.New("customer has zero-value fields")
)

type Customer struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	DNI         string `json:"dni"`
	PhoneNumber string `json:"phone_numbre"`
	Address     string `json:"address"`
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
