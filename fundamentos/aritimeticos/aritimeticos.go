package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b)

	// bitwise (bit a bit)
	fmt.Println("E =>", a&b)
	fmt.Println("OU =>", a|b)
	fmt.Println("XOR =>", a^b)

	c := 3.2
	d := 2.5

	// outras operações usando math
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(float64(a), float64(b)))
	fmt.Println("Potência =>", math.Pow(c, d))
}
