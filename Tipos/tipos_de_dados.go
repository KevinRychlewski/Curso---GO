package main

import (
	"errors"
	"fmt"
)

func main() {

	// Tipos de números inteiros
	// int8, int16, int32 e int64

	var numero int64 = 10000000
	fmt.Println(numero)

	// uint / unsygned int

	var numero2 uint32 = 1000000
	fmt.Println(numero2)

	// alias

	// RUNE = INT32
	var numero3 rune = 12345
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// Tipos de números reais
	// float32 e float 64

	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)
	
	var numeroReal2 float64 = 1230000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// Tipo String

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// Char

	char := 'B'
	fmt.Println(char)

	// Valor 0

	var texto string
	fmt.Println(texto)

	var numero5 int
	fmt.Println(numero5)

	// Tipo Booleano

	var booleano1 bool = true
	fmt.Println(booleano1)
	
	var booleano2 bool = false
	fmt.Println(booleano2)
	
	// Tipo Erro

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}