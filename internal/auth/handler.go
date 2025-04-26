package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rest-project/internal/db"
	user "rest-project/internal/models"
)

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid"})
		return
	}

	var u user.User
	db.DB.Where("email = ?", req.Email).First(&u)
	if u.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "login or password incorrect"})
		return
	}

	token, _ := GenerateJWT(u.ID)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"role":  u.Role,
	})
}

func Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Name == "" || req.Password == "" || req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var existing user.User
	if err := db.DB.Where("email = ?", req.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already taken"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	u := user.User{
		Name:     req.Name,
		Password: string(hashedPassword),
		Email:    req.Email,
		Role:     req.Role,
	}
	if err := db.DB.Create(&u).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Me(c *gin.Context) {
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, ok := userIDRaw.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var u user.User
	if err := db.DB.First(&u, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       u.ID,
		"username": u.Name,
		"email":    u.Email,
		"role":     u.Role,
	})
}
