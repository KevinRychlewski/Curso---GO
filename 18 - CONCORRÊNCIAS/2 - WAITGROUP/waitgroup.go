package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	// especifica a quantidade de goroutines que irão ser executadas
	waitGroup.Add(2)

	go func() {
		escrever("Ola Mundo")
		// waitGroup.Done() indica que a goroutine foi concluída e tira 1 do waitGroup
		waitGroup.Done()
	}()

	go func(){
		escrever("Programando em GO")
		waitGroup.Done()
	}()
	
	// waitGroup.Wait() fala pra função main aguardar até todas as goroutines serem concluídas
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {	
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
