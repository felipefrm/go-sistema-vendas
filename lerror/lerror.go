package lerror

type NotFoundError struct {
	//Name string
}

func (e *NotFoundError) Error() string {
	return "Não encontrado."
}

type InvalidKeyError struct {
	//Name string
}

func (e InvalidKeyError) Error() string {
	return "Chave invalida."
}
