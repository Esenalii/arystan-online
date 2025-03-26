package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-project/internal/models"
	"rest-project/internal/services"
	"strconv"
)

type StudentHandler struct {
	service *services.StudentService
}

// Получение списка всех студентов
func (h *StudentHandler) GetAllStudents(c *gin.Context) {
	students := h.service.GetAllStudents()
	c.JSON(http.StatusOK, students)
}

// Получение студента по ID
func (h *StudentHandler) GetStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	student, err := h.service.GetStudentById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// Создание нового студента
func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var studentCreate models.StudentEdit

	if err := c.ShouldBindJSON(&studentCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newStudent := h.service.CreateStudent(studentCreate)
	c.JSON(http.StatusCreated, newStudent)
}

// Обновление данных студента
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var studentEdit models.StudentEdit
	if err := c.ShouldBindJSON(&studentEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedStudent, err := h.service.UpdateStudent(id, studentEdit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, updatedStudent)
}

// Удаление студента
func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	if err := h.service.DeleteStudent(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
