package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-project/internal/repository"
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
