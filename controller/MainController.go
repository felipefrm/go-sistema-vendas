package controller

import (
	dao "github.com/felipefrm/go-sistema-vendas/dao"
	model "github.com/felipefrm/go-sistema-vendas/model"
	view "github.com/felipefrm/go-sistema-vendas/view"
)

type MainDaoController struct {
	View view.MainView
}

func (contrlr MainDaoController) OptionsMenu() error {
	vendadao := dao.VendaDaoMap{Model: make(map[dao.VendaIndexType]model.Venda), ClientesVendasNumero: make(map[dao.ClienteIndexType]map[dao.VendaIndexType]bool),

		ProdutosVendasNumero: make(map[dao.ProdutoIndexType]map[dao.VendaIndexType]bool)}
	clientedao := dao.ClienteDaoMap{Model: make(map[dao.ClienteIndexType]*model.Cliente), Vendadaomap: &vendadao}
	//clientedao.Init()
	produtodao := dao.ProdutoDaoMap{Model: make(map[dao.ProdutoIndexType]*model.Produto), Vendadaomap: &vendadao}

	c1 := model.Cliente{model.Pessoa{"João", "da Silva"}, "121241", "02-06-1999"}
	c2 := model.Cliente{model.Pessoa{"José", "Lucas"}, "215122", "21-02-1982"}
	c3 := model.Cliente{model.Pessoa{"Maria", "Joaquina"}, "2151225", "15-07-1987"}
	clientedao.Create(&c1)
	clientedao.Create(&c2)
	clientedao.Create(&c3)

	p1 := model.Produto{Codigo: 1, Nome: "Biscoito", Valor: 2.5}
	p2 := model.Produto{Codigo: 2, Nome: "Pó de Café", Valor: 10.5}
	p3 := model.Produto{Codigo: 3, Nome: "Xícara", Valor: 5}
	produtodao.Create(&p1)
	produtodao.Create(&p2)
	produtodao.Create(&p3)

	itens := [2]model.ItemVenda{{&p1, 2.5, 5}, {&p3, 5, 10}}
	v1 := model.Venda{1, "12-08-2021", &c1, itens[:]}
	itens = [2]model.ItemVenda{{&p2, 10.5, 1}, {&p3, 5, 10}}
	v2 := model.Venda{2, "12-08-2021", &c2, itens[:]}
	vendadao.Create(&v1)
	vendadao.Create(&v2)

	clientecontroller := ClienteDaoController{model: clientedao, view: view.ClienteView{}}
	produtocontroller := ProdutoDaoController{model: produtodao, view: view.ProdutoView{}}
	vendacontroller := VendaDaoController{vendamodel: vendadao, clientemodel: clientedao, produtomodel: produtodao, view: view.VendaView{}}

	for {
		option := contrlr.View.OptionsMenu()
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
}
