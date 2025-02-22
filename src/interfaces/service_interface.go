package interfaces

import "livraria_digital/src/models"

type Services[T any] interface {
	Validar(valor *T) (err error)
	Criar(valor *T) (err error)
	BuscarTodos(paginacao models.Paginacao) (resultado models.ResultadoPaginado[T], err error)
	BuscarPorId(Id int) (resultado T, err error)
	Atualizar(Id int, valor *T) (err error)
	Deletar(id int) (err error)
}
