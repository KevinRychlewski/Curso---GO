package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Ola mundo"), escrever("Programando em GO"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// multiplexar recebe dois canais de entrada e retorna um único canal de saída
// este canal de saída irá receber mensagens de ambos os canais de entrada
// a seleção é feita usando o comando select
func multiplexar(canalDeEntrada1, canalDeEntrada2 <- chan string) <- chan string {
	// canal de saída que irá receber as mensagens de ambos os canais de entrada
	canalDeSaida := make(chan string)

	// goroutine que irá ler as mensagens de ambos os canais de entrada e
	// enviar para o canal de saída
	go func() {
		for {
			// seleção entre os dois canais de entrada
			select{
				// caso uma mensagem seja recebida do canal de entrada 1
				case mensagem := <- canalDeEntrada1:
					// a mensagem é enviada para o canal de saída
					canalDeSaida <- mensagem
				// caso uma mensagem seja recebida do canal de entrada 2
				case mensagem := <- canalDeEntrada2:
					// a mensagem é enviada para o canal de saída
					canalDeSaida <- mensagem
			}
		}
	}()
		return canalDeSaida
}

func escrever(texto string) <- chan string {
	canal := make(chan string)

	go func(){
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 200)
		}
	}()

	return canal
}