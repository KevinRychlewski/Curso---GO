package main

import (
	"fmt"
	"teste-automatizados/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("praça dos bobos")
	fmt.Println(tipoEndereco)
}