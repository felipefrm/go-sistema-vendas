package view

import (
	"fmt"
)

type ProdutoView struct {
}

type ProdutoViewForm struct {
	Codigo int
	Nome   string
	Valor  float32
}

func (c ProdutoView) Create() (ProdutoViewForm, error) {
	var form ProdutoViewForm
	fmt.Printf("Código: ")
	fmt.Scanf("%d\n", &form.Codigo)
	fmt.Printf("\nNome: ")
	fmt.Scanf("%s\n", &form.Nome)
	fmt.Printf("Valor unitário: ")
	fmt.Scanf("%f\n", &form.Valor)
	return form, nil
}

func (c ProdutoView) RequestRG(produtos []ProdutoViewForm) (int, error) {
	var idProduto int
	fmt.Printf("\nIndique o ID do produto que deseja alterar os dados:\n")
	c.VisualizeAll(produtos)
	fmt.Printf("\n>>> ")
	fmt.Scanln(&idProduto)

	return idProduto, nil
}

func (c ProdutoView) Update(produto ProdutoViewForm) (ProdutoViewForm, error) {
	//var novodado string
	var opcao int

	for {
		opcao = -1
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
			produto.Codigo = novodado
		case 2:
			var novodado string
			fmt.Scanln(&novodado)
			produto.Nome = novodado
		case 3:
			var novodado float32
			fmt.Scanln(&novodado)
			produto.Valor = novodado
		}
	}
	return produto, nil
}

func (c ProdutoView) Visualize(v ProdutoViewForm) error {
	fmt.Printf("%d\t%s\t%f\n", v.Codigo, v.Nome, v.Valor)
	return nil
}

func (c ProdutoView) VisualizeAll(form []ProdutoViewForm) error {
	for i, v := range form {
		fmt.Printf("%d -", i+1)
		c.Visualize(v)
	}
	return nil
}
