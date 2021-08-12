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
	return model.Produto{Codigo: f.Codigo, Nome: f.Nome, Valor: f.Valor}
}

func ProdutoToProdutoViewForm(produto model.Produto) view.ProdutoViewForm {
	return view.ProdutoViewForm{
		Codigo: produto.Codigo,
		Nome:   produto.Nome,
		Valor:  produto.Valor,
	}
}

func (contrlr ProdutoDaoController) Create() error {
	var f view.ProdutoViewForm
	f, _ = contrlr.view.Create()
	produto := ProdutoViewFormToProduto(f)
	contrlr.model.Create(&produto)
	return nil
}

func (contrlr ProdutoDaoController) RequestCodigo() (int, error) {
	produtos, _ := contrlr.model.GetAll()
	var forms []view.ProdutoViewForm
	for _, x := range produtos {
		forms = append(forms, ProdutoToProdutoViewForm(x))
	}
	codigo, _ := contrlr.view.RequestCodigo(forms)
	return codigo, nil
}
func (contrlr ProdutoDaoController) Update() error {
	codigo, _ := contrlr.RequestCodigo()
	produto, _ := contrlr.model.GetById(codigo)
	form := ProdutoToProdutoViewForm(produto)
	outform, _ := contrlr.view.Update(form)
	outproduto := ProdutoViewFormToProduto(outform)
	contrlr.model.Update(codigo, &outproduto)
	return nil
}

func (contrlr ProdutoDaoController) Delete() error {
	codigo, _ := contrlr.RequestCodigo()
	produto, _ := contrlr.model.GetById(codigo)
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
