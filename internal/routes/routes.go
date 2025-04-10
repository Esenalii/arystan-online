package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rest-project/internal/delivery"
	"rest-project/internal/handler"
	"rest-project/internal/repository"
	"rest-project/internal/services"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// 📌 Auth роуттары
	r.POST("/register", func(c *gin.Context) {
		handler.Register(c, db)
	})

	r.POST("/login", func(c *gin.Context) {
		handler.Login(c, db)
	})

	// 📌 Student бөлімінің Route-тары
	studentRepo := repository.NewStudentRepository(db)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := delivery.NewStudentHandler(studentService)

	students := r.Group("api/v1/students")
	{
		students.GET("/", studentHandler.GetAllStudents)
		students.POST("/", studentHandler.CreateStudent)
		students.GET("/:id", studentHandler.GetStudent)
		students.PUT("/:id", studentHandler.UpdateStudent)
		students.DELETE("/:id", studentHandler.DeleteStudent)
	}
	courseRepo := repository.NewCourseRepository(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := delivery.NewCourseHandler(courseService)

	courses := r.Group("/api/v1/courses")
	{
		courses.GET("/", courseHandler.GetAll)
		courses.POST("/", courseHandler.Create)
	}
	userRepo := repository.NewUserRepository(db)
	userHandler := delivery.NewUserHandler(userRepo)

	users := r.Group("/api/v1/users")
	{
		users.GET("/", userHandler.GetAllUsers)
	}

}
