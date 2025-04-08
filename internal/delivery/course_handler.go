package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-project/internal/models"
	"rest-project/internal/services"
)

type CourseHandler struct {
	Service *service.CourseService
}

func NewCourseHandler(service *service.CourseService) *CourseHandler {
	return &CourseHandler{Service: service}
}

func (h *CourseHandler) GetAll(c *gin.Context) {
	courses, err := h.Service.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get courses"})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func (h *CourseHandler) Create(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.Service.CreateCourse(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	c.JSON(http.StatusOK, course)
}
