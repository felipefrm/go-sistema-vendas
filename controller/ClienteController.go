package controller

import (
	"fmt"

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
	if err := contrlr.model.Create(&cliente); err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	return nil
}

func (contrlr ClienteDaoController) RequestRG() (string, error) {
	clientes, _ := contrlr.model.GetAll()
	var forms []view.ClienteViewForm
	for _, x := range clientes {
		forms = append(forms, ClienteToClienteViewForm(*x))
	}
	rg, _ := contrlr.view.RequestRg(forms)
	return rg, nil
}
func (contrlr ClienteDaoController) Update() error {
	rg, _ := contrlr.RequestRG()
	cliente, err := contrlr.model.GetById(rg)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	} else {
		form := ClienteToClienteViewForm(*cliente)
		outform, _ := contrlr.view.Update(form)
		outcliente := ClienteViewFormToCliente(outform)
		if err := contrlr.model.Update(rg, &outcliente); err != nil {
			fmt.Printf("%v", err.Error())
			return err
		}
		return nil
	}
}

func (contrlr ClienteDaoController) Delete() error {
	rg, _ := contrlr.RequestRG()
	cliente, err := contrlr.model.GetById(rg)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	i, err := contrlr.model.GetIndex(cliente)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	if err := contrlr.model.Delete(i); err != nil {
		fmt.Printf("%v", err.Error())
	}
	return nil
}

func (contrlr ClienteDaoController) ListAll() error {
	clientes, _ := contrlr.model.GetAll()
	var forms []view.ClienteViewForm
	for _, x := range clientes {
		forms = append(forms, ClienteToClienteViewForm(*x))
	}
	contrlr.view.VisualizeAll(forms)
	return nil
}

func (contrlr ClienteDaoController) OptionsMenu() error {
	for {
		option := contrlr.view.OptionsMenu()
		switch option {
		case view.OpçãoSairCliente:
			return nil
		case view.OpçãoVerClientes:
			contrlr.ListAll()
		case view.OpçãoAdicionarCliente:
			contrlr.Create()
		case view.OpçãoAlterarCliente:
			contrlr.Update()
		case view.OpçãoRemoverCliente:
			contrlr.Delete()
		}
	}
}
