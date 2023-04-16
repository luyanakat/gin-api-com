package handlers

import (
	"gin-api/db"
	"gin-api/ent"
	"gin-api/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user ent.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := token.HashPassword(user.Password, &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"data": err.Error(),
			})
			return
		}

		u, err := db.CreateUser(c.Request.Context(), client, user.Name, user.UserName, user.Email, user.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"userId":   u.ID,
			"email":    u.Email,
			"username": u.UserName,
			"pass":     u.Password,
		})
	}
}
