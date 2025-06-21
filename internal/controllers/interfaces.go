package controllers

import "github.com/gin-gonic/gin"

type AuthControllerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type CourseControllerInterface interface {
	GetCourses(c *gin.Context)
	CreateCourse(c *gin.Context)
}
