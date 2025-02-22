package repository

import (
	"livraria_digital/src/models"

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
func (r *GenericRepository[T]) BuscarTodos(paginacao models.Paginacao) (resultado models.ResultadoPaginado[T], err error) {
	err = r.DB.
		Limit(paginacao.Limite).
		Offset(paginacao.GetOffset()).
		Find(&resultado.Dados).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return models.ResultadoPaginado[T]{}, err
	}
	return
}

// RetornarTotalRegistos: retorna uma contagem do total de registros que atedem aos critérios fornecidos
func (r *GenericRepository[T]) RetornarTotalRegistos(filtro *T) (Total int64, err error) {
	err = r.DB.Model(filtro).Where(filtro).Count(&Total).Error
	return
}

// Buscar: busca registros que correspondem aos critérios fornecidos
func (r *GenericRepository[T]) Buscar(criterio T, paginacao models.Paginacao) (resultado models.ResultadoPaginado[T], err error) {
	err = r.DB.
		Where(criterio).
		Limit(paginacao.Limite).
		Offset(paginacao.GetOffset()).
		Find(resultado.Dados).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return models.ResultadoPaginado[T]{}, err
	}
	return
}

// BuscarPrimeiro: busca o primeiro registro que corresponde aos critérios fornecidos
func (r *GenericRepository[T]) BuscarPrimeiro(criterio T) (T, error) {
	var entity T
	err := r.DB.Where(criterio).First(&entity).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return entity, err
	}
	return entity, nil
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
