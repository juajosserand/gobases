package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 13; i++ {
		printMonth(time.Month(i))
	}
}

func printMonth(m time.Month) {
	if m < 1 || m > 12 {
		fmt.Println("invalid month number")
		return
	}

	fmt.Println("month is", m.String())
}
