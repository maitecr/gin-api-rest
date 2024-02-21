package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maitecr/api-go-gin/models"
)

func GetAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func GetHome(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz: ": "E aí " + nome + ", tudo bom?",
	})
}
