package routes

import (
	"livraria_digital/src/controllers"
	"livraria_digital/src/models"
	"livraria_digital/src/repository"
	"livraria_digital/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ConfigurarLogin: Realiza a configurações necessárias para a utilização do login
func ConfigurarLogin(r *gin.Engine, db *gorm.DB) {
	db.AutoMigrate(&models.Usuario{})

	repo := repository.NewLoginRepository(db)
	service := services.NewLoginService(repo)
	controller := controllers.NewLoginController(service)

	login := r.Group("/login")
	{
		login.POST("", controller.RealizarLogin)
	}

}
