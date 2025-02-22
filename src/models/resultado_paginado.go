package models

type ResultadoPaginado[T any] struct {
	Dados []*T `json:"dados"`
	Paginacao
}
