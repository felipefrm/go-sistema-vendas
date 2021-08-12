package controller

import (
	dao "github.com/felipefrm/go-sistema-vendas/dao"
	view "github.com/felipefrm/go-sistema-vendas/view"
)

type MainDaoController struct {
	view view.MainView
}

func (contrlr MainDaoController) OptionsMenu() error {
	clientedao := dao.ClienteDaoMap{}
	produtodao := dao.ProdutoDaoMap{}
	vendadao := dao.VendaDaoMap{}
	clientecontroller := ClienteDaoController{model: clientedao, view: view.ClienteView{}}
	produtocontroller := ProdutoDaoController{model: produtodao, view: view.ProdutoView{}}
	vendacontroller := VendaDaoController{vendamodel: vendadao, clientemodel: clientedao, produtomodel: produtodao, view: view.VendaView{}}

	for {
		option := contrlr.view.OptionsMenu()
		switch option {
		case view.OpçãoSair:
			return nil
		case view.OpçãoCliente:
			clientecontroller.OptionsMenu()
		case view.OpçãoProduto:
			produtocontroller.OptionsMenu()
		case view.OpçãoVenda:
			vendacontroller.OptionsMenu()
		}
	}
	return nil
}
