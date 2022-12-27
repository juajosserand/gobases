package main

import (
	"fmt"

	"github.com/juajosserand/gobases/interfaces/2/model"
)

func main() {
	p, err := model.NewCoster(model.LargeProduct, 100)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
	fmt.Println(p.Price())
}
