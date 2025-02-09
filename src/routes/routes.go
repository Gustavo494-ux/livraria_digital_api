package routes

import (
	"livraria_digital/src/config"
	"livraria_digital/src/controllers"
	"livraria_digital/src/middlewares"
	"livraria_digital/src/models"
	"livraria_digital/src/repository"
	"livraria_digital/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	config.InitDB()
	db := config.DB

	configurarLivros(r, db)

	middlewares.Configurar(r)
	r.Use(gin.Recovery())
	return r
}

// configurarLivros: Realiza a configurações necessárias para a utilização do crud de livros
func configurarLivros(r *gin.Engine, db *gorm.DB) {
	db.AutoMigrate(&models.Livro{})

	livroRepo := repository.NewLivroRepository(db)
	livroService := services.NewLivroService(livroRepo)
	livroController := controllers.NewLivroController(livroService)

	r.Group("/livros")
	{
		r.GET("", livroController.GetAllLivros)
		r.POST("", livroController.CreateLivro)
		r.GET("/:id", livroController.GetLivroByID)
		r.PUT("/:id", livroController.UpdateLivro)
		r.DELETE("/:id", livroController.DeleteLivro)
	}

}
