package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maitecr/api-go-gin/controllers"
)

func HandleRequests() {

	r := gin.Default()

	r.GET("/alunos", controllers.GetAlunos)
	r.GET("/:nome", controllers.GetHome)
	r.POST("/alunos", controllers.CreateAluno)
	r.GET("/alunos/:id", controllers.GetAlunoById)
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	r.PATCH("/alunos/:id", controllers.EditAluno)
	r.GET("/alunos/cpf/:cpf", controllers.SearchByCpf)

	r.Run(":5000")

}
