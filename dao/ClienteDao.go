package dao

import (
	lerror "github.com/felipefrm/go-sistema-vendas/lerror"
	model "github.com/felipefrm/go-sistema-vendas/model"
	errors "github.com/pkg/errors"
)

type ClienteDao interface {
	Create(u *model.Cliente) error
	Update(i ClienteIndexType, u *model.Cliente) error
	Delete(i ClienteIndexType) error
	GetIndex(u *model.Cliente) (ClienteIndexType, error)
	GetById(i ClienteIndexType) (model.Cliente, error)
	GetAll() ([]model.Cliente, error)
}

type ClienteDaoMap map[ClienteIndexType]model.Cliente

type ClienteIndexType = string

func (dao ClienteDaoMap) Create(u *model.Cliente) error {
	if u == nil {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Cliente inválido.")
	}
	dao[u.Rg] = *u
	return nil
}

func (dao ClienteDaoMap) Update(i ClienteIndexType, u *model.Cliente) error {
	if u == nil {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Cliente inválido.")
	} else if i == "" {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Indice do cliente não é válido.")
	} else if _, err := dao[i]; !err {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Cliente não encontrado.")
	}

	delete(dao, i)
	dao[u.Rg] = *u
	return nil
}

func (dao ClienteDaoMap) Delete(i ClienteIndexType) error {
	if _, err := dao[i]; !err {
		return errors.Wrap(lerror.InvalidKeyError{}, "Cliente não encontrado.")
	}
	delete(dao, i)
	return nil
}

func (dao ClienteDaoMap) GetIndex(u *model.Cliente) (ClienteIndexType, error) {
	if u == nil {
		return "", errors.Wrap(&lerror.InvalidKeyError{}, "Cliente inválido.")
	}
	return u.Rg, nil
}

func (dao ClienteDaoMap) GetById(i ClienteIndexType) (model.Cliente, error) {
	if i == "" {
		return model.Cliente{}, errors.Wrap(&lerror.InvalidKeyError{}, "Cliente inválido.")
	} else if _, err := dao[i]; !err {
		return model.Cliente{}, errors.Wrap(&lerror.InvalidKeyError{}, "Cliente não encontrado.")
	}
	return dao[i], nil
}

func (dao ClienteDaoMap) GetAll() ([]model.Cliente, error) {
	v := make([]model.Cliente, 0, len(dao))

	for _, value := range dao {
		v = append(v, value)
	}
	return v, nil
}
