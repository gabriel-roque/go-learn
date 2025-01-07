package main

import (
	"fmt"
)

func loop(total int, c chan int) {
	for i := 0; i < total; i++ {
		c <- i + 1
		// time.Sleep(time.Second / 2)
	}
}

func main() {
	c := make(chan int)
	total := 1000

	go loop(total, c)

	for i := 0; i < total; i++ {
		fmt.Println(<-c)
	}
}
