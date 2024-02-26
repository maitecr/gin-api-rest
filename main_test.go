package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/maitecr/api-go-gin/controllers"
	"github.com/maitecr/api-go-gin/database"
	"github.com/maitecr/api-go-gin/models"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CreateAlunoMock() {
	aluno := models.Aluno{Nome: "Dracula", CPF: "12345678910", RG: "123456789"}
	database.DB.Create(&aluno)

	ID = int(aluno.ID)
}

func DeleteAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
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

func TestAlunosHandles(t *testing.T) {
	database.ConectDB()
	CreateAlunoMock()
	defer DeleteAlunoMock()

	r := SetupTestRoutes()
	r.GET("/alunos", controllers.GetAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

	fmt.Println(resposta.Body)
}

func TestSearchAlunoCPF(t *testing.T) {
	database.ConectDB()
	CreateAlunoMock()
	defer DeleteAlunoMock()

	r := SetupTestRoutes()
	r.GET("/alunos/cpf/:cpf", controllers.SearchByCpf)

	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678910", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestGetAlunoByIdHandler(t *testing.T) {
	database.ConectDB()
	CreateAlunoMock()
	defer DeleteAlunoMock()

	r := SetupTestRoutes()
	r.GET("/alunos/:id", controllers.GetAlunoById)

	pathSearch := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("GET", pathSearch, nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno

	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	fmt.Println(alunoMock.Nome)

	assert.Equal(t, "Dracula", alunoMock.Nome, "nomes devem ser iguais")
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeleteAlunoMock(t *testing.T) {
	database.ConectDB()
	CreateAlunoMock()

	r := SetupTestRoutes()
	r.DELETE("/alunos/:id", controllers.DeleteAluno)

	pathSearch := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathSearch, nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditAluno(t *testing.T) {
	database.ConectDB()
	CreateAlunoMock()
	defer DeleteAlunoMock()

	r := SetupTestRoutes()
	r.PATCH("/alunos/:id", controllers.EditAluno)

	aluno := models.Aluno{Nome: "Dracula", CPF: "10987654321", RG: "987654321"}
	valorJson, _ := json.Marshal(aluno)

	pathEdit := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("PATCH", pathEdit, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)

	assert.Equal(t, "Dracula", alunoMockAtualizado.Nome)
	assert.Equal(t, "10987654321", alunoMockAtualizado.CPF)
	assert.Equal(t, "987654321", alunoMockAtualizado.RG)

	fmt.Println(alunoMockAtualizado.Nome)
	fmt.Println(alunoMockAtualizado.CPF)
	fmt.Println(alunoMockAtualizado.RG)

}
