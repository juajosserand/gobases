package main

import "fmt"

func main() {
	employees := map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	// printing benjamin's age
	fmt.Println("benjamin's age:", employees["Benjamin"])

	// employees older than 21
	var count int
	for _, age := range employees {
		if age > 21 {
			count++
		}
	}
	fmt.Println("employees older than 21:", count)

	// adding federico (25) to list
	employees["Federico"] = 25
	fmt.Println(employees)

	// removing Pedro from list
	delete(employees, "Pedro")
	fmt.Println(employees)
}
