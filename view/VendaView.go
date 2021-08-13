package view

import (
	"fmt"
	"time"
)

type VendaView struct {
	clienteview ClienteView
	produtoview ProdutoView
}

type ItemVendaViewForm struct {
	Produto ProdutoViewForm
	Valor   float32
	Qtd     int
}

type VendaViewForm struct {
	Numero  int
	Data    string
	Cliente ClienteViewForm
	Itens   []ItemVendaViewForm
}

type VendaOption int

const (
	OpçãoSairVenda VendaOption = iota
	OpçãoVerVendas
	OpçãoAdicionarVenda
	OpçãoAlterarVenda
	OpçãoRemoverVenda
)

func (v VendaView) OptionsMenu() VendaOption {

	var opcao VendaOption

	for {
		fmt.Printf("\n[1] Visualizar Vendas\n[2] Adicionar Venda\n[3] Alterar Venda\n[4] Remover Venda\n[0] Voltar\n>>> ")
		_, err := fmt.Fscan(stdin, &opcao)
		stdin.ReadString('\n')
		if err == nil {
			break
		}
		fmt.Print(err)
	}

	return opcao
}

func (c VendaView) Create(clientesform []ClienteViewForm, produtosform []ProdutoViewForm) (VendaViewForm, error) {
	data := time.Now().Format("01-02-2006")
	//var produtoId, qtd int
	var qtd, numero int

	for {
		fmt.Printf("\nInforme o número da venda: ")
		_, err := fmt.Fscan(stdin, &numero)
		stdin.ReadString('\n')
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}

	clienteId, _ := c.clienteview.RequestRg(clientesform)

	var itens []ItemVendaViewForm

	for {
		produtoId, _ := c.produtoview.RequestCodigo(produtosform)

		if produtoId == -1 {
			break
		}

		for {
			fmt.Printf("\nInforme a quantidade: ")
			_, err := fmt.Fscan(stdin, &qtd)
			stdin.ReadString('\n')
			if err != nil {
				fmt.Print(err)
			} else {
				break
			}
		}

		itens = append(itens, ItemVendaViewForm{Produto: ProdutoViewForm{Codigo: produtoId}, Qtd: qtd})
	}

	resultform := VendaViewForm{Numero: numero, Data: data, Cliente: ClienteViewForm{Rg: clienteId}, Itens: itens}
	return resultform, nil
}

func (vv VendaView) RequestNumero(vendas []VendaViewForm) (int, error) {
	//var form VendaViewForm
	var idVenda int
	for {
		fmt.Printf("\nIndique o número da venda:\n")
		vv.VisualizeAll(vendas)
		fmt.Printf("\n>>> ")
		_, err := fmt.Fscan(stdin, &idVenda)
		stdin.ReadString('\n')
		if err != nil {
			fmt.Print(err)
		} else {
			break
		}
	}
	return idVenda, nil
}

func (vv VendaView) Update(venda VendaViewForm, clientes []ClienteViewForm, produtos []ProdutoViewForm) (VendaViewForm, error) {

	var opcao int

	for {
		for {
			opcao = -1
			fmt.Printf("\nIndique a informação que deseja alterar:\n")
			fmt.Printf("\n[1] Número\n[2] Data\n[3] Cliente\n[4] Itens\n[0] Voltar\n>>> ")
			_, err := fmt.Fscan(stdin, &opcao)
			stdin.ReadString('\n')
			if err != nil {
				fmt.Print(err)
			} else {
				break
			}
		}
		if opcao == 0 {
			break
		} else if opcao == 1 {
			var novodado int
			for {
				fmt.Printf("\nIndique o novo número da venda: ")
				_, err := fmt.Fscan(stdin, &novodado)
				stdin.ReadString('\n')
				if err != nil {
					fmt.Print(err)
				} else {
					break
				}
			}
			venda.Numero = novodado
		} else if opcao == 2 {
			var novodado string
			for {
				fmt.Printf("\nIndique a nova data da venda: ")
				_, err := fmt.Fscan(stdin, &novodado)
				stdin.ReadString('\n')
				if err != nil {
					fmt.Print(err)
				} else {
					break
				}
			}
			venda.Data = novodado
		} else if opcao == 3 {
			clienteId, _ := vv.clienteview.RequestRg(clientes)
			venda.Cliente.Rg = clienteId
		} else if opcao == 4 {
			for {
				for {
					fmt.Printf("\n[1] Adicionar item\n[2] Remover item\n[3] Alterar item\n[0] Voltar\n>>> ")
					_, err := fmt.Fscan(stdin, &opcao)
					stdin.ReadString('\n')
					if err != nil {
						fmt.Print(err)
					} else {
						break
					}
				}
				if opcao == 0 {
					break
				} else if opcao == 1 {
					for {
						produtoId, _ := vv.produtoview.RequestCodigo(produtos)
						if produtoId == -1 {
							break
						}
						var qtd int
						for {
							fmt.Printf("\nInforme a quantidade:\n")
							_, err := fmt.Fscan(stdin, &qtd)
							stdin.ReadString('\n')
							if err != nil {
								fmt.Print(err)
							} else {
								break
							}
						}
						venda.Itens = append(venda.Itens, ItemVendaViewForm{Produto: ProdutoViewForm{Codigo: produtoId}, Qtd: qtd})
					}
				} else if opcao == 2 {
					// remove
				} else if opcao == 3 {
					// modifica
				} else {
					fmt.Println("Digite uma opção válida.")
				}
			}
		} else {
			fmt.Println("Digite uma opção válida.")
		}
	}
	return venda, nil
}

func (vv VendaView) Visualize(form VendaViewForm) error {
	fmt.Printf("(%d)\t%s\t%s\t%v\n", form.Numero, form.Data, form.Cliente.Nome, form.Itens)
	return nil
}

func (vv VendaView) VisualizeAll(form []VendaViewForm) error {
	for _, v := range form {
		vv.Visualize(v)
	}
	return nil
}
