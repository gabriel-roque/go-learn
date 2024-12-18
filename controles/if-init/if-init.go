package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	for i := 0; i < 10; i++ {
		if i := randomNumber(); i > 5 {
			fmt.Println("Winner!")
		} else {
			fmt.Println("Loser!")
		}
	}
}
