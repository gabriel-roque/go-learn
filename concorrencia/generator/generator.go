package generator

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// Google I/O 2012 - Go Concurrency Patterns

// <-chan - canal somente leitura
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				c <- fmt.Sprintf("Erro ao acessar %s: %v", url, err)
				return
			}
			defer resp.Body.Close()

			html, err := io.ReadAll(resp.Body)
			if err != nil {
				c <- fmt.Sprintf("Erro ao ler corpo da resposta de %s: %v", url, err)
				return
			}

			r, _ := regexp.Compile("<title>(.*?)</title>")
			match := r.FindStringSubmatch(string(html))
			if len(match) > 1 {
				c <- fmt.Sprintf("Título de %s: %s", url, match[1])
			} else {
				c <- fmt.Sprintf("Título não encontrado em %s", url)
			}
		}(url)
	}

	return c
}

func main() {
	t1 := Titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := Titulo("https://www.amazon.com", "https://www.youtube.com")
	fmt.Println(<-t1)
	fmt.Println(<-t1)
	fmt.Println(<-t2)
	fmt.Println(<-t2)
}
