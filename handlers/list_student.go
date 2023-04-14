package handlers

import (
	"gin-api/db"
	"gin-api/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAllStudent(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		students, err := db.GetAllStudent(c.Request.Context(), client)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, students)
	}
}
