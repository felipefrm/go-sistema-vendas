package dao

import (
	model "github.com/felipefrm/go-sistema-vendas/model"
	lerror "github.com/felipefrm/go-sistema-vendas/lerror"
	errors "github.com/pkg/errors"
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
	if u == nil{
		return errors.Wrap(&lerror.InvalidKeyError{}, "Produto inválido.")
	}
	
	dao.Model[u.Codigo] = *u
	return nil
}

func (dao ProdutoDaoMap) Update(i ProdutoIndexType, u *model.Produto) error {
	if u == nil {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Produto inválido.")
	} else if i == "" {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Indice do produto não é válido.")
	} else if _, err := dao.Model[i]; !err {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Produto não encontrado.")
	}

	newindex, _ := dao.GetIndex(u)
	if newindex != i {
		dao.Vendadaomap.UpdateProdutoKey(i, newindex)
	}
	dao.Delete(i)
	dao.Model[u.Codigo] = *u
	return nil
}

func (dao ProdutoDaoMap) Delete(i ProdutoIndexType) error {
	if _, err := dao.Model[i]; !err {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Produto não encontrado.")
	}

	dao.Vendadaomap.ProdutoRemove(i)
	delete(dao.Model, i)
	return nil
}

func (dao ProdutoDaoMap) GetIndex(u *model.Produto) (ProdutoIndexType, error) {
	if u == nil {
		return "", errors.Wrap(&lerror.InvalidKeyError{}, "Produto inválido.")
	}

	return u.Codigo, nil
}

func (dao ProdutoDaoMap) GetById(i ProdutoIndexType) (model.Produto, error) {
	if i == "" {
		return model.Cliente{}, errors.Wrap(&lerror.InvalidKeyError{}, "Produto inválido.")
	} else if _, err := dao.Model[i]; !err {
		return model.Cliente{}, errors.Wrap(&lerror.InvalidKeyError{}, "Produto não encontrado.")
	}

	return dao.Model[i], nil
}

func (dao ProdutoDaoMap) GetAll() ([]model.Produto, error) {
	v := make([]model.Produto, 0, len(dao.Model))

	for _, value := range dao.Model {
		v = append(v, value)
	}
	return v, nil
}
