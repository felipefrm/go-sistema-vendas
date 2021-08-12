package model

//type ClienteIndexType string

type Cliente struct {
	Pessoa
	Rg         string
	Nascimento string
}

//func (c Cliente) Index() string {
	//return c.rg
//}
