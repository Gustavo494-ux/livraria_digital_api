package routes

import (
	"livraria_digital/src/config"
	"livraria_digital/src/controllers"
	"livraria_digital/src/middlewares"
	"livraria_digital/src/models"
	"livraria_digital/src/repository"
	"livraria_digital/src/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Inicializa o banco de dados
	config.InitDB()
	db := config.DB

	// Auto-migrate (cria as tabelas no banco de dados)
	db.AutoMigrate(&models.Livro{})

	// Inicializa o repositório e o service
	livroRepo := repository.NewLivroRepository(db)
	livroService := services.NewLivroService(livroRepo)
	livroController := controllers.NewLivroController(livroService)

	// Rotas para Livros
	r.GET("/livros", livroController.GetAllLivros)
	r.GET("/livros/:id", livroController.GetLivroByID)
	r.POST("/livros", livroController.CreateLivro)
	r.PUT("/livros/:id", livroController.UpdateLivro)
	r.DELETE("/livros/:id", livroController.DeleteLivro)

	middlewares.Configurar(r)
	r.Use(gin.Recovery())
	return r
}
