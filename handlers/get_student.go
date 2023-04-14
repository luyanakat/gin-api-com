package handlers

import (
	"gin-api/db"
	"gin-api/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudentByID(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		student, err := db.GetStudentByID(c.Request.Context(), client, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, student)
	}
}
