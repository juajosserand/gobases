package main

import (
	"fmt"
	"time"
)

func main() {
	s := &Student{
		Firstname: "Steve",
		Lastname:  "Gopher",
		DNI:       "12.345.678",
		Date:      time.Date(2017, time.December, 15, 0, 0, 0, 0, time.UTC),
	}

	print(s)
}

type Detailer interface {
	Detail() string
}

func print(d Detailer) {
	fmt.Println(d.Detail())
}

type Student struct {
	Firstname string
	Lastname  string
	DNI       string
	Date      time.Time
}

func (s *Student) Detail() string {
	return fmt.Sprintf(
		"Firstname: %s\nLastname: %s\nDNI: %s\nAdmission Date: %v",
		s.Firstname,
		s.Lastname,
		s.DNI,
		s.Date,
	)
}
