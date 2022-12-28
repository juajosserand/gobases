package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}

		fmt.Println("execution finished")
	}()

	_, err := os.ReadFile("customers.txt")
	if err != nil {
		panic("file not found or damaged")
	}
}
