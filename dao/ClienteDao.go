package dao
import(
model "github.com/felipefrm/go-sistema-vendas/model"
)

type ClienteDao interface{
	Create(u *model.Cliente) error
	Update(i ClienteIndexType, u *model.Cliente) error
	Delete(i ClienteIndexType) error
	GetIndex(u *model.Cliente) (ClienteIndexType, error)
	GetById(i ClienteIndexType) (model.Cliente, error)
	GetAll() ([]model.Cliente, error)
}

type ClienteDaoMap struct{
	clientes map[ClienteIndexType]model.Cliente
}
type ClienteIndexType = string

func (dao ClienteDaoMap) Create(u *model.Cliente) error{
	dao.clientes[u.Rg] = *u
	return nil
}

func (dao ClienteDaoMap) Update(i ClienteIndexType, u *model.Cliente) error{
	delete(dao.clientes,i)
	dao.clientes[u.Rg] = *u
	return nil
}

func (dao ClienteDaoMap) Delete(i ClienteIndexType) error{
	delete(dao.clientes,i)
	return nil
}

func (dao ClienteDaoMap) GetIndex(u *model.Cliente) ClienteIndexType{
	return u.Rg
}

func (dao ClienteDaoMap) GetById(i ClienteIndexType) (model.Cliente, error){
	return dao.clientes[i],nil
}

func (dao ClienteDaoMap) GetAll() ([]model.Cliente, error){
	v := make([]model.Cliente, 0, len(dao.clientes))

	for  _, value := range dao.clientes {
		 v = append(v, value)
	}
	return v, nil
}
