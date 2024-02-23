package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/maitecr/api-go-gin/controllers"
)

func SetupTestRoutes() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestStatusCodeHome(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:nome", controllers.GetHome)
	req, _ := http.NewRequest("GET", "/batata", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code, "They should be equals")

	mockResp := `{"API diz": "E aí batata, tudo bom?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(t, mockResp, string(respostaBody))

	fmt.Println(string(respostaBody))
	fmt.Println(string(mockResp))
}
