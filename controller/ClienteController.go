package controller

import (
	dao "github.com/felipefrm/go-sistema-vendas/dao"
	model "github.com/felipefrm/go-sistema-vendas/model"
	view "github.com/felipefrm/go-sistema-vendas/view"
)

type ClienteDaoController struct {
	model dao.ClienteDao
	view  view.ClienteView
}

func ClienteViewFormToCliente(f view.ClienteViewForm) model.Cliente {
	return model.Cliente{Pessoa: model.Pessoa{Nome: f.Nome, Sobrenome: f.Sobrenome}, Rg: f.Rg, Nascimento: f.Nascimento}
}

func ClienteToClienteViewForm(cliente model.Cliente) view.ClienteViewForm {
	return view.ClienteViewForm{
		Nome:       cliente.Nome,
		Sobrenome:  cliente.Sobrenome,
		Rg:         cliente.Rg,
		Nascimento: cliente.Nascimento,
	}
}

func (contrlr ClienteDaoController) Create() error {
	var f view.ClienteViewForm
	f, _ = contrlr.view.Create()
	cliente := ClienteViewFormToCliente(f)
	contrlr.model.Create(&cliente)
	return nil
}

func (contrlr ClienteDaoController) RequestRG() (string, error) {
	clientes, _ := contrlr.model.GetAll()
	var forms []view.ClienteViewForm
	for _, x := range clientes {
		forms = append(forms, ClienteToClienteViewForm(x))
	}
	rg, _ := contrlr.view.RequestRG(forms)
	return rg, nil
}
func (contrlr ClienteDaoController) Update() error {
	rg, _ := contrlr.RequestRG()
	cliente, _ := contrlr.model.GetById(rg)
	form := ClienteToClienteViewForm(cliente)
	outform, _ := contrlr.view.Update(form)
	outcliente := ClienteViewFormToCliente(outform)
	contrlr.model.Update(rg, &outcliente)
	return nil
}

func (contrlr ClienteDaoController) Delete() error {
	rg, _ := contrlr.RequestRG()
	cliente, _ := contrlr.model.GetById(rg)
	i, _ := contrlr.model.GetIndex(&cliente)
	contrlr.model.Delete(i)
	return nil
}

func (contrlr ClienteDaoController) ListAll() error {
	clientes, _ := contrlr.model.GetAll()
	var forms []view.ClienteViewForm
	for _, x := range clientes {
		forms = append(forms, ClienteToClienteViewForm(x))
	}
	contrlr.view.VisualizeAll(forms)
	return nil
}
