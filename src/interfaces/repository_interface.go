package interfaces

import "livraria_digital/src/models"

type Repository[T any] interface {
	Criar(valor *T) (err error)
	Atualizar(valor *T) (err error)
	RetornarTotalRegistos(filtro *T) (Total int64, err error)
	BuscarTodos(paginacao models.Paginacao) (resultado models.ResultadoPaginado[T], err error)
	Buscar(criterio T, paginacao models.Paginacao) (resultado models.ResultadoPaginado[T], err error)
	BuscarPrimeiro(valor T) (resultado T, err error)
	Deletar(valor T) (err error)
}
