package main

import (
	"github.com/maitecr/api-go-gin/models"
	"github.com/maitecr/api-go-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "nome", CPF: "515465", RG: "1556454"},
	}
	routes.HandleRequests()
}
