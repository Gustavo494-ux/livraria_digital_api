package services

import (
	"errors"
	"livraria_digital/src/models"
	Generics "livraria_digital/src/pkg"
	"livraria_digital/src/repository"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// GenericServices é uma estrutura genérica para operações de CRUD
type GenericServices[T any] struct {
	Repo      *repository.GenericRepository[T]
	Validator *validator.Validate
}

// NewGenericServices cria uma nova instância do serviço genérico
func NewGenericServices[T any](repo *repository.GenericRepository[T]) *GenericServices[T] {
	return &GenericServices[T]{
		Repo:      repo,
		Validator: validator.New(),
	}
}

// validar: verifica se a entidade possui dados válidos e se não existe no banco
func (s *GenericServices[T]) Validar(valor *T) (err error) {
	if err := s.Validator.Struct(valor); err != nil {
		return err
	}

	entidadeBanco, err := s.Repo.BuscarPrimeiro(*valor)
	if err != nil {
		return err
	}

	if !reflect.ValueOf(entidadeBanco).IsZero() {
		return errors.New("registro existente")
	}

	return
}

// Criar: cria um novo registro no banco de dados
func (s *GenericServices[T]) Criar(valor *T) (err error) {
	if err := s.Validar(valor); err != nil {
		return err
	}
	return s.Repo.Criar(valor)
}

// BuscarTodos: busca todos os registros no banco de dados
func (s *GenericServices[T]) BuscarTodos(paginacao models.Paginacao) (resultado models.ResultadoPaginado[T], err error) {
	entidadeVazia := new(T)
	paginacao.Total, err = s.Repo.RetornarTotalRegistos(entidadeVazia)
	if err != nil {
		return models.ResultadoPaginado[T]{}, err
	}

	resultado, err = s.Repo.BuscarTodos(paginacao)
	resultado.Paginacao = paginacao
	resultado.Paginacao.CalcularQuantidadePaginas()

	return
}

// Buscar: busca o registro com Id fornecido
func (s *GenericServices[T]) BuscarPorId(Id int) (resultado T, err error) {
	var valor T
	Generics.SetarCampo(valor, "ID", Id)
	return s.Repo.BuscarPrimeiro(valor)

}

// Atualizar: atualiza um registro no banco de dados
func (s *GenericServices[T]) Atualizar(Id int, valor *T) (err error) {
	var registroBanco T
	Generics.SetarCampo(&registroBanco, "ID", uint(Id))

	registroBanco, err = s.Repo.BuscarPrimeiro(registroBanco)
	if err != nil {
		return err
	}
	if Generics.RetornarCampo(registroBanco, "ID", 0).(uint) < 1 {
		return errors.New("nenhum registro encontrado")
	}

	Generics.CopiarSeVazio(
		&registroBanco,
		valor,
	)

	return s.Repo.Atualizar(valor)
}

// Deletar: remove um registro do banco de dados pelo ID
func (s *GenericServices[T]) Deletar(Id int) error {
	var entidade T
	Generics.SetarCampo(&entidade, "ID", uint(Id))

	entidadeBanco, err := s.Repo.BuscarPrimeiro(entidade)
	if err != nil {
		return err
	}

	if (Generics.RetornarCampo(entidadeBanco, "ID", 0)).(uint) < 1 {
		return errors.New("registro não encontrado")
	}

	return s.Repo.Deletar(entidadeBanco)
}
