package config

import (
	"fmt"
	"os"
	"strconv"
)

var (
	CAMINHO_BANCO = ""
	PORTA_API     = 0
	CHAVE_JWT     = ""
)

// ConfigurarVariaveis: busca as variaveis definidas no .env e carrega como variaveis globais
func ConfigurarVariaveis() {
	CarregarEnv()

	CAMINHO_BANCO = os.Getenv("CAMINHO_BANCO")
	CHAVE_JWT = os.Getenv("CHAVE_JWT")

	var err error
	PORTA_API, err = strconv.Atoi(os.Getenv("PORTA_API"))
	if err != nil {
		fmt.Printf("ocorreu um erro ao configurar as variaveis, error: %s\n", err.Error())
	}

	fmt.Println(CAMINHO_BANCO)
	fmt.Println(PORTA_API)
	fmt.Println(CHAVE_JWT)
}
