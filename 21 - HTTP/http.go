package main

import (
	"log"
	"net/http"
)

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE (FAZ REQUISIÇÃO)- SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)

	// Resquest - Response

	// Rotas
	// URI - Identidicador do Recurso
	// Método - GET, POST, PUT, DELETE,

	// COMO CRIAR UMA ROTA(Primeiro nós passamos o URI(/home) e o segundo parametro é uma função que vai receber o /home e vai saber como lidar com ele) 
	// O HTTP.RESPONSEWRITER E O *HTTP.REQUEST SÃO OS QUE VAO FAZER A RESPONSE E A REQUEST
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Ola mundo"))
	})

	http.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("ola usuario"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}