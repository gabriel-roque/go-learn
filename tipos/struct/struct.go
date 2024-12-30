package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func (p produto) format() {
	fmt.Printf("%s no preço de R$ %2f\n", p.nome, p.precoComDesconto())
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Caneta",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(produto1, produto1.precoComDesconto())
	produto1.format()

	produto2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println(produto2, produto2.precoComDesconto())
}
