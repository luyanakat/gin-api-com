package main

import (
	"gin-api/db"
	"gin-api/handlers"
	loggerconfig "gin-api/internal/log"
	"gin-api/internal/middleware"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// init log
	logger, err := loggerconfig.NewLog()
	if err != nil {
		log.Println(err)
	}
	defer logger.Sync()

	// connect db
	client, err := db.Connect()
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// run schema
	if err := db.Schema(client); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// gin init
	r := gin.Default()

	// use middleware
	r.Use(middleware.LogRequest(logger))

	// api route
	v1 := r.Group("/v1")
	{
		students := v1.Group("/students")
		{
			students.GET("", handlers.ListAllStudent(client))
			students.POST("", handlers.CreateStudent(client))
			students.GET("/:id", handlers.GetStudentByID(client))
			students.PATCH("/:id", handlers.UpdateStudentByID(client))
			students.DELETE("/:id", handlers.DeleteStudentByID(client))
		}
		users := v1.Group("/user")
		{
			users.POST("/register", handlers.RegisterUser(client))
			users.POST("/token", handlers.GenerateToken(client))
			secured := users.Group("/secured").Use(middleware.Auth())
			{
				secured.GET("/ping", handlers.Ping())
			}
		}
	}

	// run app (default 8080)
	r.Run()
}
