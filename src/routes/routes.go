package routes

import (
	"livraria_digital/src/config"
	"livraria_digital/src/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	config.InitDB()
	db := config.DB

	ConfigurarRotas(r, db)

	middlewares.Configurar(r)
	r.Use(gin.Recovery())
	return r
}

// ConfigurarRotas: configura as rotas
func ConfigurarRotas(r *gin.Engine, db *gorm.DB) {
	ConfigurarLogin(r, db)
	ConfigurarUsuarios(r, db)
	ConfigurarLivros(r, db)
	ConfigurarAutor(r, db)
	ConfigurarEditora(r, db)
	ConfigurarCategoria(r, db)

}
