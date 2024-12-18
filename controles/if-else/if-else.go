package main

import "fmt"

func printNote(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}

}

func main() {
	printNote(7.3)
	printNote(5.1)
}
