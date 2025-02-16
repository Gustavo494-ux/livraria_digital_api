package routes

import (
	"livraria_digital/src/controllers"
	"livraria_digital/src/models"
	"livraria_digital/src/repository"
	"livraria_digital/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ConfigurarLivros: Realiza a configurações necessárias para a utilização do crud de livros
func ConfigurarLivros(r *gin.Engine, db *gorm.DB) {
	db.AutoMigrate(&models.Livro{})

	repo := repository.NewLivroRepository(db).ToGenericRepository()
	service := services.NewLivroService(repo).ToGenericService()
	controller := controllers.NewLivroController(service)

	grupo := r.Group("/livros")
	{
		grupo.GET("", controller.BuscarTodos)
		grupo.POST("", controller.Criar)
		grupo.GET("/:id", controller.BuscarPorId)
		grupo.PUT("/:id", controller.Atualizar)
		grupo.DELETE("/:id", controller.Deletar)
	}
}
