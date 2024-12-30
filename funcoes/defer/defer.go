package main

import "fmt"

// defer - executa sempre até o ultimo momento possível da função

func getResult(value int) int {
	defer fmt.Println("Fim!")

	if value > 5000 {
		fmt.Println("Valor alto!")
		return 5000
	}

	fmt.Println("Valor baixo!")
	return value
}

func main() {
	fmt.Println(getResult(6000))
	fmt.Println(getResult(3000))
}
