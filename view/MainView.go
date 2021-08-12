package view

import "fmt"

type MainOption int

const (
	OpçãoSair MainOption = iota
	OpçãoCliente
	OpçãoProduto
	OpçãoVenda
)

type MainView struct {
}

func (v MainView) OptionsMenu() MainOption {

	var opcao MainOption

	for {
		fmt.Printf("\n[1] Gerenciar Clientes\n[2] Gerenciar Produtos\n[3] Gerenciar Vendas\n[0] Sair\n>>> ")
		_, err := fmt.Fscan(stdin, &opcao)
		stdin.ReadString('\n')
		if err == nil {
			break
		}
		fmt.Print(err)
	}

	return opcao
}
