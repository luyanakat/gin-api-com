package main

import (
	"context"
	"fmt"
	"gin-api/db"
	"gin-api/handlers"
	loggeri22 "gin-api/internal/log"
	"gin-api/internal/middleware"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// init log
	logger, err := loggeri22.NewLog()
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

	student, err := db.GetStudentByID(context.Background(), client, "e4d453bc-e04f-4161-b404-f794ff813")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(student)

	// gin init
	r := gin.Default()

	// use middleware
	r.Use(middleware.LogRequest(logger))
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
	}

	r.Run()
}
