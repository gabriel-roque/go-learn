package main

import "fmt"

func main() {
	count := 10

	i := 1

	for i <= count {
		fmt.Println(i)
		i++
	}

	for i := 0; i < count; i++ {
		fmt.Println(i + 1)
	}

	for i := 0; i < count; i += 2 {
		fmt.Println(i + 1)
	}
}
