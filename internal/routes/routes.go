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

	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.AdminOnly())
	{
		adminRoutes.POST("/courses", courseHandler.Create)
	}

	r.GET("/courses", courseHandler.GetAll)

	studentRepo := repository.NewStudentRepository(db.DB)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := delivery.NewStudentHandler(studentService)

	adminRoutes.POST("/students", studentHandler.CreateStudent)
	adminRoutes.PUT("/students/:id", studentHandler.UpdateStudent)
	adminRoutes.DELETE("/students/:id", studentHandler.DeleteStudent)

	teacherRoutes := r.Group("/teacher")
	teacherRoutes.Use(middleware.TeacherOnly())
	{
		teacherRoutes.GET("/students", studentHandler.GetAllStudents)
		teacherRoutes.GET("/students/:id", studentHandler.GetStudent)
	}
}
