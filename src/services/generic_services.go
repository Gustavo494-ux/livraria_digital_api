package services

import (
	"errors"
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
func (s *GenericServices[T]) validar(valor *T) (err error) {
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
func (s *GenericServices[T]) Criar(valor *T) error {
	if err := s.validar(valor); err != nil {
		return err
	}
	return s.Repo.Criar(valor)
}

// BuscarTodos: busca todos os registros no banco de dados
func (s *GenericServices[T]) BuscarTodos() ([]T, error) {
	return s.Repo.BuscarTodos()
}

// Buscar: busca todos os registros no banco de dados que cumprirem os requisitos dos parametros
func (s *GenericServices[T]) Buscar(valor T) ([]T, error) {
	return s.Repo.Buscar(valor)
}

// Buscar: busca o primeiro registro que cumprir os requisitos
func (s *GenericServices[T]) BuscarPrimeiro(valor T) (T, error) {
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

	return s.Repo.Atualizar(valor)
}

// Deletar: remove um registro do banco de dados pelo ID
func (s *GenericServices[T]) Deletar(Id int) error {
	var instancia T
	Generics.SetarCampo(&instancia, "ID", uint(Id))

	return s.Repo.Deletar(instancia)
}
