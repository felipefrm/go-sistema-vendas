package main

import "fmt"

func (c ListaCliente) VisualizarClientes() {
	for i, v := range c {
		fmt.Printf("(%d)\t%s\t%s\t%s\t%s\n", i, v.Pessoa.nome, v.Pessoa.sobrenome, v.rg, v.nascimento)
	}
}

func (c *ListaCliente) AdicionarCliente() {
	var nome, sobrenome, rg, nascimento string
	fmt.Printf("\nNome: ")
	fmt.Scanf("%s\n", &nome)
	fmt.Printf("Sobrenome: ")
	fmt.Scanf("%s\n", &sobrenome)
	fmt.Printf("RG: ")
	fmt.Scanf("%s\n", &rg)
	fmt.Printf("Data de nascimento: ")
	fmt.Scanf("%s\n", &nascimento)

	novoCliente := Cliente{Pessoa{nome, sobrenome}, rg, nascimento}

	*c = append(*c, novoCliente)
}

func (c *ListaCliente) ModificarCliente() {

	var idCliente int
	var novodado string

	fmt.Printf("\nIndique o ID do cliente que deseja alterar os dados:\n")
	c.VisualizarClientes()
	fmt.Printf("\n>>> ")
	fmt.Scanln(&idCliente)
	cliente := &(*c)[idCliente]
	for {
		opcao := -1
		fmt.Printf("\nIndique a informação que deseja alterar:\n")
		fmt.Printf("\n[1] Nome\n[2] Sobrenome\n[3] RG\n[4] Data de nascimento\n[0] Voltar\n>>> ")
		fmt.Scanln(&opcao)
		if opcao == 0 {
			break
		}
		fmt.Printf("\nIndique a nova informação a ser inserida: ")
		fmt.Scanln(&novodado)
		switch opcao {
		case 1:
			cliente.Pessoa.nome = novodado
		case 2:
			cliente.Pessoa.sobrenome = novodado
		case 3:
			cliente.rg = novodado
		case 4:
			cliente.nascimento = novodado
		}
	}
}

func (c *ListaCliente) RemoverCliente() {
	var idCliente int
	fmt.Printf("\nIndique o ID do cliente que deseja remover:\n")
	c.VisualizarClientes()
	fmt.Printf("\n>>> ")
	fmt.Scanln(&idCliente)

	aux := make([]Cliente, len(*c))
	copy(aux[:], *c)

	*c = append(aux[:idCliente], aux[idCliente+1:]...)
}

func MenuCliente(cliente *ListaCliente) {

	var opcao int

	for {
		fmt.Printf("\n[1] Visualizar Clientes\n[2] Adicionar Cliente\n[3] Alterar Cliente\n[4] Remover Cliente\n[0] Voltar\n>>> ")
		fmt.Scanf("%d\n", &opcao)
		if opcao == 0 {
			break
		} else if opcao == 1 {
			cliente.VisualizarClientes()
		} else if opcao == 2 {
			cliente.AdicionarCliente()
		} else if opcao == 3 {
			cliente.ModificarCliente()
		} else if opcao == 4 {
			cliente.RemoverCliente()
		} else {
			// errado
		}
	}
}
