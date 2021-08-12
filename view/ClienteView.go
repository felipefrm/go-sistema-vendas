package view
import ("fmt" "os" "bufio")
var stdin = bufio.NewReader(os.Stdin)
//type ClienteDao interface{
	//Create(u *model.Cliente) error
	//Update(i ClienteIndexType, u *model.Cliente) error
	//Delete(i ClienteIndexType) error
	//GetIndex(u *model.Cliente) (ClienteIndexType, error)
	//GetById(i ClienteIndexType) (model.Cliente, error)
	//GetAll() ([]model.Cliente, error)
//}

type ClienteView struct{
	name string
}

type ClienteViewForm struct{
	Nome string
	Sobrenome string
	Rg string
	Nascimento string
}
func (c ClienteView) Create() ClienteViewForm{
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
	return form
}

func (c ClienteView) RequestRG(clientes []model.Cliente) string{
	var form ClienteViewForm
	var idCliente string
	for {
		fmt.Printf("\nIndique o RG do cliente que deseja alterar os dados: [Digite -1 para voltar]\n")
		//c.VisualizarClientes()
		fmt.Printf("\n>>> ")
		_, err := fmt.Fscan(stdin, &idCliente)
		if err != nil {
			fmt.Print(err)
		} else if idCliente < 0 {
			return
		} else if len(*c) == 0 || idCliente >= len(*c) {
			fmt.Println("Digite um RG válido.")
		} else {
			break
		}
		stdin.ReadString('\n')
	}
	//form.Rg = idCliente
	return idCliente
}

func (c ClienteView) Update(form ClienteViewForm) ClienteViewForm{
	//var form ClienteViewForm
	var opcao int
	var cliente ClienteViewForm
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
	return form
}

func (c ClienteView) Visualize(form ClienteViewForm) error{
	fmt.Printf("(%d)\t%s\t%s\t%s\t%s\n", v.Nome, v.Sobrenome, v.Rg, v.Nascimento)
}

func (c ClienteView) VisualizeList(form []ClienteViewForm) error{
	for i, v := range form {
		c.Visualize(v)
	}
}
