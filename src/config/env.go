package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func CarregarEnv() {
	caminho, err := os.Getwd()
	if err != nil {
		fmt.Printf("Ocorreu um erro ao buscar o diretorio principal, erro: %s\n", err.Error())
	}

	err = godotenv.Load(fmt.Sprintf("%s/.env", caminho))
	if err != nil {
		fmt.Printf("Ocorreu um erro ao carregar o arquivo env, error: %s\n", err.Error())
	}
}
