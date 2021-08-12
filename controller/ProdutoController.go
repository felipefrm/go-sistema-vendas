package controller

import (
	dao "github.com/felipefrm/go-sistema-vendas/dao"
	model "github.com/felipefrm/go-sistema-vendas/model"
	view "github.com/felipefrm/go-sistema-vendas/view"
)

type ProdutoDaoController struct {
	model dao.ProdutoDao
	view  view.ProdutoView
}

func ProdutoViewFormToProduto(f view.ProdutoViewForm) model.Produto {
	return model.Produto{Pessoa: model.Pessoa{Nome: f.Nome, Sobrenome: f.Sobrenome}, Codigo: f.Codigo, Nascimento: f.Nascimento}
}

func ProdutoToProdutoViewForm(c model.Produto) view.ProdutoViewForm {
	return view.ProdutoViewForm{
		Nome:       c.Nome,
		Sobrenome:  c.Sobrenome,
		Codigo:     c.Codigo,
		Nascimento: c.Nascimento,
	}
}

func (contrlr ProdutoDaoController) Create() error {
	var f view.ProdutoViewForm = contrlr.view.Create()
	c := ProdutoViewFormToProduto(f)
	contrlr.model.Create(&c)
	return nil
}

func (contrlr ProdutoDaoController) RequestCODIGO() (int, error) {
	produtos, _ := contrlr.model.GetAll()
	var forms []view.ProdutoViewForm
	for _, x := range produtos {
		forms = append(forms, ProdutoToProdutoViewForm(x))
	}
	Codigo := contrlr.view.RequestCODIGO(forms)
	return Codigo, nil
}
func (contrlr ProdutoDaoController) Update() error {
	Codigo, _ := contrlr.RequestCODIGO()
	produto, _ := contrlr.model.GetById(Codigo)
	form := ProdutoToProdutoViewForm(produto)
	outform, _ := contrlr.view.Update(form)
	outproduto := ProdutoViewFormToProduto(outform)
	contrlr.model.Update(Codigo, &outproduto)
	return nil
}

func (contrlr ProdutoDaoController) Delete() error {
	Codigo, _ := contrlr.RequestCODIGO()
	produto, _ := contrlr.model.GetById(Codigo)
	i, _ := contrlr.model.GetIndex(&produto)
	contrlr.model.Delete(i)
	return nil
}

func (contrlr ProdutoDaoController) ListAll() error {
	produtos, _ := contrlr.model.GetAll()
	var forms []view.ProdutoViewForm
	for _, x := range produtos {
		forms = append(forms, ProdutoToProdutoViewForm(x))
	}
	contrlr.view.VisualizeAll(forms)
	return nil
}
