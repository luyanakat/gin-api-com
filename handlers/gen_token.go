package handlers

import (
	"gin-api/db"
	"gin-api/ent"
	"gin-api/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request TokenRequest
		var user ent.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err := db.GetUserByEmail(c.Request.Context(), client, user.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		credentialError := token.CheckPassword(request.Password, &user)
		if credentialError != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid credentials",
			})
		}

		tokenString, err := token.GenerateJWT(user.Email, user.UserName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	}
}
