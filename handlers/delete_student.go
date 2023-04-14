package handlers

import (
	"gin-api/db"
	"gin-api/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteStudentByID(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := db.DeleteStudentByID(c.Request.Context(), client, id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"data": "Student deleted",
		})
	}
}
