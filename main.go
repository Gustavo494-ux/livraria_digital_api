package main

import "livraria_digital/src/routes"

func main() {
	r := routes.SetupRoutes()
	r.Run(":8080")
}
