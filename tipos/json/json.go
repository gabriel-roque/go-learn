package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	Id    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// STRUCT -> JSON

	p1 := produto{
		Id:    1,
		Nome:  "Notebook",
		Preco: 1899.90,
		Tags:  []string{"Promoção", "Eletrônico"},
	}

	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// JSON -> STRUCT
	var p2 produto
	// jsonString := `{"id":1,"nome":"Notebook","preco":1899.9,"tags":["Promoção","Eletrônico"]}`
	json.Unmarshal(p1Json, &p2)
	fmt.Println(p2)
}
