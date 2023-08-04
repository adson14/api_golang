package routes

import (
	"api_golang/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRota(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		livros := main.Group("livros")
		{
			livros.GET("/:id", controllers.MostrarLivro)
			livros.GET("/", controllers.MostraLivros)
			livros.POST("/", controllers.CriarLivro)
			livros.PUT("/", controllers.EditarLivro)
			livros.DELETE("/:id", controllers.DeleteLivro)
		}
	}

	return router
}
