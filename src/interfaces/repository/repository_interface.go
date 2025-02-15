package interfaces

type Repository[T any] interface {
	Criar(valor *T) (err error)
	Atualizar(valor *T) (err error)
	BuscarTodos() (resultado []T, err error)
	BuscarPrimeiro(valor T) (resultado T, err error)
	Deletar(valor T) (err error)
}
