package main

import "fmt"

func main() {
	fmt.Println("var 1nombre string -> invalid")
	fmt.Println("var apellido string -> valid")
	fmt.Println("var edad int -> valid")
	fmt.Println("1apellido := 6 -> invalid")
	fmt.Println("var licencia_de_conducir = true -> valid")
	fmt.Println("var estatura de la persona int -> invalid")
	fmt.Println("cantidadDeHijos := 2 -> valid")

	fmt.Println()
	fmt.Println("Correct:")
	fmt.Println("var nombre string")
	fmt.Println("var apellido string")
	fmt.Println("var edad int")
	fmt.Println("apellido := <string>")
	fmt.Println("var licenciaDeConducir = true")
	fmt.Println("var estaturaDeLaPersona int")
	fmt.Println("cantidadDeHijos := 2")
}
