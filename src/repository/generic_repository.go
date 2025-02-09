package repository

import (
	"gorm.io/gorm"
)

// GenericRepository é uma estrutura genérica para operações de CRUD
type GenericRepository[T any] struct {
	DB *gorm.DB
}

// NewGenericRepository cria uma nova instância do repositório genérico
func NewGenericRepository[T any](db *gorm.DB) *GenericRepository[T] {
	return &GenericRepository[T]{DB: db}
}

// Criar: cria um novo registro no banco de dados
func (r *GenericRepository[T]) Criar(entity *T) error {
	return r.DB.Create(entity).Error
}

// BuscarTodos: busca todos os registros no banco de dados
func (r *GenericRepository[T]) BuscarTodos() ([]T, error) {
	var entities []T
	err := r.DB.Find(&entities).Error
	return entities, err
}

// Buscar: busca registros que correspondem aos critérios fornecidos
func (r *GenericRepository[T]) Buscar(criteria T) ([]T, error) {
	var entities []T
	err := r.DB.Where(criteria).Find(&entities).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return entities, nil
}

// BuscarPrimeiro: busca o primeiro registro que corresponde aos critérios fornecidos
func (r *GenericRepository[T]) BuscarPrimeiro(criteria T) (*T, error) {
	var entity T
	err := r.DB.Where(criteria).First(&entity).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &entity, nil
}

// Atualizar: atualiza um registro no banco de dados
func (r *GenericRepository[T]) Atualizar(entity *T) error {
	return r.DB.Save(entity).Error
}

// Deletar: remove um registro do banco de dados pelo ID
func (r *GenericRepository[T]) Deletar(entityParametro T) error {
	var entity T
	return r.DB.Delete(&entity, entityParametro).Error
}
