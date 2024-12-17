package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3
	i += 3
	i -= 3
	i /= 3
	i *= 3
	i %= 3

	fmt.Println(i)

	// Múltipla atribuição
	x, y := 1, 2
	fmt.Println(x, y)

	// Inverter
	x, y = y, x
	fmt.Println(x, y)
}
