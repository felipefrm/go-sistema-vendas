package main

import "fmt"

type Database struct {
	cliente []Cliente
	produto []Produto
	venda   []Venda
}

func main() {

	var opcao, opcaocrud int
	// db := Database{}

	var EntityMap = map[int]string{
		1: "Cliente",
		2: "Produto",
		3: "Venda",
	}

	for {
		fmt.Printf("\n[1] CRUD Cliente\n[2] CRUD Produto\n[3] CRUD Venda\n[0] Sair\n>>> ")
		fmt.Scanf("%d\n", &opcao)

		if opcao == 0 {
			break
		} else if opcao > 0 && opcao <= 3 {

			for {

				fmt.Printf("\n[1] Visualizar %ss\n[2] Adicionar %s\n[3] Alterar %s\n[4] Remover %s\n[0] Voltar\n>>> ",
					EntityMap[opcao], EntityMap[opcao], EntityMap[opcao], EntityMap[opcao])
				fmt.Scanf("%d\n", &opcaocrud)
				if opcaocrud == 0 {
					break
				}
			}
		} else {
			// errado
		}

	}
}
