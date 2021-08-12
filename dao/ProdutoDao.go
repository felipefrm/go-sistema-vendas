package dao
import(
model "github.com/felipefrm/go-sistema-vendas/model"
)

type ProdutoIndexType = int

type ProdutoDao interface{
	Create(u *model.Produto) error
	Update(i ProdutoIndexType, u *model.Produto) error
	Delete(i ProdutoIndexType) error
	GetIndex(u *model.Produto) (ProdutoIndexType, error)
	GetById(i ProdutoIndexType) (model.Produto, error)
	GetAll() ([]model.Produto, error)
}

type ProdutoDaoMap struct{
	produtos map[ProdutoIndexType]model.Produto
}


func (dao ProdutoDaoMap) Create(u *model.Produto) error{
	dao.produtos[u.Codigo] = *u
	return nil
}

func (dao ProdutoDaoMap) Update(i ProdutoIndexType, u *model.Produto) error{
	delete(dao.produtos,i)
	dao.produtos[u.Codigo] = *u
	return nil
}

func (dao ProdutoDaoMap) Delete(i ProdutoIndexType) error{
	delete(dao.produtos,i)
	return nil
}

func (dao ProdutoDaoMap) GetIndex(u *model.Produto) ProdutoIndexType{
	return u.Codigo
}

func (dao ProdutoDaoMap) GetById(i ProdutoIndexType) (model.Produto, error){
	return dao.produtos[i],nil
}

func (dao ProdutoDaoMap) GetAll() ([]model.Produto, error){
	v := make([]model.Produto, 0, len(dao.produtos))

	for  _, value := range dao.produtos {
		 v = append(v, value)
	}
	return v, nil
}
