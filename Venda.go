package main

type Venda struct {
	numero  int
	data    string
	cliente Cliente
	itens   []ItemVenda
}

func (venda Venda) total() float32 {
	var SomaTotal float32 = 0.0
	for _, item := range venda.itens {
		SomaTotal += item.total()
	}
	return SomaTotal
}
