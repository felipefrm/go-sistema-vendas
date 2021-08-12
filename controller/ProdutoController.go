package controller
import(
model "github.com/felipefrm/go-sistema-vendas/model"
view "github.com/felipefrm/go-sistema-vendas/view"
dao "github.com/felipefrm/go-sistema-vendas/dao"
)

//type ProdutoDao interface{
	//Create(u *model.Produto) error
	//Update(i ProdutoIndexType, u *model.Produto) error
	//Delete(i ProdutoIndexType) error
	//GetIndex(u *model.Produto) (ProdutoIndexType, error)
	//GetById(i ProdutoIndexType) (model.Produto, error)
	//GetAll() ([]model.Produto, error)
//}

type ProdutoDaoController struct{
	model dao.ProdutoDao
	view view.ProdutoView
}

func ProdutoViewFormToProduto(f view.ProdutoViewForm) model.Produto{
	return model.Produto{Pessoa: model.Pessoa{Nome: f.Nome, Sobrenome: f.Sobrenome},Rg: f.Rg, Nascimento: f.Nascimento}
}

func (contrlr ProdutoDaoController) Create() error{
	var f view.ProdutoViewForm = contrlr.view.Create()
	c := ProdutoViewFormToProduto(f)
	contrlr.model.Create(&c)
	return nil
}

func (contrlr ProdutoDaoController) Update() error{
	var f view.ProdutoViewForm = contrlr.view.Create()
	//var i int
	c := ProdutoViewFormToProduto(f)
	i, _ := contrlr.model.GetIndex(&c)
	contrlr.model.Update(i, &c)
	return nil
}

func (contrlr ProdutoDaoController) Delete() error{
	var f view.ProdutoViewForm = contrlr.view.Create()
	//var i int
	c := ProdutoViewFormToProduto(f)
	i, _ := contrlr.model.GetIndex(&c)
	contrlr.model.Delete(i)
	return nil
}

func (contrlr ProdutoDaoController) ListAll() error{
	return nil
}
