package model

type ItemVenda struct {
	Produto Produto
	Valor   float32
	Qtd     int
}

func (item ItemVenda) total() float32 {
	return item.Valor * float32(item.Qtd)
}
