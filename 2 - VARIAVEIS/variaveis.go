package main

import "fmt"

func main() {
	// Declarar variável
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// Declarar mais de uma variável junta
	var (
		variavel3 string = "lalala"
		variavel4 string = "lalala"
	)
	
	fmt.Println(variavel3, variavel4)

	// Declarar variáveis na mesma linha
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)
	
	// Criar uma constante 
	const constante1 string = "Constante"
	fmt.Println(constante1)

	// Trocar o valor de uma variável pela outra
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}