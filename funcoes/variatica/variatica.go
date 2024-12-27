package main

import "fmt"

func media(numbers ...float64) float64 {
	total := 0.0

	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers))
}

func main() {
	fmt.Println(media(1.4, 6.4, 3.6))
}
