package handlers

import (
	"gin-api/db"
	"gin-api/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateStudentByID(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var student ent.Student

		id := c.Param("id")

		if err := c.ShouldBind(&student); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		}

		if err := db.UpdateStudentByID(c.Request.Context(),
			client,
			id,
			student.Name,
			student.School,
			student.Age); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": "updated",
		})
	}
}
