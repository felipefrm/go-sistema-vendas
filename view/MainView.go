package view
//import ("fmt")

type MainOption int
const (
	OpçãoSair MainOption = iota
	OpçãoCliente
	OpçãoProduto
	OpçãoVenda
)


type MainView struct{

}

func (c MainView) OptionsMenu() MainOption{
	// TODO
	return OpçãoSair
}
