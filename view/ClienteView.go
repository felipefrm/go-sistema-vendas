package view

import (
	"bufio"
	"fmt"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

//type ClienteDao interface{
//Create(u *model.Cliente) error
//Update(i ClienteIndexType, u *model.Cliente) error
//Delete(i ClienteIndexType) error
//GetIndex(u *model.Cliente) (ClienteIndexType, error)
//GetById(i ClienteIndexType) (model.Cliente, error)
//GetAll() ([]model.Cliente, error)
//}

type ClienteView struct {
}

type ClienteViewForm struct {
	Nome       string
	Sobrenome  string
	Rg         string
	Nascimento string
}

type ClienteOption int

const (
	OpçãoSairCliente ClienteOption = iota
	OpçãoVerClientes
	OpçãoAdicionarCliente
	OpçãoAlterarCliente
	OpçãoRemoverCliente
)

func (v ClienteView) OptionsMenu() ClienteOption {

	var opcao ClienteOption

	for {
		fmt.Printf("\n[1] Visualizar Clientes\n[2] Adicionar Cliente\n[3] Alterar Cliente\n[4] Remover Cliente\n[0] Voltar\n>>> ")
		_, err := fmt.Fscan(stdin, &opcao)
		if err == nil {
			break
		}
		fmt.Print(err)
		stdin.ReadString('\n')
	}

	return opcao
}

func (c ClienteView) Create() (ClienteViewForm, error) {
	var form ClienteViewForm
	for {
		fmt.Printf("\nNome: ")
		_, err := fmt.Fscan(stdin, &form.Nome)
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}
	for {
		fmt.Printf("Sobrenome: ")
		_, err := fmt.Fscan(stdin, &form.Sobrenome)
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}
	for {
		fmt.Printf("RG: ")
		_, err := fmt.Fscan(stdin, &form.Rg)
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}
	for {
		fmt.Printf("Data de nascimento: ")
		_, err := fmt.Fscan(stdin, &form.Nascimento)
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}
	return form, nil
}

func (c ClienteView) RequestRg(clientes []ClienteViewForm) (string, error) {
	//var form ClienteViewForm
	var idCliente string
	for {
		fmt.Printf("\nIndique o RG do cliente:\n")
		//c.VisualizarClientes()
		c.VisualizeAll(clientes)
		fmt.Printf("\n>>> ")
		_, err := fmt.Fscan(stdin, &idCliente)
		if err != nil {
			fmt.Print(err)
		} else if len(idCliente) <= 0 {
			continue
		} else {
			break
		}
		stdin.ReadString('\n')
	}
	//form.Rg = idCliente
	return idCliente, nil
}

func (c ClienteView) Update(cliente ClienteViewForm) (ClienteViewForm, error) {
	var novodado string
	var opcao int
	for {
		fmt.Printf("\nIndique a informação que deseja alterar:\n")
		fmt.Printf("\n[1] Nome\n[2] Sobrenome\n[3] RG\n[4] Data de nascimento\n[0] Voltar\n>>> ")
		_, err := fmt.Fscan(stdin, &opcao)
		if err != nil {
			fmt.Print(err)
		} else {
			if opcao == 0 {
				break
			}
			for {
				fmt.Printf("\nIndique a nova informação a ser inserida: ")
				fmt.Scanln(&novodado)
				if len(novodado) > 0 {
					break
				}
			}
			switch opcao {
			case 1:
				cliente.Nome = novodado
			case 2:
				cliente.Sobrenome = novodado
			case 3:
				cliente.Rg = novodado
			case 4:
				cliente.Nascimento = novodado
			default:
				fmt.Println("Digite uma opção válida.")
			}
		}
	}
	return cliente, nil
}

func (c ClienteView) Visualize(form ClienteViewForm) error {
	fmt.Printf("%s\t%s\t%s\t%s\n", form.Nome, form.Sobrenome, form.Rg, form.Nascimento)
	return nil
}

func (c ClienteView) VisualizeAll(form []ClienteViewForm) error {
	for i, v := range form {
		fmt.Printf("%d -", i+1)
		c.Visualize(v)
	}
	return nil
}
