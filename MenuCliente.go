package main

import (
	"fmt"
)

func (c ListaCliente) VisualizarClientes() {
	for i, v := range c {
		fmt.Printf("(%d)\t%s\t%s\t%s\t%s\n", i, v.Pessoa.nome, v.Pessoa.sobrenome, v.rg, v.nascimento)
	}
}

func (c *ListaCliente) AdicionarCliente() {
	var nome, sobrenome, rg, nascimento string
	for {
		fmt.Printf("\nNome: ")
		_, err := fmt.Fscan(stdin, &nome)
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}
	for {
		fmt.Printf("Sobrenome: ")
		_, err := fmt.Fscan(stdin, &sobrenome)
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}
	for {
		fmt.Printf("RG: ")
		_, err := fmt.Fscan(stdin, &rg)
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}
	for {
		fmt.Printf("Data de nascimento: ")
		_, err := fmt.Fscan(stdin, &nascimento)
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}

	novoCliente := Cliente{Pessoa{nome, sobrenome}, rg, nascimento}

	*c = append(*c, novoCliente)
}

func (c *ListaCliente) ModificarCliente() {

	var idCliente int
	var novodado string

	for {
		fmt.Printf("\nIndique o ID do cliente que deseja alterar os dados: [Digite -1 para voltar]\n")
		c.VisualizarClientes()
		fmt.Printf("\n>>> ")
		_, err := fmt.Fscan(stdin, &idCliente)
		if err != nil {
			fmt.Print(err)
		} else if idCliente < 0 {
			return
		} else if len(*c) == 0 || idCliente >= len(*c) {
			fmt.Println("Digite um ID válido.")
		} else {
			break
		}
		stdin.ReadString('\n')
	}

	cliente := &(*c)[idCliente]

	var opcao int
	for {
		fmt.Printf("\nIndique a informação que deseja alterar:\n")
		fmt.Printf("\n[1] Nome\n[2] Sobrenome\n[3] RG\n[4] Data de nascimento\n[0] Voltar\n>>> ")
		_, err := fmt.Fscan(stdin, &opcao)
		if err != nil {
			fmt.Print(err)
		} else {
			if opcao == 0 {
				return
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
			default:
				fmt.Println("Digite uma opção válida.")
			}
		}
	}
}

func (c *ListaCliente) RemoverCliente() {
	var idCliente int
	for {
		fmt.Printf("\nIndique o ID do cliente que deseja remover: [Digite -1 para voltar]\n")
		c.VisualizarClientes()
		fmt.Printf("\n>>> ")
		_, err := fmt.Fscan(stdin, &idCliente)
		if err != nil {
			fmt.Print(err)
		} else if idCliente < 0 {
			return
		} else if len(*c) == 0 || idCliente > len(*c) {
			fmt.Println("Digite um ID válido.")
		} else {
			break
		}
		stdin.ReadString('\n')
	}

	aux := make([]Cliente, len(*c))
	copy(aux[:], *c)

	*c = append(aux[:idCliente], aux[idCliente+1:]...)
}

func MenuCliente(cliente *ListaCliente) {

	var opcao int

	for {
		fmt.Printf("\n[1] Visualizar Clientes\n[2] Adicionar Cliente\n[3] Alterar Cliente\n[4] Remover Cliente\n[0] Voltar\n>>> ")
		_, err := fmt.Fscan(stdin, &opcao)
		if err != nil {
			fmt.Print(err)
		} else {
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
				fmt.Println("Você deve selecionar uma das opções apresentadas.")
			}
		}
		stdin.ReadString('\n')
	}
}
