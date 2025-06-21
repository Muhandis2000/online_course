package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/online-school/internal/controllers"
	"github.com/yourusername/online-school/internal/middleware"
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
