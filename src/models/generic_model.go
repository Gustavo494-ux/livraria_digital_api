package models

import (
	Generics "livraria_digital/src/pkg"

	"gorm.io/gorm"
)

type GenericModel[T any] struct {
	gorm.Model
}

// NewGenericModel: cria uma nova instância do model genérico
func NewGenericModel[T any]() *GenericModel[T] {
	return &GenericModel[T]{}
}

// IsID: verifica se o Id é válido. ID != 0
func (m *GenericModel[T]) IsID() bool {
	return m.ID != 0
}

// Merge: Copia os campos(recurssivamente) do parametro para a instância atual. desde que o campo esteja vazio no destino
func (m *GenericModel[T]) Merge(source *GenericModel[T]) {
	Generics.CopiarSeVazio(m, source)
}
