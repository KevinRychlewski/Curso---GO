package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Ola mundo", canal)
	// quando usamos <- canal ele vai receber o valor do canal


	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim")

}


func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {	
		//quando usamos canal <- ele vai enviar o valor para o canal
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}



