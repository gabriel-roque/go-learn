package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		// time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Bom dia!", 3)
	// fale("Maria", "Só posso falar depois de você", 1)

	go fale("Maria", "Ei...", 500)
	go fale("João", "Opa...", 500)

	time.Sleep(time.Second * 5)
}
