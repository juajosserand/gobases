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
			fmt.Println(err)
		}
	}()

	c1 := model.Customer{
		ID:          1,
		Name:        "Customer 1",
		DNI:         "12.345.678",
		PhoneNumber: "0123-456789",
		Address:     "Customer 1 Address",
	}

	// c2 := model.Customer{
	// 	ID:          2,
	// 	Name:        "Customer 2",
	// 	DNI:         "12.345.679",
	// 	PhoneNumber: "0123-456710",
	// 	Address:     "Customer 2 Address",
	// }

	s, err := storage.NewCustomerStorage()
	if err != nil {
		panic(err)
	}

	// adding customer, no error expected if c1 and c2 not in file
	// err = s.Add(c1)
	// if err != nil {
	// 	panic(err)
	// }

	// err = s.Add(c2)
	// if err != nil {
	// 	panic(err)
	// }

	// expecting customer validation error
	c1.ID = 0
	err = s.Add(c1)
	if err != nil {
		panic(err)
	}

	// expecting existing customer error
	c1.ID = 1
	err = s.Add(c1)
	if err != nil {
		panic(err)
	}
}
