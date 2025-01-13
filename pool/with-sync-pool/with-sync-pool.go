package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	fileSize   = 5 * 1024 * 1024 * 1024 // 1 GB
	bufferSize = 4 * 1024               // 4 KB
)

func main() {
	var pool = sync.Pool{
		New: func() interface{} {
			return make([]byte, bufferSize)
		},
	}

	file, err := os.Create("file_with_pool.bin")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	// Monitoramento inicial
	startTime := time.Now()
	var memStatsBefore, memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)

	var totalWritten int64
	for totalWritten < fileSize {
		buffer := pool.Get().([]byte)

		_, err := rand.Read(buffer)
		if err != nil {
			fmt.Println("Erro ao preencher o buffer:", err)
			return
		}

		n, err := file.Write(buffer)
		if err != nil {
			fmt.Println("Erro ao escrever no arquivo:", err)
			return
		}

		totalWritten += int64(n)
		pool.Put(buffer)
	}

	// Monitoramento final
	runtime.ReadMemStats(&memStatsAfter)
	elapsed := time.Since(startTime)

	// Exibindo resultados
	fmt.Println("Arquivo criado com sucesso usando sync.Pool!")
	fmt.Printf("Tempo de execução: %v\n", elapsed)
	fmt.Printf("Uso de memória antes: %d KB\n", memStatsBefore.Alloc/1024)
	fmt.Printf("Uso de memória após: %d KB\n", memStatsAfter.Alloc/1024)
	fmt.Printf("Total de garbage collections: %d\n", memStatsAfter.NumGC-memStatsBefore.NumGC)
}
