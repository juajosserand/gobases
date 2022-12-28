package main

import (
	"fmt"

	"github.com/juajosserand/gobases/error-handling/3/model"
	"github.com/juajosserand/gobases/error-handling/3/storage"
)

func main() {
	defer func() {
		fmt.Println("execution finished")

		if err := recover(); err != nil {
			fmt.Println("errors were detected at runtime")
		}
	}()

	c1 := model.Customer{
		ID:          1,
		Name:        "Customer 1",
		DNI:         "12.345.678",
		PhoneNumber: "0123-456789",
		Address:     "Customer 1 Address",
	}

	var err error
	s := storage.NewCustomerStorage()

	// adding customer, no error expected
	s.Add(c1)

	// expecting customer validation error
	c1.ID = 0 // invalid id
	err = s.Add(c1)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	// expecting existing customer error
	c1.ID = 1
	err = s.Add(c1)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
