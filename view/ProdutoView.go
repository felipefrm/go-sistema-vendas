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

type ProdutoOption int

const (
	OpçãoSairProduto ProdutoOption = iota
	OpçãoVerProdutos
	OpçãoAdicionarProduto
	OpçãoAlterarProduto
	OpçãoRemoverProduto
)

func (v ProdutoView) OptionsMenu() ProdutoOption {

	var opcao ProdutoOption

	for {
		fmt.Printf("\n[1] Visualizar Produtos\n[2] Adicionar Produto\n[3] Alterar Produto\n[4] Remover Produto\n[0] Voltar\n>>> ")
		_, err := fmt.Fscan(stdin, &opcao)
		stdin.ReadString('\n')
		if err == nil {
			break
		}
		fmt.Print(err)
	}

	return opcao
}

func (c ProdutoView) Create() (ProdutoViewForm, error) {
	var form ProdutoViewForm

	for {
		fmt.Printf("\nCódigo: ")
		_, err := fmt.Fscan(stdin, &form.Codigo)
		stdin.ReadString('\n')
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}
	for {
		fmt.Printf("Nome: ")
		_, err := fmt.Fscan(stdin, &form.Nome)
		stdin.ReadString('\n')
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}
	for {
		fmt.Printf("Valor unitário: ")
		_, err := fmt.Fscan(stdin, &form.Valor)
		stdin.ReadString('\n')
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}

	return form, nil
}

func (c ProdutoView) RequestCodigo(produtos []ProdutoViewForm) (int, error) {
	var idProduto int
	for {
		fmt.Printf("\nIndique o ID do produto: [Digite -1 para finalizar inserção de produtos]\n")
		c.VisualizeAll(produtos)
		fmt.Printf("\n>>> ")
		_, err := fmt.Fscan(stdin, &idProduto)
		stdin.ReadString('\n')
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}

	return idProduto, nil
}

func (c ProdutoView) Update(produto ProdutoViewForm) (ProdutoViewForm, error) {

	var opcao int

	for {
		opcao = -1
		fmt.Printf("\nIndique a informação que deseja alterar:\n")
		fmt.Printf("\n[1] Código\n[2] Nome\n[3] Valor unitário\n[0] Voltar\n>>> ")
		_, err := fmt.Fscan(stdin, &opcao)
		stdin.ReadString('\n')
		if err != nil {
			fmt.Print(err)
		} else {
			if opcao == 0 {
				break
			}
			fmt.Printf("\nIndique a nova informação a ser inserida: ")
			switch opcao {
			case 1:
				var novodado int
				for {
					_, err := fmt.Fscan(stdin, &novodado)
					stdin.ReadString('\n')
					if err != nil {
						fmt.Print(err)
					} else {
						break
					}
				}
				produto.Codigo = novodado
			case 2:
				var novodado string
				for {
					_, err := fmt.Fscan(stdin, &novodado)
					stdin.ReadString('\n')
					if err != nil {
						fmt.Print(err)
					} else if len(novodado) <= 0 {
						continue
					} else {
						break
					}
				}
				produto.Nome = novodado
			case 3:
				var novodado float32
				for {
					_, err := fmt.Fscan(stdin, &novodado)
					stdin.ReadString('\n')
					if err != nil {
						fmt.Print(err)
					} else {
						break
					}
				}
				produto.Valor = novodado
			default:
				fmt.Println("Digite uma opção válida.")
			}
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
