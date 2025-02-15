package interfaces

type Services[T any] interface {
	Validar(valor *T) (err error)
	Criar(valor *T) (err error)
	Atualizar(Id int, valor *T) (err error)
	BuscarTodos() (resultado []T, err error)
	BuscarPorId(Id int) (resultado T, err error)
	Deletar(id int) (err error)
}
