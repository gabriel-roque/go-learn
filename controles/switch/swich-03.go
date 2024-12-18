package main

import (
	"fmt"
	"time"
)

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Não definido"
	}
}

func main() {
	fmt.Println(getType(1))
	fmt.Println(getType(1.2))
	fmt.Println(getType("Hello!"))
	fmt.Println(getType(func() {}))
	fmt.Println(getType(time.Now()))
}
