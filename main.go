package main

import (
	"livraria_digital/src/config"
	"livraria_digital/src/routes"
)

func main() {
	config.ConfigurarVariaveis()
	routes.InicializarAPI()
}
