package services

import (
	"errors"
	"livraria_digital/src/interfaces"
	"livraria_digital/src/models"
	Generics "livraria_digital/src/pkg"
	"livraria_digital/src/repository"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type IUsuarioService interface {
	interfaces.Services[models.Usuario]
}

type UsuarioService struct {
	repo      *repository.UsuarioRepository
	validator *validator.Validate
}

// NewUsuarioService cria uma nova instância do serviço de usuário
func NewUsuarioService(repo *repository.UsuarioRepository) *UsuarioService {
	return &UsuarioService{
		repo:      repo,
		validator: validator.New(),
	}
}

// Validar: verifica se o usuário possui dados válidos e se não existe no banco
func (s *UsuarioService) Validar(usuario *models.Usuario) (err error) {
	if err := s.validator.Struct(usuario); err != nil {
		return err
	}

	entidadeBanco, err := s.repo.BuscarPrimeiro(*usuario)
	if err != nil {
		return err
	}

	if !reflect.ValueOf(entidadeBanco).IsZero() {
		return errors.New("registro existente")
	}

	return
}

// Criar: cria um novo registro no banco de dados
func (s *UsuarioService) Criar(usuario *models.Usuario) (err error) {
	if err := s.Validar(usuario); err != nil {
		return err
	}

	usuario.EnriptarSenha()
	return s.repo.Criar(usuario)
}

// BuscarTodos: busca todos os registros no banco de dados
func (s *UsuarioService) BuscarTodos(paginacao models.Paginacao) (resultado models.ResultadoPaginado[models.Usuario], err error) {
	paginacao.Total, err = s.repo.RetornarTotalRegistos(&models.Usuario{})
	if err != nil {
		return models.ResultadoPaginado[models.Usuario]{}, err
	}

	resultado, err = s.repo.BuscarTodos(paginacao)
	resultado.Paginacao = paginacao
	resultado.Paginacao.CalcularQuantidadePaginas()

	return
}

// BuscarPorId: busca o registro com Id fornecido
func (s *UsuarioService) BuscarPorId(Id int) (resultado models.Usuario, err error) {
	usuarioFiltro := models.Usuario{}
	usuarioFiltro.ID = uint(Id)
	return s.repo.BuscarPrimeiro(usuarioFiltro)
}

// Atualizar: atualiza um registro no banco de dados
func (s *UsuarioService) Atualizar(Id int, usuarioParametro *models.Usuario) (err error) {
	registroBanco := models.Usuario{}
	registroBanco.ID = uint(Id)

	registroBanco, err = s.repo.BuscarPrimeiro(registroBanco)
	if err != nil {
		return err
	}

	if !registroBanco.IsID() {
		return errors.New("usuário não encontrado")
	}

	Generics.CopiarSeVazio(
		&registroBanco,
		usuarioParametro,
	)

	return s.repo.Atualizar(usuarioParametro)
}

// Deletar: remove um registro do banco de dados pelo ID
func (s *UsuarioService) Deletar(Id int) error {
	usuarioFiltro := models.Usuario{}
	usuarioFiltro.ID = uint(Id)

	usuarioBanco, err := s.repo.BuscarPrimeiro(usuarioFiltro)
	if err != nil {
		return err
	}

	if !usuarioBanco.IsID() {
		return errors.New("usuário não encontrado")
	}

	return s.repo.Deletar(usuarioBanco)
}
