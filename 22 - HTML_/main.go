package main

import (
	"log"
	"net/http"
	"html/template"
)

// A GENTE CRIA UMA VÁRIAVEL PARA ARMAZENAR TODOS OS ARQUIVOS DE TEMPLATE QUE A GENTE TEM DENTRO DA NOSSA APLICAÇÃO
var templates *template.Template

type usuario struct {
	Nome string
	Email string
}

func main() {
	// A GENTE JOGA NA VARIAVEL TEMPLATES TODOS OS NOSSOS TEMPLATES QUE A GENTE VAI CRIAR
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request){
		// CRIANDO UM USUARIO
		u := usuario{"Davi", "davi.felipe@gmail.com"}

		templates.ExecuteTemplate(w, "home.html", u)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}