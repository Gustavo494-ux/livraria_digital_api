package routes

import (
	"fmt"
	"livraria_digital/src/config"
	"livraria_digital/src/middlewares"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InicializarAPI() {
	server := SetupRoutes()
	server.ListenAndServe()
}

func SetupRoutes() *http.Server {
	router := gin.Default()

	config.InitDB()
	db := config.DB

	ConfigurarRotas(router, db)

	middlewares.Configurar(router)
	router.Use(gin.Recovery())

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.PORTA_API),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server
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
