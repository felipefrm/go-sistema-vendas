package model

type Venda struct {
	Numero  int
	Data    string
	Cliente Cliente
	Itens   []ItemVenda
}

func (venda Venda) total() float32 {
	var SomaTotal float32 = 0.0
	for _, item := range venda.Itens {
		SomaTotal += item.total()
	}
	return SomaTotal
}
