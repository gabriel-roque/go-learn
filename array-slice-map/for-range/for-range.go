package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador define o tamanho do array

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}
}
