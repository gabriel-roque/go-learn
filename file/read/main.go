package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Person struct {
	Name   string
	Gender string
	Age    int
}

func worker(id int, lines <-chan []string, people chan<- Person, wg *sync.WaitGroup) {
	defer wg.Done()

	for line := range lines {
		if len(line) != 3 {
			continue
		}

		age, err := strconv.Atoi(line[2])
		if err != nil {
			continue
		}

		p := Person{
			Name:   line[0],
			Gender: line[1],
			Age:    age,
		}

		people <- p
	}
}

func main() {
	file, err := os.Open("random_people.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	lines := make(chan []string, 100)
	people := make(chan Person, 100)

	const numWorkers = 8
	var wg sync.WaitGroup

	for i := range numWorkers {
		wg.Add(1)
		go worker(i, lines, people, &wg)
	}

	go func() {
		for {
			record, err := reader.Read()
			if err != nil {
				break
			}
			lines <- record
		}
		close(lines)
	}()

	go func() {
		wg.Wait()
		close(people)
	}()

	count := 0
	for p := range people {
		fmt.Printf("Nome: %-10s | Gênero: %-6s | Idade: %d\n", p.Name, p.Gender, p.Age)
		count++
	}

	fmt.Printf("Total de pessoas processadas: %d\n", count)
}
