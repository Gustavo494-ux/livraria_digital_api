package routes

import (
	"livraria_digital/src/controllers"
	"livraria_digital/src/models"
	"livraria_digital/src/repository"
	"livraria_digital/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ConfigurarUsuarios: Realiza a configurações necessárias para a utilização do crud de usuários
func ConfigurarUsuarios(r *gin.Engine, db *gorm.DB) {
	db.AutoMigrate(&models.Usuario{})

	repo := repository.NewUsuarioRepository(db)
	service := services.NewUsuarioService(repo)
	controller := controllers.NewUsuarioController(service)

	usuarios := r.Group("/usuarios")
	{
		usuarios.GET("", controller.BuscarTodos)
		usuarios.POST("", controller.Criar)
		usuarios.GET("/:id", controller.BuscarPorId)
		usuarios.PUT("/:id", controller.Atualizar)
		usuarios.DELETE("/:id", controller.Deletar)
	}

}
