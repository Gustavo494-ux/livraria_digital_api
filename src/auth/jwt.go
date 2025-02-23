package auth

import (
	"errors"
	"fmt"
	"livraria_digital/src/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// CriarTokenJWT retorna um token assinado com as permissões do usuário
func CriarTokenJWT(usuarioID uint, permissoes map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Define as claims (permissões) do token
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix() // Expira em 6 horas
	claims["usuarioId"] = usuarioID

	// Adiciona as permissões personalizadas ao token
	for key, value := range permissoes {
		claims[key] = value
	}

	// Assina o token com a chave JWT
	return token.SignedString([]byte(config.CHAVE_JWT))
}

// ValidarToken verifica se o token passado na requisição é válido
func ValidarToken(c *gin.Context) error {
	tokenString := ExtrairToken(c)
	_, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}
	return nil
}

// ExtrairUsuarioID retorna o usuarioId que está salvo no token
func ExtrairUsuarioID(c *gin.Context) (uint, error) {
	usuarioID, err := ExtrairPermissao(c, "usuarioId")
	if err != nil {
		return 0, err
	}

	usuarioIDuInt, ok := usuarioID.(uint)
	if !ok {
		return 0, errors.New("usuarioId no token não é um número válido")
	}

	return usuarioIDuInt, nil
}

// ExtrairPermissao retorna o valor de uma permissão específica do token
func ExtrairPermissao(c *gin.Context, chave string) (interface{}, error) {
	tokenString := ExtrairToken(c)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return nil, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if valor, existe := permissoes[chave]; existe {
			return valor, nil
		}
		return nil, fmt.Errorf("permissão '%s' não encontrada no token", chave)
	}

	return nil, errors.New("token inválido")
}

// ExtrairToken: Extrai o Token da requisição
func ExtrairToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}

	return []byte(config.CHAVE_JWT), nil
}
