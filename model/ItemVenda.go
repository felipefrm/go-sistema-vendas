package model

type ItemVenda struct {
	Produto *Produto
	Valor   float32
	Qtd     int
}

func (item ItemVenda) Total() float32 {
	return item.Valor * float32(item.Qtd)
}
