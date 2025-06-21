package controllers

import (
	"net/http"

	"github.com/Muhandis2000/online-school/internal/models"
	"github.com/Muhandis2000/online-school/internal/services"
	"github.com/gin-gonic/gin"
)

type CourseController struct {
	service *services.CourseService
}

func NewCourseController(service *services.CourseService) *CourseController {
	return &CourseController{service}
}

func (c *CourseController) GetCourses(ctx *gin.Context) {
	courses, err := c.service.GetCourses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}
	ctx.JSON(http.StatusOK, courses)
}

func (c *CourseController) CreateCourse(ctx *gin.Context) {
	var course models.Course
	if err := ctx.BindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id, err := c.service.CreateCourse(course)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	course.ID = id
	ctx.JSON(http.StatusCreated, course)
}
