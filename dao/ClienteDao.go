package dao

import (
	model "github.com/felipefrm/go-sistema-vendas/model"
	"errors"
	"fmt"
)

type ClienteDao interface {
	Create(u *model.Cliente) error
	Update(i ClienteIndexType, u *model.Cliente) error
	Delete(i ClienteIndexType) error
	GetIndex(u *model.Cliente) (ClienteIndexType, error)
	GetById(i ClienteIndexType) (model.Cliente, error)
	GetAll() ([]model.Cliente, error)
}

type ClienteDaoMap struct {
	clientes map[ClienteIndexType]model.Cliente
}
type ClienteIndexType = string

func (dao ClienteDaoMap) Create(u *model.Cliente) error {
	if u == nil{
		err := erros.New("Cliente não válido.")
		fmt.Print(err)
		return err
	}
	dao.clientes[u.Rg] = *u
	return nil
}

func (dao ClienteDaoMap) Update(i ClienteIndexType, u *model.Cliente) error {
	if u == nil{
		err := errors.New("Cliente não válido.")
		fmt.Print(err)
		return err
	}else if i == nil{
		err := errors.New("Indice não válido.")
		fmt.Print(err)
		return err
	}
	delete(dao.clientes, i)
	dao.clientes[u.Rg] = *u
	return nil
}

func (dao ClienteDaoMap) Delete(i ClienteIndexType) error {
	if i == nil{
		err := errors.New("Indice não válido.")
		fmt.Print(err)
		return err
	}
	delete(dao.clientes, i)
	return nil
}

func (dao ClienteDaoMap) GetIndex(u *model.Cliente) (ClienteIndexType, error) {
	if u == nil{
		err := errors.New("Cliente não válido.")
		fmt.Print(err)
		return err
	return u.Rg, nil
}

func (dao ClienteDaoMap) GetById(i ClienteIndexType) (model.Cliente, error) {
	if i == nil{
		err := errors.New("Indice não válido.")
		fmt.Print(err)
		return err
	}
	return dao.clientes[i], nil
}

func (dao ClienteDaoMap) GetAll() ([]model.Cliente, error) {
	v := make([]model.Cliente, 0, len(dao.clientes))

	for _, value := range dao.clientes {
		v = append(v, value)
	}
	return v, nil
}
