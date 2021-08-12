package view
import ("fmt")
//type ClienteDao interface{
	//Create(u *model.Cliente) error
	//Update(i ClienteIndexType, u *model.Cliente) error
	//Delete(i ClienteIndexType) error
	//GetIndex(u *model.Cliente) (ClienteIndexType, error)
	//GetById(i ClienteIndexType) (model.Cliente, error)
	//GetAll() ([]model.Cliente, error)
//}

type ClienteView struct{

}

type ClienteViewForm struct{
	Nome string
	Sobrenome string
	Rg string
	Nascimento string
}
func (c ClienteView) Create() ClienteViewForm{
	var form ClienteViewForm
	fmt.Printf("Nome:")
	fmt.Scanf("%s\n", &form.Nome)
	fmt.Printf("Sobrenome:")
	fmt.Scanf("%s\n", &form.Sobrenome)
	fmt.Printf("RG:")
	fmt.Scanf("%s\n", &form.Rg)
	fmt.Printf("Data de nascimento:")
	fmt.Scanf("%s\n", &form.Nascimento)
	return form
}

func (c ClienteView) Delete() ClienteViewForm{
	var form ClienteViewForm
	fmt.Printf("Digite o RG para remoção:")
	fmt.Scanf("%s\n", &form.Rg)
	return form
}
func (c ClienteView) RequestRG() ClienteViewForm{
	var form ClienteViewForm
	fmt.Printf("Digite o RG para alteração:")
	fmt.Scanf("%s\n", &form.Rg)
	return form
}
func (c ClienteView) Update(form ClienteViewForm) ClienteViewForm{
	//ClienteViewForm form
	var tmpinput string
	fmt.Printf("Nome (%s):",form.Nome)
	fmt.Scanf("%s\n", &tmpinput)
	if len(tmpinput) > 0 {
		form.Nome= tmpinput
	}
	fmt.Printf("Sobrenome (%s):",form.Sobrenome)
	fmt.Scanf("%s\n", &form.Sobrenome)
	fmt.Printf("RG (%s):",form.Rg)
	fmt.Scanf("%s\n", &form.Rg)
	fmt.Printf("Data de nascimento (%s):",form.Nascimento)
	fmt.Scanf("%s\n", &form.Nascimento)
	return form
}
