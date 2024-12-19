package main

import "fmt"

func main() {
	employeeSalary := map[string]float64{
		"José João":     11325.45,
		"Gabriel Silva": 15456.78,
		"Pedro Junior":  1200.0,
	}

	employeeSalary["Rafael Filho"] = 1350.0
	delete(employeeSalary, "Any Employee")

	for employee, salary := range employeeSalary {
		fmt.Println(employee, salary)
	}
}
