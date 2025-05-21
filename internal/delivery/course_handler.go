package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Курстарды алу сәтсіз болды"})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func (h *CourseHandler) Create(c *gin.Context) {
	roleVal, exists := c.Get("role")
	if !exists || roleVal != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Тек админдер курс қоса алады"})
		return
	}

	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Деректер жарамсыз"})
		return
	}

	if err := h.Service.CreateCourse(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Курс құру сәтсіз болды"})
		return
	}
	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) GetLessons(c *gin.Context) {
	idParam := c.Param("course_id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course_id"})
		return
	}

	lessons, err := h.Service.GetLessonsByCourseID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, lessons)
}

func (h *CourseHandler) GetCourseDetails(c *gin.Context) {
	idParam := c.Param("course_id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course_id"})
		return
	}

	course, err := h.Service.GetCourseWithLessons(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Курс табылмады"})
		return
	}
	c.JSON(http.StatusOK, course)
}

func (h *CourseHandler) Delete(c *gin.Context) {
	roleVal, exists := c.Get("role")
	if !exists || roleVal != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Тек админдер ғана курс өшіре алады"})
		return
	}

	idParam := c.Param("course_id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "course_id дұрыс емес"})
		return
	}

	if err := h.Service.DeleteCourse(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Курсты өшіру сәтсіз болды"})
		return
	}

	c.Status(http.StatusNoContent)
}
