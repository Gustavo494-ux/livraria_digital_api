package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// GerarHash: gera um hash seguro a partir de uma string
func GerarHash(texto string) (textoEncriptado string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(texto), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Erro ao gerar hash:", err)
		return
	}
	textoEncriptado = string(hash)
	return
}

func ValidarHash(texto, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(texto)) == nil
}
