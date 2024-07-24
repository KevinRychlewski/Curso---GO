package main

import "fmt"

func main() {

	usuario := map[string]string {
		"nome": "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string { // cria um map dentro de outro
		"nome": {
			"primeiro": "João",
			"ultimo": "Pedro",
		},
		"curso": {
			"nome": "Engenharia",
			"faculdade": "USP",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2,"nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)
}
