package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-project/internal/models"
	"rest-project/internal/repository"
	"strconv"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.Repo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Қолданушыларды алу сәтсіз аяқталды"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Қате ID"})
		return
	}

	user, err := h.Repo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Қолданушы табылмады"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userCreate models.UserEdit

	if err := c.ShouldBindJSON(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Қате сұраныс денесі"})
		return
	}

	newUser, err := h.Repo.Create(userCreate.Name, userCreate.Email, userCreate.Password, userCreate.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Қолданушы құру сәтсіз"})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Қате ID"})
		return
	}

	var userEdit models.UserEdit
	if err := c.ShouldBindJSON(&userEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Қате сұраныс денесі"})
		return
	}

	updatedUser, err := h.Repo.Update(id, &userEdit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Қолданушы табылмады"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Қате ID"})
		return
	}

	if err := h.Repo.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Қолданушы табылмады"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Қолданушы сәтті өшірілді"})
}
