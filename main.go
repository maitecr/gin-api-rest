package main

import (
	"github.com/maitecr/api-go-gin/database"
	"github.com/maitecr/api-go-gin/models"
	"github.com/maitecr/api-go-gin/routes"
)

func main() {
	database.ConectDB()
	models.Alunos = []models.Aluno{
		{Nome: "nome", CPF: "515465", RG: "1556454"},
	}
	routes.HandleRequests()
}
