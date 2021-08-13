package controller

import (
	"fmt"

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
	p := ProdutoViewFormToProduto(itemvenda.Produto)
	return model.ItemVenda{Produto: &p, Valor: itemvenda.Valor, Qtd: itemvenda.Qtd}
}

func ItemVendaToItemVendaViewForm(itemvenda model.ItemVenda) view.ItemVendaViewForm {
	return view.ItemVendaViewForm{Produto: ProdutoToProdutoViewForm(*itemvenda.Produto), Valor: itemvenda.Valor, Qtd: itemvenda.Qtd, Total: itemvenda.Total()}
}

func VendaViewFormToVenda(venda view.VendaViewForm) model.Venda {
	var itens []model.ItemVenda
	for _, x := range venda.Itens {
		itens = append(itens, ItemVendaViewFormToItemVenda(x))
	}
	cliente := ClienteViewFormToCliente(venda.Cliente)
	return model.Venda{Numero: venda.Numero, Data: venda.Data, Cliente: &cliente, Itens: itens}
}

func VendaToVendaViewForm(venda model.Venda) view.VendaViewForm {
	var itensforms []view.ItemVendaViewForm
	for _, x := range venda.Itens {
		itensforms = append(itensforms, ItemVendaToItemVendaViewForm(x))
	}
	return view.VendaViewForm{Numero: venda.Numero, Data: venda.Data, Cliente: ClienteToClienteViewForm(*venda.Cliente), Itens: itensforms}
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
	c, err := contrlr.clientemodel.GetIndex(venda.Cliente)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	ogclient, err := contrlr.clientemodel.GetById(c)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	venda.Cliente = &ogclient

	for i, x := range venda.Itens {
		prod, err := contrlr.produtomodel.GetIndex(x.Produto)
		if err != nil {
			fmt.Printf("%v", err.Error())
			return err
		}
		ogprod, err := contrlr.produtomodel.GetById(prod)
		if err != nil {
			fmt.Printf("%v", err.Error())
			return err
		}
		venda.Itens[i].Produto = &ogprod
	}

	if err := contrlr.vendamodel.Create(&venda); err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
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
	venda, err := contrlr.vendamodel.GetById(numero)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	form := VendaToVendaViewForm(venda)
	outform, _ := contrlr.view.Update(form, clientesforms, produtosforms)
	outvenda := VendaViewFormToVenda(outform)
	c, _ := contrlr.clientemodel.GetIndex(outvenda.Cliente)
	ogclient, _ := contrlr.clientemodel.GetById(c)
	outvenda.Cliente = &ogclient

	for i, x := range outvenda.Itens {
		prod, _ := contrlr.produtomodel.GetIndex(x.Produto)
		ogprod, _ := contrlr.produtomodel.GetById(prod)
		outvenda.Itens[i].Produto = &ogprod
	}
	if err := contrlr.vendamodel.Update(numero, &outvenda); err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	return nil
}

func (contrlr VendaDaoController) Delete() error {
	numero, _ := contrlr.RequestNumero()
	venda, err := contrlr.vendamodel.GetById(numero)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	i, err := contrlr.vendamodel.GetIndex(&venda)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	if err := contrlr.vendamodel.Delete(i); err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
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

func (contrlr VendaDaoController) OptionsMenu() error {
	for {
		option := contrlr.view.OptionsMenu()
		switch option {
		case view.OpçãoSairVenda:
			return nil
		case view.OpçãoVerVendas:
			contrlr.ListAll()
		case view.OpçãoAdicionarVenda:
			contrlr.Create()
		case view.OpçãoAlterarVenda:
			contrlr.Update()
		case view.OpçãoRemoverVenda:
			contrlr.Delete()
		}
	}
}
