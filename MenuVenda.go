package main

import (
	"fmt"
	"time"
)

var contVenda = 0

func (venda ListaVenda) VisualizarVendas() {
	for i, v := range venda {
		fmt.Printf("(%d)\t%d\t%s\t%s\t%v\n", i, v.numero, v.data, v.cliente.Pessoa.nome, v.itens)
	}
}

func (v *ListaVenda) AdicionarVenda(Cliente ListaCliente, Produto ListaProduto) {
	data := time.Now().Format("01-02-2006")
	var clienteId, produtoId, qtd int
	fmt.Printf("\nIndique o cliente que está fazendo a compra:\n")
	Cliente.VisualizarClientes()
	fmt.Printf("\n>>> ")
	fmt.Scanln(&clienteId)

	var itens []ItemVenda

	for {
		fmt.Printf("\nSelecione um produto para adicioná-lo à venda [Digite -1 para finalizar a venda]:\n")
		Produto.VisualizarProdutos()
		fmt.Printf("\n>>> ")
		fmt.Scanln(&produtoId)

		if produtoId == -1 {
			break
		}

		fmt.Printf("\nInforme a quantidade:\n")
		fmt.Scanln(&qtd)

		itens = append(itens, ItemVenda{Produto[produtoId], Produto[produtoId].valor, qtd})
	}

	novaVenda := Venda{contVenda, data, Cliente[clienteId], itens}
	contVenda += 1

	*v = append(*v, novaVenda)
}

func (v *ListaVenda) ModificarVenda(Cliente ListaCliente, Produto ListaProduto) {

	var idVenda int

	fmt.Printf("\nIndique o ID da Venda que deseja alterar os dados:\n")
	v.VisualizarVendas()
	fmt.Printf("\n>>> ")
	fmt.Scanln(&idVenda)
	venda := &(*v)[idVenda]
	for {
		opcao := -1
		fmt.Printf("\nIndique a informação que deseja alterar:\n")
		fmt.Printf("\n[1] Número\n[2] Data\n[3] Cliente\n[4] Itens\n[0] Voltar\n>>> ")
		fmt.Scanln(&opcao)
		if opcao == 0 {
			break
		} else if opcao == 1 {
			fmt.Printf("\nIndique o novo número da venda: ")
			var novodado int
			fmt.Scanln(&novodado)
			venda.numero = novodado
		} else if opcao == 2 {
			fmt.Printf("\nIndique a nova data da venda: ")
			var novodado string
			fmt.Scanln(&novodado)
			venda.data = novodado
		} else if opcao == 3 {
			var clienteId int
			fmt.Printf("\nIndique o novo cliente da venda:\n")
			Cliente.VisualizarClientes()
			fmt.Printf("\n>>> ")
			fmt.Scanln(&clienteId)
			venda.cliente = Cliente[clienteId]
		} else if opcao == 4 {
			// remove, adiciona, modifica itens
		}
	}
}

func (v *ListaVenda) RemoverVenda() {
	var idVenda int
	fmt.Printf("\nIndique o ID do cliente que deseja remover:\n")
	v.VisualizarVendas()
	fmt.Printf("\n>>> ")
	fmt.Scanln(&idVenda)

	aux := make(ListaVenda, len(*v))
	copy(aux[:], *v)

	*v = append(aux[:idVenda], aux[idVenda+1:]...)
}

func MenuVenda(Venda *ListaVenda, Cliente *ListaCliente, Produto *ListaProduto) {

	var opcao int

	for {
		fmt.Printf("\n[1] Visualizar Vendas\n[2] Adicionar Venda\n[3] Alterar Venda\n[4] Remover Venda\n[0] Voltar\n>>> ")
		fmt.Scanf("%d\n", &opcao)
		if opcao == 0 {
			break
		} else if opcao == 1 {
			Venda.VisualizarVendas()
		} else if opcao == 2 {
			Venda.AdicionarVenda(*Cliente, *Produto)
		} else if opcao == 3 {
			Venda.ModificarVenda(*Cliente, *Produto)
		} else if opcao == 4 {
			Venda.RemoverVenda()
		} else {
			// errado
		}
	}
}
