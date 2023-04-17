package handlers

import (
	"gin-api/db"
	"gin-api/ent"
	"gin-api/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterUserData struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

func RegisterUser(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user RegisterUserData

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPass, err := token.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		u, err := db.CreateUser(c.Request.Context(), client, user.Name, user.UserName, user.Email, string(hashedPass))
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
			"pass":     u.Password, // just test senstive struct tag
		})
	}
}
