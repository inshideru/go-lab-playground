package main

import (
	"fmt"
)

// Estado descreve um estado com nome, população e capital
type Estado struct {
	nome      string
	populacao int
	capital   string
}

func main() {
	estados := make(map[string]Estado, 6)

	estados["GO"] = Estado{"Goiás", 6434052, "Goiânia"}
	estados["PB"] = Estado{"Paraíba", 3914418, "João Pessoa"}
	estados["PR"] = Estado{"Paraná", 10997462, "Curitiba"}
	estados["RN"] = Estado{"Rio Grande do Norte", 3373960, "Natal"}
	estados["AM"] = Estado{"Amazonas", 3807923, "Manaus"}
	estados["SE"] = Estado{"Sergipe", 2228489, "Aracaju"}
	estados["SP"] = Estado{"São Paulo", 2228489, "São Paulo"}

	fmt.Println(estados)

	saoPaulo, encontrado := estados["SP"]
	if encontrado {
		fmt.Println(saoPaulo)
	}

	delete(estados, "AM")

	for sigla, estado := range estados {
		fmt.Printf("%s (%s) possui %d habitantes.\n",
			estado.nome, sigla, estado.populacao)
	}

}
