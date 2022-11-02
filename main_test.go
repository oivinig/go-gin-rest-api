package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/oivinig/api-go-gin/controllers"
	"github.com/oivinig/api-go-gin/database"
)

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func TestGreeting(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:student_name", controllers.Greeting)
	req, _ := http.NewRequest("GET", "/name", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "Should be equal.")
	responseMock := `{"API":"Hello name, whats up?"}`
	responseBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, responseMock, string(responseBody))
}

func TestFindAllHandler(t *testing.T) {
	database.Connect()
	r := SetupTestRoutes()
	r.GET("/students", controllers.DisplayStudents)
	req, _ := http.NewRequest("GET", "/students", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}
