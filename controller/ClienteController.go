package controller
import(
	model "github.com/felipefrm/go-sistema-vendas/model"
	view "github.com/felipefrm/go-sistema-vendas/view"
	dao "github.com/felipefrm/go-sistema-vendas/dao"
)

//type ClienteDao interface{
//Create(u *model.Cliente) error
//Update(i ClienteIndexType, u *model.Cliente) error
//Delete(i ClienteIndexType) error
//GetIndex(u *model.Cliente) (ClienteIndexType, error)
//GetById(i ClienteIndexType) (model.Cliente, error)
//GetAll() ([]model.Cliente, error)
//}

type ClienteDaoController struct{
	model dao.ClienteDao
	view view.ClienteView
}

func ClienteViewFormToCliente(f view.ClienteViewForm) model.Cliente{
	return model.Cliente{Pessoa: model.Pessoa{Nome: f.Nome, Sobrenome: f.Sobrenome},Rg: f.Rg, Nascimento: f.Nascimento}
}

func (contrlr ClienteDaoController) Create() error{
	var f view.ClienteViewForm = contrlr.view.Create()
	c := ClienteViewFormToCliente(f)
	contrlr.model.Create(&c)
	return nil
}

func (contrlr ClienteDaoController) Update() error{
	var f view.ClienteViewForm = contrlr.view.Create()
	//var i int
	c := ClienteViewFormToCliente(f)
	i, _ := contrlr.model.GetIndex(&c)
	contrlr.model.Update(i, &c)
	return nil
}

func (contrlr ClienteDaoController) Delete() error{
	var f view.ClienteViewForm = contrlr.view.Create()
	//var i int
	c := ClienteViewFormToCliente(f)
	i, _ := contrlr.model.GetIndex(&c)
	contrlr.model.Delete(i)
	return nil
}

func (contrlr ClienteDaoController) ListAll() error{
	return nil
}
