package controller
import(
	model "github.com/felipefrm/go-sistema-vendas/model"
	view "github.com/felipefrm/go-sistema-vendas/view"
	dao "github.com/felipefrm/go-sistema-vendas/dao"
)

type MainDaoController struct{
	view view.MainView
}

func (contrlr MainDaoController) OptionsMenu() error{
	clientedao := dao.ClienteDaoMap{}
	produtodao := dao.ProdutoDaoMap{}
	vendadao := dao.VendaDaoMap{}
	clientecontroller := ClienteDaoController{model: clientedao, view: view.ClienteView{}}
	produtocontroller := ClienteDaoController{model: clientedao, view: view.ClienteView{}}
	for {
		option := contrlr.view.OptionsMenu()
		switch option{
		case view.OpçãoSair:
			return nil
		case view.OpçãoCliente:
			//ClienteDaoController{model: 
			//clientedao
			return nil
		case view.OpçãoProduto:
			return nil
		case view.OpçãoVenda:
			return nil
		}
	}
	return nil
}
