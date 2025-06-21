package routes

import (
	"github.com/Muhandis2000/online-school/internal/controllers"
	"github.com/Muhandis2000/online-school/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authCtrl *controllers.AuthController, courseCtrl *controllers.CourseController) {
	r.POST("/register", authCtrl.Register)
	r.POST("/login", authCtrl.Login)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/courses", courseCtrl.GetCourses)
		api.POST("/courses", courseCtrl.CreateCourse)
	}
}
