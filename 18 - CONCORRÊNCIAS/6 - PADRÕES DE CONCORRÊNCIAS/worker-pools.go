package main

import "fmt"

// Você usa o workerpool quando vc tiver uma fila de tarefas grandes pra ser executadas
// Pode fazer com que cada uma tenham varios processor, cada uma pegando uma item dessa fila e executando de forma independente

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// Usando a concorrencia, nós temos 2 processos fazendo a mesma tarefa ao mesmo tempo, entao fica bem mais rapida
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)


	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

// Colocando a seta antes do chan significa que vai ser um canal que só recebe dados
// Colocando a seta depois do chan significa que vai ser um canal que só vai enviar dados
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)

}