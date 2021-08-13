package dao

import (
	model "github.com/felipefrm/go-sistema-vendas/model"
)

type ProdutoIndexType = int

type ProdutoDao interface {
	Create(u *model.Produto) error
	Update(i ProdutoIndexType, u *model.Produto) error
	Delete(i ProdutoIndexType) error
	GetIndex(u *model.Produto) (ProdutoIndexType, error)
	GetById(i ProdutoIndexType) (model.Produto, error)
	GetAll() ([]model.Produto, error)
}

type ProdutoDaoMap struct {
	Model       map[ProdutoIndexType]model.Produto
	Vendadaomap *VendaDaoMap
}

func (dao ProdutoDaoMap) Create(u *model.Produto) error {
	dao.Model[u.Codigo] = *u
	return nil
}

func (dao ProdutoDaoMap) Update(i ProdutoIndexType, u *model.Produto) error {
	newindex, _ := dao.GetIndex(u)
	if newindex != i {
		dao.Vendadaomap.UpdateProdutoKey(i, newindex)
	}
	dao.Delete(i)
	dao.Model[u.Codigo] = *u
	return nil
}

func (dao ProdutoDaoMap) Delete(i ProdutoIndexType) error {
	dao.Vendadaomap.ProdutoRemove(i)
	delete(dao.Model, i)
	return nil
}

func (dao ProdutoDaoMap) GetIndex(u *model.Produto) (ProdutoIndexType, error) {
	return u.Codigo, nil
}

func (dao ProdutoDaoMap) GetById(i ProdutoIndexType) (model.Produto, error) {
	return dao.Model[i], nil
}

func (dao ProdutoDaoMap) GetAll() ([]model.Produto, error) {
	v := make([]model.Produto, 0, len(dao.Model))

	for _, value := range dao.Model {
		v = append(v, value)
	}
	return v, nil
}
