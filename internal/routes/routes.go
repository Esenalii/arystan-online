package routes

import (
	"github.com/gin-gonic/gin"
	"rest-project/internal/auth"
	"rest-project/internal/db"
	"rest-project/internal/delivery"
	"rest-project/internal/middleware"
	"rest-project/internal/repository"
	"rest-project/internal/services"
)

func SetupRoutes(r *gin.Engine) {
	authRoutes := r.Group("api/v1/auth")
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
		authRoutes.GET("/me", middleware.AuthRequired(), auth.Me)
	}

	courseRepo := repository.NewCourseRepository(db.DB)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := delivery.NewCourseHandler(courseService)

	userRepo := repository.NewUserRepository(db.DB)
	userHandler := delivery.NewUserHandler(userRepo)

	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.AdminOnly())
	{
		adminRoutes.POST("/courses", courseHandler.Create)
		adminRoutes.POST("/users", userHandler.CreateUser)
		adminRoutes.GET("/users", userHandler.GetAllUsers)
		adminRoutes.PUT("/users/:id", userHandler.UpdateUser)
		adminRoutes.DELETE("/users/:id", userHandler.DeleteUser)
	}

	r.GET("/courses", courseHandler.GetAll)

	teacherRoutes := r.Group("/teacher")
	teacherRoutes.Use(middleware.TeacherOnly())
	{
		teacherRoutes.GET("/users", userHandler.GetAllUsers)
	}
}
