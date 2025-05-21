package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"rest-project/internal/models"
	"rest-project/internal/services"
)

type LessonHandler struct {
	Service *service.LessonService
}

func NewLessonHandler(r *gin.Engine, svc *service.LessonService) *LessonHandler {
	return &LessonHandler{Service: svc}
}

func (h *LessonHandler) RegisterRoutes(r *gin.Engine) {
	// Public lessons
	r.GET("/lessons", h.GetAll)
	r.GET("/courses/:course_id/lessons", h.GetByCourse)

	// Protected routes
	grp := r.Group("/lessons")
	grp.Use()
	grp.POST("/courses/:course_id/lessons", h.Create)
	grp.PUT("/lessons/:id", h.Update)
	grp.DELETE("/lessons/:id", h.Delete)
}

func (h *LessonHandler) GetAll(c *gin.Context) {
	list, err := h.Service.GetAllLessons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch lessons"})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *LessonHandler) GetByCourse(c *gin.Context) {
	courseID, err := strconv.ParseUint(c.Param("course_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}
	list, err := h.Service.GetLessonsByCourseID(uint(courseID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch lessons for course"})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *LessonHandler) Create(c *gin.Context) {
	courseID, _ := strconv.ParseUint(c.Param("course_id"), 10, 64)
	var lesson models.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	lesson.CourseID = uint(courseID)
	if err := h.Service.CreateLesson(&lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create lesson"})
		return
	}
	c.JSON(http.StatusCreated, lesson)
}

func (h *LessonHandler) Update(c *gin.Context) {
	lessonID, err := strconv.ParseUint(c.Param("lesson_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID"})
		return
	}

	existing, err := h.Service.GetLessonByID(uint(lessonID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		return
	}

	var input models.Lesson
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	existing.Title = input.Title
	existing.Content = input.Content
	existing.Video = input.Video
	existing.CourseID = input.CourseID

	if err := h.Service.UpdateLesson(existing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update lesson"})
		return
	}

	c.JSON(http.StatusOK, existing)
}

func (h *LessonHandler) Delete(c *gin.Context) {
	lessonID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID"})
		return
	}

	courseID := c.Param("course_id")
	if courseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Course ID is required"})
		return
	}

	if err := h.Service.DeleteLesson(uint(lessonID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete lesson"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *LessonHandler) LessonOneId(c *gin.Context) {
	lessonIDParam := c.Param("lesson_id")
	lessonID, err := strconv.ParseUint(lessonIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID"})
		return
	}

	lesson, err := h.Service.GetLessonByID(uint(lessonID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lesson not found"})
		return
	}

	c.JSON(http.StatusOK, lesson)
}
