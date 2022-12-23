package main

import "fmt"

func main() {
	checkLoan(24, 12, true, 100001) // expected "Loan rejected"
}

// employedSince is represented in months
func checkLoan(age, employedSince int, employed bool, salary float64) {
	if age <= 22 || !employed || employedSince < 12 {
		fmt.Println("Loan rejected")
		return
	}

	if salary > 100000 {
		fmt.Println("Loan accepted without interests")
	} else {
		fmt.Println("Loan accepted with interests")
	}
}
