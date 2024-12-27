package main

import "fmt"

func f1() {
	fmt.Printf("F1")
}

func f2(p1 string, p2 string) {
	fmt.Println(p1, p2)
}

func f3() string {
	return "F3"
}

// Mesmo tipo
func f4(p1, p2 string) string {
	return "ANY value"
}

func f5() (string, string) {
	return "Value 1", "Value 2"
}

func main() {
	f1()
	f2("Hello", "World")

	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3, r4)

	r51, r52 := f5()
	fmt.Println(r51, r52)
}
