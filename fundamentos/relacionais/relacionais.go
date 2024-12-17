package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("String", "banana" == "banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{Nome: "João"}
	p2 := Pessoa{Nome: "João"}

	fmt.Println("Pessoas:", p1 == p2)
}
