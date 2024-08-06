package main

import "fmt"

// Canal com buffer ele só vai bloquear esse tráfego de dados quando ele atingir a capacidade máxima
// E o canal sem buffer ele sempre vai bloquear quando estiver uma operação de envio e de recebemento
// A principal diferença entre um canal com buffer e um sem buffer é a capacidade de armazenamento temporário de valores. Um canal com buffer permite que um número limitado de valores seja armazenado temporariamente, enquanto um canal sem buffer bloqueia a comunicação até que um receptor esteja pronto para receber o valor.
func main() {
	canal := make(chan string, 2)

	canal <- "Ola mundo"
	canal <- "Programando em GO"


	mensagem := <-canal
	fmt.Println(mensagem)
}