package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Int
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... unit8 unit16 unit32 unit64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2) // 97 runa 'a'

	// boolean
	bo := true
	fmt.Println("O tipo bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Lorem Lorem text"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string múltiplas linhas
	s2 := `
		Hello
		World!
	`
	fmt.Println(s2)
}
