package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct{
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Pedro", 170, 20}
	fmt.Println(p1)
	est1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(est1)
}