package main

type ItemVenda struct {
	produto Produto
	valor   float32
	qtd     int
}

func (item ItemVenda) total() float32 {
	return item.valor * float32(item.qtd)
}
