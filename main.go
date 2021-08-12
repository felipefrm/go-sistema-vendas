package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

type ListaCliente []Cliente
type ListaProduto []Produto
type ListaVenda []Venda

type Database struct {
	cliente ListaCliente
	produto ListaProduto
	venda   ListaVenda
}

func main() {

	var opcao int
	db := Database{}

	for {
		fmt.Printf("\n[1] Gerenciar Clientes\n[2] Gerenciar Produtos\n[3] Gerenciar Vendas\n[0] Sair\n>>> ")
		_, err := fmt.Fscan(stdin, &opcao)
		if err != nil {
			fmt.Print(err)
		} else {
			if opcao == 0 {
				break
			} else if opcao == 1 {
				MenuCliente(&(db.cliente))
			} else if opcao == 2 {
				MenuProduto(&(db.produto))
			} else if opcao == 3 {
				MenuVenda(&db.venda, &db.cliente, &db.produto)
			} else {
				fmt.Println("Você deve selecionar uma das opções apresentadas.")
			}
		}
		stdin.ReadString('\n')
	}
}
