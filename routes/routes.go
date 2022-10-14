package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oivinig/api-go-gin/controllers"
)

func Handler() {
	r := gin.Default()
	r.GET("/students", controllers.DisplayStudents)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.FindStudent)
	r.PATCH("/students/:id", controllers.EditStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.GET("/students/document/:document", controllers.FindStudentByDocument)
	r.GET("/:student_name", controllers.Greeting)
	r.Run()
}
