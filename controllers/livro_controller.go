package controllers

import (
	"api_golang/database"
	"api_golang/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MostrarLivro(c *gin.Context) {

	id := c.Param("id")

	newid, erro := strconv.Atoi(id)

	if erro != nil {
		c.JSON(400, gin.H{
			"error": "O id deve ser um inteiro",
		})

		return
	}

	db := database.GetDatabase()

	var livro models.Livro
	erro = db.First(&livro, newid).Error

	if erro != nil {
		c.JSON(400, gin.H{
			"error": "Nenhum livro encontrado :" + erro.Error(),
		})

		return
	}

	c.JSON(200, livro)

}

func CriarLivro(c *gin.Context) {

	db := database.GetDatabase()

	var livro models.Livro

	erro := c.ShouldBindJSON(&livro)

	if erro != nil {
		c.JSON(400, gin.H{
			"error": "Nenhum Json encontrado :" + erro.Error(),
		})

		return
	}

	erro = db.Create(&livro).Error

	if erro != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível criar o livro  :" + erro.Error(),
		})

		return
	}

	c.JSON(200, livro)
}

func MostraLivros(c *gin.Context) {
	db := database.GetDatabase()

	var livros []models.Livro
	erro := db.Find(&livros).Error

	if erro != nil {
		c.JSON(400, gin.H{
			"error": "Nenhum livro cadastrado  :" + erro.Error(),
		})

		return
	}

	c.JSON(200, livros)
}

func EditarLivro(c *gin.Context) {
	db := database.GetDatabase()

	var livro models.Livro

	erro := c.ShouldBindJSON(&livro)

	if erro != nil {
		c.JSON(400, gin.H{
			"error": "Nenhum Json encontrado :" + erro.Error(),
		})

		return
	}

	erro = db.Save(&livro).Error

	if erro != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível atualizar o livro  :" + erro.Error(),
		})

		return
	}

	c.JSON(200, livro)
}

func DeleteLivro(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID deve ser um inteiro",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Livro{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possível deletar o livro: " + err.Error(),
		})
		return
	}

	c.Status(204)
}
