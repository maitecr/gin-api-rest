package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maitecr/api-go-gin/controllers"
)

func HandleRequests() {

	r := gin.Default()

	r.GET("/alunos", controllers.GetAlunos)
	r.GET("/:nome", controllers.GetHome)

	r.Run(":5000")

}
