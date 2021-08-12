package view

import (
	"fmt"
	"time"
)

type VendaView struct {
	clienteview ClienteView
	produtoview ProdutoView
}

type ItemVendaViewForm struct {
	Produto ProdutoViewForm
	Valor   float32
	Qtd     int
}

type VendaViewForm struct {
	Numero  int
	Data    string
	Cliente ClienteViewForm
	Itens   []ItemVendaViewForm
}

func (c VendaView) Create(clientesform []ClienteViewForm, produtosform []ProdutoViewForm) (VendaViewForm, error) {
	data := time.Now().Format("01-02-2006")
	//var produtoId, qtd int
	var qtd int
	clienteId, _ := c.clienteview.RequestRg(clientesform)

	var itens []ItemVendaViewForm

	for {
		produtoId, _ := c.produtoview.RequestCodigo(produtosform)

		fmt.Printf("\nInforme a quantidade:\n")
		fmt.Scanln(&qtd)

		itens = append(itens, ItemVendaViewForm{Produto: ProdutoViewForm{Codigo: produtoId}, Qtd: qtd})
	}

	resultform := VendaViewForm{Data: data, Cliente: ClienteViewForm{Rg: clienteId}, Itens: itens}
	return resultform, nil
}

func (vv VendaView) RequestNumero(vendas []VendaViewForm) (int, error) {
	//var form VendaViewForm
	var idVenda int
	fmt.Printf("\nIndique o número da venda que deseja remover:\n")
	vv.VisualizeAll(vendas)
	fmt.Printf("\n>>> ")
	fmt.Fscan(stdin, &idVenda)
	return idVenda, nil
}

func (vv VendaView) Update(venda VendaViewForm) (VendaViewForm, error) {

	var opcao int

	for {
		opcao = -1
		fmt.Printf("\nIndique a informação que deseja alterar:\n")
		fmt.Printf("\n[1] Número\n[2] Data\n[3] Cliente\n[4] Itens\n[0] Voltar\n>>> ")
		fmt.Scanln(&opcao)
		if opcao == 0 {
			break
		} else if opcao == 1 {
			fmt.Printf("\nIndique o novo número da venda: ")
			var novodado int
			fmt.Scanln(&novodado)
			venda.Numero = novodado
		} else if opcao == 2 {
			fmt.Printf("\nIndique a nova data da venda: ")
			var novodado string
			fmt.Scanln(&novodado)
			venda.Data = novodado
		} else if opcao == 3 {
			// troca de cliente
		} else if opcao == 4 {
			// remove, adiciona, modifica itens
		}
	}
	return venda, nil
}

func (vv VendaView) Visualize(form VendaViewForm) error {
	fmt.Printf("%d\t%s\t%s\t%v\n", form.Numero, form.Data, form.Cliente.Nome, form.Itens)
	return nil
}

func (vv VendaView) VisualizeAll(form []VendaViewForm) error {
	for i, v := range form {
		fmt.Printf("%d -", i+1)
		vv.Visualize(v)
	}
	return nil
}
