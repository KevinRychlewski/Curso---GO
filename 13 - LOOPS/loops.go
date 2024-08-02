package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)
	
	for j := 0; j < 10; j+=2 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	
	// FOR - RANGE
	nomes := [3]string{"João", "Maria", "Pedro"}
	
	for indice, nomes := range nomes{
		fmt.Println(indice, nomes)
	}


	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string {
		"nome": "Leonardo",
		"sobrenome": "Silva",

	}

	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}