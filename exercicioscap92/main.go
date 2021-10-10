package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p *Pessoa) Falar() {
	fmt.Printf("Olá, meu nome é %s e tenho %v anos de vida\n\n", p.Nome, p.Idade)
}

type Humanos interface {
	Falar()
}

func DigaAlgo(coisaQuefala Humanos) {
	coisaQuefala.Falar()
}

func main() {
	felipe := Pessoa{
		Nome:  "Felipe Alafy",
		Idade: 15,
	}

	denis := Pessoa{
		Nome:  "Denis Alexandre",
		Idade: 16,
	}

	DigaAlgo(&felipe)
	DigaAlgo(&denis)
}
