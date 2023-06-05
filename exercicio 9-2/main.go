package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Printf("Meu nome é %v e eu tenho %v anos \n", p.nome, p.idade)
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	fmt.Println("Esse aqui é a interface")
	h.falar()
}

func main() {
	pessoa1 := pessoa{"Wesley", 20}

	pessoa1.falar()
	dizerAlgumaCoisa(&pessoa1)
}
