package main

import "fmt"

type item struct {
	produtoID int
	qtd       int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{1, 20, 2.99},
			{2, 1, 29.99},
			{3, 5, 1.99},
		},
	}
	fmt.Printf("Valor total do pedido é de R$%.2f", pedido.valorTotal())
}
