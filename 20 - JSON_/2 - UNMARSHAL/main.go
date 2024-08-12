package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

// Como converter um JSON em struct ou map usando UNSMARSHAL
func main() {
	cachorroEmJson := `{"nome": "Rex", "raca": "Dalmata", "idade": 3}`
	var c cachorro
	// Convertendo JSON para struct 
	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	// Convertendo JSON para map
	cachorro2EmJson := `{"nome": "Toby", "raca": "Poodle"}`

	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro2EmJson), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
