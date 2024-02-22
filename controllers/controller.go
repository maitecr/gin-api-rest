package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maitecr/api-go-gin/database"
	"github.com/maitecr/api-go-gin/models"
)

func GetHome(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz: ": "E aí " + nome + ", tudo bom?",
	})
}

func CreateAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func GetAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func GetAlunoById(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
	}

	c.JSON(http.StatusOK, aluno)
}

func DeleteAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

func EditAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName(("id"))
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)

	c.JSON(http.StatusOK, aluno)
}

func SearchByCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")

	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
	}

	c.JSON(http.StatusOK, aluno)
}
