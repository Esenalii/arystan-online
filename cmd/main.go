package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"rest-project/internal/models"
	"rest-project/internal/routes"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Course{})

	if err != nil {
		log.Fatal("Error on migrating to the DB", err)
	}

	r := gin.Default()
	r.Use(cors.Default())
	routes.SetupRoutes(r, db)
	r.Run(":8080")
}
