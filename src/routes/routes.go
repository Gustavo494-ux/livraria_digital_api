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

	configurarUsuarios(r, db)
	configurarLivros(r, db)

	middlewares.Configurar(r)
	r.Use(gin.Recovery())
	return r
}

// configurarUsuarios: Realiza a configurações necessárias para a utilização do crud de usuários
func configurarUsuarios(r *gin.Engine, db *gorm.DB) {
	db.AutoMigrate(&models.Usuario{})

	var repo = repository.NewUsuarioRepository(db).ToGenericRepository()
	service := services.NewUsuarioService(repo).ToGenericService()
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

// configurarLivros: Realiza a configurações necessárias para a utilização do crud de livros
func configurarLivros(r *gin.Engine, db *gorm.DB) {
	db.AutoMigrate(&models.Livro{})

	livroRepo := repository.NewLivroRepository(db)
	livroService := services.NewLivroService(livroRepo)
	livroController := controllers.NewLivroController(livroService)

	livros := r.Group("/livros")
	{
		livros.GET("", livroController.GetAllLivros)
		livros.POST("", livroController.CreateLivro)
		livros.GET("/:id", livroController.GetLivroByID)
		livros.PUT("/:id", livroController.UpdateLivro)
		livros.DELETE("/:id", livroController.DeleteLivro)
	}

}
