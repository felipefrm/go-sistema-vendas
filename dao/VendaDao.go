package dao

import (
	model "github.com/felipefrm/go-sistema-vendas/model"
)

type VendaIndexType = int

type VendaDao interface {
	Create(u *model.Venda) error
	Update(i VendaIndexType, u *model.Venda) error
	Delete(i VendaIndexType) error
	GetIndex(u *model.Venda) (VendaIndexType, error)
	GetById(i VendaIndexType) (model.Venda, error)
	GetAll() ([]model.Venda, error)
}

type VendaDaoMap map[VendaIndexType]model.Venda

func (dao VendaDaoMap) Create(u *model.Venda) error {
	dao[u.Numero] = *u
	return nil
}

func (dao VendaDaoMap) Update(i VendaIndexType, u *model.Venda) error {
	delete(dao, i)
	dao[u.Numero] = *u
	return nil
}

func (dao VendaDaoMap) Delete(i VendaIndexType) error {
	delete(dao, i)
	return nil
}

func (dao VendaDaoMap) GetIndex(u *model.Venda) (VendaIndexType, error) {
	return u.Numero, nil
}

func (dao VendaDaoMap) GetById(i VendaIndexType) (model.Venda, error) {
	return dao[i], nil
}

func (dao VendaDaoMap) GetAll() ([]model.Venda, error) {
	v := make([]model.Venda, 0, len(dao))

	for _, value := range dao {
		v = append(v, value)
	}
	return v, nil
}
