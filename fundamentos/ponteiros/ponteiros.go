package main

import "fmt"

func main() {
	i := 1

	// GO não tem aritmética de ponteiros
	var p *int = nil

	p = &i // pegando o endereço da variável na memória
	*p++   // desreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i, &i)
}
