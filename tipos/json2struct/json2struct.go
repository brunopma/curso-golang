package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct 2 json
	p1 := produto{1, "Notebook", 1999.90, []string{"Promocao", "Eletronico"}}
	p1Json, _ := json.MarshalIndent(p1, "", "  ")
	fmt.Println(string(p1Json))

	// json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":8.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}
