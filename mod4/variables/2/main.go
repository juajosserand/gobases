package main

import "fmt"

func main() {
	var (
		t float32 = 26.4
		h float32 = 0.5   // percentage, range: 0-1
		p float64 = 1.009 // unit: mBar
	)

	fmt.Println("Temperature:", t)
	fmt.Println("Humidity:", h)
	fmt.Println("Temperature:", p)
}
