package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[7675] = "Maria"
	aprovados[2454] = "João"
	aprovados[5453] = "Lucas"
	fmt.Println(aprovados)

	for id, nome := range aprovados {
		fmt.Printf("%s (ID: %d)\n", nome, id)
	}

	fmt.Println(aprovados[7675])
	delete(aprovados, 7675)
	fmt.Println(aprovados)
}
