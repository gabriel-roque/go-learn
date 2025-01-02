package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	parts := strings.Split(nomeCompleto, " ")
	p.nome = parts[0]
	p.sobrenome = parts[1]
}

func main() {
	p1 := pessoa{nome: "Gabriel", sobrenome: "Roque"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Rebeca Santos")
	fmt.Println(p1.getNomeCompleto())
}
