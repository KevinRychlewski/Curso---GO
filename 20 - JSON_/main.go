package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string	`json:"raca"`
	Idade uint `json:"idade"`
}

// Marshal para converter um map em JSON ou um struct em JSON
// json.Marshal()
// // Unmarshal para converter JSON em um map ou struct
// json.Unmarshal()


func main() {
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)
	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJson)

	// Para converter um slice de bytes em JSON
	fmt.Println(bytes.NewBuffer(cachorroEmJson))

	// Exemplo de conversão de um map em JSON
	c2 := map[string]string {
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJson, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJson)
	fmt.Println(bytes.NewBuffer(cachorro2EmJson))
}