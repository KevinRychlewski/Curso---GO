package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		// Usando o select eu deixo de perder o tempo de sincronização entre o canal 1 e 2, por conta que com o select, ele apenas envia os dados entre o canal quando eles estão prontos
		// Então se o canal 1 receber 4 valor em 2 segundos e o canal 2 receber apenas 1, o canal 1 vai liberar 4 valores e depois o canal 2 vai liberar 1
		// Caso contrário, o canal 1 ficaria na dependencia do canal 2 pra enviar mais dados, e consequentemente o canal 2 ficaria travando o envio de dados do canal 1, fazendo ele perder tempo esperando
		
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}