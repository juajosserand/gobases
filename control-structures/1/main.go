package main

import (
	"fmt"
)

func main() {
	meses := []string{"Enero", "Febrero", "Junio", "Agosto", "Abril"}
	// ObtenerEstacion(meses)
	obtenerEstaciones(meses)
}

func ObtenerEstacion(meses []string) {
	for _, mes := range meses {
		if mes == "Enero" || mes == "Febrero" || mes == "Marzo" {
			fmt.Println("Verano")
		} else if mes == "Abril" || mes == "Mayo" || mes == "Junio" {
			fmt.Println("Otoño")
		} else if mes == "Julio" || mes == "Agosto" || mes == "Septiembre" {
			fmt.Println("Invierno")
		} else if mes == "Octubre" || mes == "Noviembre" || mes == "Diciembre" {
			fmt.Println("Primavera")
		} else {
			fmt.Println("No existe este mes")
		}
	}
}

func obtenerEstaciones(months []string) {
	for _, m := range months {
		switch m {
		case "Enero", "Febrero", "Marzo":
			fmt.Println("Verano")
		case "Abril", "Mayo", "Junio":
			fmt.Println("Otoño")
		case "Julio", "Agosto", "Septiembre":
			fmt.Println("Invierno")
		case "Octubre", "Noviembre", "Diciembre":
			fmt.Println("Primavera")
		}
	}
}
