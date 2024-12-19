package main

import "fmt"

func main() {
	employeeByLetter := map[string]map[string]float64{
		"G": {
			"Gabriel Silva": 14234.33,
			"Lucas Marcos":  14234.33,
		},
		"J": {
			"Júlia Narcos": 1231.4,
		},
		"P": {
			"Pedro Paulo": 13223.32,
		},
	}

	delete(employeeByLetter, "P")

	for letter, group := range employeeByLetter {
		for employee, salary := range group {
			fmt.Println(letter, employee, salary)
		}
	}
}
