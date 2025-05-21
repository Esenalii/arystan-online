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

	lessonRepo := repository.NewLessonRepository(db.DB)
	lessonService := service.NewLessonService(lessonRepo)
	lessonHandler := delivery.NewLessonHandler(r, lessonService)

	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.AdminOnly())
	{
		adminRoutes.POST("/courses", courseHandler.Create)
		adminRoutes.DELETE("/courses/:course_id", courseHandler.Delete)
		adminRoutes.GET("/lessons", lessonHandler.GetAll)
		adminRoutes.POST("/users", userHandler.CreateUser)
		adminRoutes.GET("/users", userHandler.GetAllUsers)
		adminRoutes.PUT("/users/:id", userHandler.UpdateUser)
		adminRoutes.DELETE("/users/:id", userHandler.DeleteUser)
		adminRoutes.POST("/courses/:course_id/lessons", lessonHandler.Create)
		adminRoutes.DELETE("/courses/:course_id/lessons/:id", lessonHandler.Delete)
		adminRoutes.PUT("/courses/:course_id/lessons/:id", lessonHandler.Update)
	}

	r.GET("/courses", courseHandler.GetAll)
	r.GET("/courses/:course_id", courseHandler.GetLessons)
	r.GET("/courses/:course_id/:lesson_id", lessonHandler.LessonOneId)
	r.GET("/courses/:course_id/details", courseHandler.GetCourseDetails)

	teacherRoutes := r.Group("/teacher")
	teacherRoutes.Use(middleware.TeacherOnly())
	{
		teacherRoutes.GET("/users", userHandler.GetAllUsers)
		teacherRoutes.POST("/courses/:course_id/lessons", lessonHandler.Create)
		teacherRoutes.DELETE("/courses/:course_id/lessons/:id", lessonHandler.Delete)
		teacherRoutes.PUT("/courses/:course_id/lessons/:lesson_id", lessonHandler.Update)
	}
}
