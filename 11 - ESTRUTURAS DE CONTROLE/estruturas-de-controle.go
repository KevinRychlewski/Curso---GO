package main

import "fmt"


func main() {
	fmt.Println("Estruturas de Controle")

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número maior que 0")
	} else if numero < -10{
		fmt.Println("Número menor ou igual -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}
}
