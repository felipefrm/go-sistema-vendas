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

type ProdutoDaoMap map[ProdutoIndexType]model.Produto

func (dao ProdutoDaoMap) Create(u *model.Produto) error {
	dao[u.Codigo] = *u
	return nil
}

func (dao ProdutoDaoMap) Update(i ProdutoIndexType, u *model.Produto) error {
	delete(dao, i)
	dao[u.Codigo] = *u
	return nil
}

func (dao ProdutoDaoMap) Delete(i ProdutoIndexType) error {
	delete(dao, i)
	return nil
}

func (dao ProdutoDaoMap) GetIndex(u *model.Produto) (ProdutoIndexType, error) {
	return u.Codigo, nil
}

func (dao ProdutoDaoMap) GetById(i ProdutoIndexType) (model.Produto, error) {
	return dao[i], nil
}

func (dao ProdutoDaoMap) GetAll() ([]model.Produto, error) {
	v := make([]model.Produto, 0, len(dao))

	for _, value := range dao {
		v = append(v, value)
	}
	return v, nil
}
