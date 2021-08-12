package controller

import (
	dao "github.com/felipefrm/go-sistema-vendas/dao"
	model "github.com/felipefrm/go-sistema-vendas/model"
	view "github.com/felipefrm/go-sistema-vendas/view"
)

type VendaDaoController struct {
	vendamodel   dao.VendaDao
	clientemodel dao.ClienteDao
	produtomodel dao.ProdutoDao
	view         view.VendaView
}

func ItemVendaViewFormToItemVenda(itemvenda view.ItemVendaViewForm) model.ItemVenda {
	return model.ItemVenda{Produto: ProdutoViewFormToProduto(itemvenda.Produto), Valor: itemvenda.Valor, Qtd: itemvenda.Qtd}
}

func ItemVendaToItemVendaViewForm(itemvenda model.ItemVenda) view.ItemVendaViewForm {
	return view.ItemVendaViewForm{Produto: ProdutoToProdutoViewForm(itemvenda.Produto), Valor: itemvenda.Valor, Qtd: itemvenda.Qtd}
}

func VendaViewFormToVenda(venda view.VendaViewForm) model.Venda {
	var itens []model.ItemVenda
	for _, x := range venda.Itens {
		itens = append(itens, ItemVendaViewFormToItemVenda(x))
	}
	return model.Venda{Numero: venda.Numero, Data: venda.Data, Cliente: ClienteViewFormToCliente(venda.Cliente), Itens: itens}
}

func VendaToVendaViewForm(venda model.Venda) view.VendaViewForm {
	var itensforms []view.ItemVendaViewForm
	for _, x := range venda.Itens {
		itensforms = append(itensforms, ItemVendaToItemVendaViewForm(x))
	}
	return view.VendaViewForm{Numero: venda.Numero, Data: venda.Data, Cliente: ClienteToClienteViewForm(venda.Cliente), Itens: itensforms}
}

func (contrlr VendaDaoController) Create() error {
	var f view.VendaViewForm
	clientes, _ := contrlr.clientemodel.GetAll()
	var clientesforms []view.ClienteViewForm
	for _, x := range clientes {
		clientesforms = append(clientesforms, ClienteToClienteViewForm(x))
	}

	produtos, _ := contrlr.produtomodel.GetAll()
	var produtosforms []view.ProdutoViewForm
	for _, x := range produtos {
		produtosforms = append(produtosforms, ProdutoToProdutoViewForm(x))
	}
	f, _ = contrlr.view.Create(clientesforms, produtosforms)
	venda := VendaViewFormToVenda(f)
	contrlr.vendamodel.Create(&venda)
	return nil
}

func (contrlr VendaDaoController) RequestNumero() (int, error) {
	vendas, _ := contrlr.vendamodel.GetAll()
	var forms []view.VendaViewForm
	for _, x := range vendas {
		forms = append(forms, VendaToVendaViewForm(x))
	}
	numero, _ := contrlr.view.RequestNumero(forms)
	return numero, nil
}
func (contrlr VendaDaoController) Update() error {
	clientes, _ := contrlr.clientemodel.GetAll()
	var clientesforms []view.ClienteViewForm
	for _, x := range clientes {
		clientesforms = append(clientesforms, ClienteToClienteViewForm(x))
	}
	produtos, _ := contrlr.produtomodel.GetAll()
	var produtosforms []view.ProdutoViewForm
	for _, x := range produtos {
		produtosforms = append(produtosforms, ProdutoToProdutoViewForm(x))
	}

	numero, _ := contrlr.RequestNumero()
	venda, _ := contrlr.vendamodel.GetById(numero)
	form := VendaToVendaViewForm(venda)
	outform, _ := contrlr.view.Update(form, clientesforms, produtosforms)
	outvenda := VendaViewFormToVenda(outform)
	contrlr.vendamodel.Update(numero, &outvenda)
	return nil
}

func (contrlr VendaDaoController) Delete() error {
	numero, _ := contrlr.RequestNumero()
	venda, _ := contrlr.vendamodel.GetById(numero)
	i, _ := contrlr.vendamodel.GetIndex(&venda)
	contrlr.vendamodel.Delete(i)
	return nil
}

func (contrlr VendaDaoController) ListAll() error {
	vendas, _ := contrlr.vendamodel.GetAll()
	var forms []view.VendaViewForm
	for _, x := range vendas {
		forms = append(forms, VendaToVendaViewForm(x))
	}
	contrlr.view.VisualizeAll(forms)
	return nil
}
