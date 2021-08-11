package main

import "fmt"

func (p ListaProduto) VisualizarProdutos() {
	for i, v := range p {
		fmt.Printf("(%d)\t%d\t%s\t%f\n", i, v.codigo, v.nome, v.valor)
	}
}

func (p *ListaProduto) AdicionarProduto() {
	var codigo int
	var nome string
	var valor float32
	fmt.Printf("Código: ")
	fmt.Scanf("%d\n", &codigo)
	fmt.Printf("\nNome: ")
	fmt.Scanf("%s\n", &nome)
	fmt.Printf("Valor unitário: ")
	fmt.Scanf("%f\n", &valor)

	novoProduto := Produto{codigo, nome, valor}

	*p = append(*p, novoProduto)
}

func (p *ListaProduto) ModificarProduto() {

	var idProduto int

	fmt.Printf("\nIndique o ID do produto que deseja alterar os dados:\n")
	p.VisualizarProdutos()
	fmt.Printf("\n>>> ")
	fmt.Scanln(&idProduto)
	produto := &(*p)[idProduto]
	for {
		opcao := -1
		fmt.Printf("\nIndique a informação que deseja alterar:\n")
		fmt.Printf("\n[1] Código\n[2] Nome\n[3] Valor unitário\n[0] Voltar\n>>> ")
		fmt.Scanln(&opcao)
		if opcao == 0 {
			break
		}
		fmt.Printf("\nIndique a nova informação a ser inserida: ")

		switch opcao {
		case 1:
			var novodado int
			fmt.Scanln(&novodado)
			produto.codigo = novodado
		case 2:
			var novodado string
			fmt.Scanln(&novodado)
			produto.nome = novodado
		case 3:
			var novodado float32
			fmt.Scanln(&novodado)
			produto.valor = novodado

		}
	}
}

func (p *ListaProduto) RemoverProduto() {
	var idProduto int
	fmt.Printf("\nIndique o ID do produto que deseja remover:\n")
	p.VisualizarProdutos()
	fmt.Printf("\n>>> ")
	fmt.Scanln(&idProduto)

	aux := make(ListaProduto, len(*p))
	copy(aux[:], *p)

	*p = append(aux[:idProduto], aux[idProduto+1:]...)
}

func MenuProduto(produto *ListaProduto) {

	var opcao int

	for {
		fmt.Printf("\n[1] Visualizar Produtos\n[2] Adicionar Produto\n[3] Alterar Produto\n[4] Remover Produto\n[0] Voltar\n>>> ")
		fmt.Scanf("%d\n", &opcao)
		if opcao == 0 {
			break
		} else if opcao == 1 {
			produto.VisualizarProdutos()
		} else if opcao == 2 {
			produto.AdicionarProduto()
		} else if opcao == 3 {
			produto.ModificarProduto()
		} else if opcao == 4 {
			produto.RemoverProduto()
		} else {
			// errado
		}
	}

}
