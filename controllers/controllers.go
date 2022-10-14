package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oivinig/api-go-gin/database"
	"github.com/oivinig/api-go-gin/models"
)

func DisplayStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("student_name")
	c.JSON(200, gin.H{
		"me": "hello you " + name + "! whats up?",
	})
}

func CreateStudent(c *gin.Context) {
	var newStudent models.Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&newStudent)
	c.JSON(http.StatusOK, newStudent)
}

func FindStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Student " + id + " deleted sucessfully.",
	})
}

func EditStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func FindStudentByDocument(c *gin.Context) {
	var student models.Student
	document := c.Param("document")
	database.DB.Where(&models.Student{Document: document}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	c.JSON(http.StatusOK, student)
}
