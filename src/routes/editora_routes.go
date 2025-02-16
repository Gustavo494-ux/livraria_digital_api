package routes

import (
	"livraria_digital/src/controllers"
	"livraria_digital/src/models"
	"livraria_digital/src/repository"
	"livraria_digital/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ConfigurarEditora: Realiza a configurações necessárias para a utilização do crud
func ConfigurarEditora(r *gin.Engine, db *gorm.DB) {
	db.AutoMigrate(&models.Editora{})

	var repo = repository.NewGenericRepository[models.Editora](db)
	service := services.NewGenericServices(repo)
	controller := controllers.NewGenericController(service)

	groupo := r.Group("/editora")
	{

		groupo.GET("", controller.BuscarTodos)
		groupo.POST("", controller.Criar)
		groupo.GET("/:id", controller.BuscarPorId)
		groupo.PUT("/:id", controller.Atualizar)
		groupo.DELETE("/:id", controller.Deletar)
	}
}
