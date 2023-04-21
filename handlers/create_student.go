package handlers

import (
	"gin-api/db"
	"gin-api/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStudent(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var student ent.Student

		if err := c.ShouldBind(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		st, err := db.CreateStudent(c.Request.Context(), client, student.Name, student.Age, student.School)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": st.ID,
		})

	}
}
