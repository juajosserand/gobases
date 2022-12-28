package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/juajosserand/gobases/error-handling/3/model"
)

var (
	errCustomerExists = errors.New("customer already exists")
)

type CustomerStorage interface {
	Add(model.Customer) error
}

type storage struct {
	filename  string
	Customers []model.Customer `json:"customers"`
}

func NewCustomerStorage() (CustomerStorage, error) {
	s := &storage{
		filename: "./customers.json",
	}

	err := s.readFromFile()
	if err != nil {
		return nil, err
	}

	return s, nil
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

	s.Customers = append(s.Customers, c)

	err = s.writeToFile()
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) exists(c model.Customer) error {
	for _, customer := range s.Customers {
		if customer.ID == c.ID {
			return fmt.Errorf("[storage.exists] error: %w", errCustomerExists)
		}
	}

	return nil
}

func (s *storage) readFromFile() error {
	data, err := os.ReadFile(s.filename)
	if err != nil {
		return fmt.Errorf("[storage.readFromFile] error: %w", err)
	}

	err = json.Unmarshal(data, &s)
	if err != nil {
		return fmt.Errorf("[storage.readFromFile] error: %w", err)
	}

	return nil
}

func (s *storage) writeToFile() error {
	bytes, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("[storage.writeToFile] error: %w", err)
	}

	err = os.WriteFile(s.filename, bytes, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("[storage.writeToFile] error: %w", err)
	}

	return nil
}
