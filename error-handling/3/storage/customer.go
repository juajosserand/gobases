package storage

import (
	"errors"
	"fmt"

	"github.com/juajosserand/gobases/error-handling/3/model"
)

var (
	errCustomerExists = errors.New("customer already exists")
)

type CustomerStorage interface {
	Add(model.Customer) error
}

type storage struct {
	customers []model.Customer
}

func NewCustomerStorage() CustomerStorage {
	return &storage{
		customers: []model.Customer{},
	}
}

// Not concurrent safe!
func (s *storage) Add(c model.Customer) error {
	err := c.Validate()
	if err != nil {
		return err
	}

	err = s.exists(c)
	if err != nil {
		return err
	}

	s.customers = append(s.customers, c)
	return nil
}

func (s *storage) exists(c model.Customer) error {
	for _, customer := range s.customers {
		if customer.ID == c.ID {
			return fmt.Errorf("[storage.exists] error: %w", errCustomerExists)
		}
	}

	return nil
}
