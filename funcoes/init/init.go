package main

import "fmt"

// Função init e executada sempre no inicio de cada arquivo GO
func init() {
	fmt.Println("Init!")
}

func main() {
	fmt.Println("Main!")
}
