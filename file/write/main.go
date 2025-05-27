package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

var maleNames = []string{"João", "Pedro", "Lucas", "Carlos", "Mateus"}
var femaleNames = []string{"Maria", "Ana", "Beatriz", "Clara", "Juliana"}
var genders = []string{"male", "female"}

func generateRandomPerson() (string, string, int) {
	gender := genders[rand.Intn(len(genders))]
	var name string
	if gender == "male" {
		name = maleNames[rand.Intn(len(maleNames))]
	} else {
		name = femaleNames[rand.Intn(len(femaleNames))]
	}

	age := rand.Intn(82) + 18
	return name, gender, age
}

func worker(id int, jobs <-chan int, results chan<- []string, wg *sync.WaitGroup) {
	defer wg.Done()

	for range jobs {
		name, gender, age := generateRandomPerson()
		results <- []string{name, gender, strconv.Itoa(age)}
	}
}

func main() {
	file, err := os.Create("random_people.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"name", "gender", "age"})

	numRecords := 1000 * 1000
	numWorkers := 1

	jobs := make(chan int, numRecords)
	results := make(chan []string, numRecords)

	var wg sync.WaitGroup

	for w := range numWorkers {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := range numRecords {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for record := range results {
		writer.Write(record)
	}

	fmt.Println("Finalizado com sucesso.")
}
