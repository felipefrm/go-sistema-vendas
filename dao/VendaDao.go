package dao

import (
	lerror "github.com/felipefrm/go-sistema-vendas/lerror"
	model "github.com/felipefrm/go-sistema-vendas/model"
	errors "github.com/pkg/errors"
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

type VendaDaoMap struct {
	Model                map[VendaIndexType]model.Venda
	ClientesVendasNumero map[ClienteIndexType]map[VendaIndexType]bool
	ProdutosVendasNumero map[ProdutoIndexType]map[VendaIndexType]bool
}

func (dao VendaDaoMap) InsertClientesVendasNumero(i ClienteIndexType, j VendaIndexType, value bool) error {
	if i == "" {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Indice do cliente não é válido.")
	} else if j < 0 {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Indice de venda não é válido.")
	}

	if dao.ClientesVendasNumero[i] == nil {
		dao.ClientesVendasNumero[i] = make(map[VendaIndexType]bool)
	}
	dao.ClientesVendasNumero[i][j] = value
	return nil
}

func (dao VendaDaoMap) InsertProdutosVendasNumero(i ProdutoIndexType, j VendaIndexType, value bool) error {
	if i < 0 {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Indice do cliente não é válido.")
	} else if j < 0 {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Indice de venda não é válido.")
	}

	if dao.ProdutosVendasNumero[i] == nil {
		dao.ProdutosVendasNumero[i] = make(map[VendaIndexType]bool)
	}
	dao.ProdutosVendasNumero[i][j] = value
	return nil
}

func (dao VendaDaoMap) Create(u *model.Venda) error {
	if u == nil {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Venda inválida.")
	} else if _, err := dao.Model[(*u).Numero]; err {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Número de venda já em uso.")
	}

	dao.Model[u.Numero] = *u

	//dao.ClientesVendasNumero[u.Cliente.Rg][u.Numero] = true
	dao.InsertClientesVendasNumero(u.Cliente.Rg, u.Numero, true)

	for _, item := range u.Itens {
		dao.InsertProdutosVendasNumero(item.Produto.Codigo, u.Numero, true)
	}
	//dao.ClientesVendasNumero[u.Cliente.Rg]
	return nil
}

func (dao VendaDaoMap) Update(i VendaIndexType, u *model.Venda) error {
	if u == nil {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Venda inválida.")
	} else if i < 0 {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Indice da venda não é válido.")
	} else if _, err := dao.Model[i]; !err {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Venda não encontrada.")
	} else if _, err := dao.Model[(*u).Numero]; (*u).Numero != i && err {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Número de venda já em uso.")
	}

	i2, _ := dao.GetIndex(u)

	dao.Delete(i)
	dao.Model[i2] = *u
	dao.InsertClientesVendasNumero(u.Cliente.Rg, i2, true)
	for _, item := range u.Itens {
		dao.InsertProdutosVendasNumero(item.Produto.Codigo, u.Numero, true)
	}
	return nil
}

func (dao VendaDaoMap) Delete(i VendaIndexType) error {
	if _, err := dao.Model[i]; !err {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Venda não encontrado.")
	}

	u, _ := dao.GetById(i)
	delete(dao.ClientesVendasNumero[u.Cliente.Rg], i)
	for _, item := range u.Itens {
		delete(dao.ProdutosVendasNumero[item.Produto.Codigo], u.Numero)
	}
	delete(dao.Model, i)
	return nil
}

func (dao VendaDaoMap) GetIndex(u *model.Venda) (VendaIndexType, error) {
	if u == nil {
		return -1, errors.Wrap(&lerror.InvalidKeyError{}, "Venda inválida.")
	}

	return u.Numero, nil
}

func (dao VendaDaoMap) GetById(i VendaIndexType) (model.Venda, error) {
	if i < 0 {
		return model.Venda{}, errors.Wrap(&lerror.InvalidKeyError{}, "Venda inválida.")
	} else if _, err := dao.Model[i]; !err {
		return model.Venda{}, errors.Wrap(&lerror.InvalidKeyError{}, "Venda não encontrada.")
	}

	return dao.Model[i], nil
}

func (dao VendaDaoMap) GetAll() ([]model.Venda, error) {
	v := make([]model.Venda, 0, len(dao.Model))

	for _, value := range dao.Model {
		v = append(v, value)
	}
	return v, nil
}

func (dao VendaDaoMap) ClienteRemove(i ClienteIndexType) error {
	if i == "" {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Indice do cliente não é válido.")
	} else if _, err := dao.ClientesVendasNumero[i]; !err {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Cliente não encontrado.")
	}

	for idx := range dao.ClientesVendasNumero[i] {
		dao.Delete(idx)
	}
	return nil
}
func (dao VendaDaoMap) UpdateClienteKey(ckey ClienteIndexType, newkey ClienteIndexType) error {
	if ckey == "" {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Indice do cliente não é válido.")
	} else if _, err := dao.ClientesVendasNumero[ckey]; !err {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Cliente não encontrado.")
	} else if _, err := dao.ClientesVendasNumero[newkey]; !err {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Key já em uso.")
	}

	for idx := range dao.ClientesVendasNumero[ckey] {
		delete(dao.ClientesVendasNumero[ckey], idx)
		dao.InsertClientesVendasNumero(newkey, idx, true)
	}
	return nil
}

func (dao VendaDaoMap) ProdutoRemove(i ProdutoIndexType) error {
	if i < 0 {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Indice do produto não é válido.")
	} else if _, err := dao.Model[i]; !err {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Produto não encontrado.")
	}

	for idx := range dao.ProdutosVendasNumero[i] {
		dao.Delete(idx)
	}
	return nil
}

func (dao VendaDaoMap) UpdateProdutoKey(ckey ProdutoIndexType, newkey ProdutoIndexType) error {
	if ckey < 0 {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Indice do produto não é válido.")
	} else if _, err := dao.Model[ckey]; !err {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Produto não encontrado.")
	} else if _, err := dao.Model[newkey]; !err {
		return errors.Wrap(&lerror.InvalidKeyError{}, "Key já em uso.")
	}

	for idx := range dao.ProdutosVendasNumero[ckey] {
		delete(dao.ProdutosVendasNumero[ckey], idx)
		dao.InsertProdutosVendasNumero(newkey, idx, true)
	}
	return nil
}
