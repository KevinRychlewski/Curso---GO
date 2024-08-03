package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	go escrever("Ola Mundo") // goroutine = ele continua rodando o código simultaneamente de forma paralela
	escrever("Programando em GO")
}

func escrever(texto string) {
	for{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
