package main

import "fmt"

func printAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")

	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Lucas", "Gabriel"}
	printAprovados(aprovados...)
}
