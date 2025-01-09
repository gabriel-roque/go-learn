package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d\n", pessoa, i)
		}
		close(c) // Fechar o channel após enviar todas as mensagens
	}()

	return c
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		defer close(c) // Garantir que o channel `c` será fechado ao final

		for entrada1 != nil || entrada2 != nil {
			select {
			case s, ok := <-entrada1:
				if !ok {
					entrada1 = nil // Finalizar leitura de entrada1
					continue
				}

				c <- s
			case s, ok := <-entrada2:
				if !ok {
					entrada2 = nil // Finalizar leitura de entrada2
					continue
				}

				c <- s
			}
		}

	}()

	return c
}

func main() {
	c := juntar(falar("João"), falar("Maria"))

	for msg := range c {
		fmt.Printf(msg)
	}

	fmt.Println("Leitura concluída!")
}
