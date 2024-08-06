package main

import "fmt"

// Canal com buffer ele só vai bloquear esse tráfego de dados quando ele atingir a capacidade máxima
// E o canal sem buffer ele sempre vai bloquear quando estiver uma operação de envio e de recebemento
func main() {
	canal := make(chan string, 2)

	canal <- "Ola mundo"
	canal <- "Programando em GO"


	mensagem := <-canal
	fmt.Println(mensagem)
}