package main

import "fmt"

func getResult(nota float64) string {
	if nota > 6 {
		return "Approved"
	}

	return "Reproved"
}

func main() {
	fmt.Println(getResult(6.2))
}
