package controller

import (
	"fmt"

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
	if err := contrlr.model.Create(&produto); err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	return nil
}

func (contrlr ProdutoDaoController) RequestCodigo() (int, error) {
	produtos, _ := contrlr.model.GetAll()
	var forms []view.ProdutoViewForm
	for _, x := range produtos {
		forms = append(forms, ProdutoToProdutoViewForm(*x))
	}
	codigo, _ := contrlr.view.RequestCodigo(forms)
	return codigo, nil
}
func (contrlr ProdutoDaoController) Update() error {
	codigo, _ := contrlr.RequestCodigo()
	produto, err := contrlr.model.GetById(codigo)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	form := ProdutoToProdutoViewForm(*produto)
	outform, _ := contrlr.view.Update(form)
	outproduto := ProdutoViewFormToProduto(outform)
	if err := contrlr.model.Update(codigo, &outproduto); err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	return nil
}

func (contrlr ProdutoDaoController) Delete() error {
	codigo, _ := contrlr.RequestCodigo()
	produto, err := contrlr.model.GetById(codigo)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	i, err := contrlr.model.GetIndex(produto)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	if err := contrlr.model.Delete(i); err != nil {
		fmt.Printf("%v", err.Error())
		return err
	}
	return nil
}

func (contrlr ProdutoDaoController) ListAll() error {
	produtos, _ := contrlr.model.GetAll()
	var forms []view.ProdutoViewForm
	for _, x := range produtos {
		forms = append(forms, ProdutoToProdutoViewForm(*x))
	}
	contrlr.view.VisualizeAll(forms)
	return nil
}

func (contrlr ProdutoDaoController) OptionsMenu() error {
	for {
		option := contrlr.view.OptionsMenu()
		switch option {
		case view.Op????oSairProduto:
			return nil
		case view.Op????oVerProdutos:
			contrlr.ListAll()
		case view.Op????oAdicionarProduto:
			contrlr.Create()
		case view.Op????oAlterarProduto:
			contrlr.Update()
		case view.Op????oRemoverProduto:
			contrlr.Delete()
		}
	}
}
