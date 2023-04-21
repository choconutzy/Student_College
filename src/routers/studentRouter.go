package routers

import (
	"Jobhun_Mahasiswa/src/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome",
		})
	})

	router.GET("/hobbies", controllers.GetHobbies)
	router.GET("/majors", controllers.GetMajors)
	router.GET("/students", controllers.GetStudents)
	router.POST("/student", controllers.CreateStudent)
	router.GET("/student/:id", controllers.GetStudentById)
	router.PUT("/student/:id", controllers.UpdateStudent)
	router.DELETE("/student/:id", controllers.DeleteStudent)

	return router
}
