package main

import (
	"fmt"
	"time"
)

func main() {
	e := Employee{
		Person: Person{
			ID:          1,
			Name:        "Steve",
			DateOfBirth: time.Now(),
		},
		ID:       2,
		Position: "Manager",
	}

	e.PrintEmployee()
}

type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

type Employee struct {
	Person
	ID       int
	Position string
}

func (e *Employee) PrintEmployee() {
	fmt.Println(e)
}
