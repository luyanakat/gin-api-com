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

func Login(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request TokenRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		user, err := db.GetUserByEmail(c.Request.Context(), client, request.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		credentialError := token.CheckPassword(request.Password, user)
		if credentialError != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid credentials",
			})
			return
		}

		tokenString, err := token.GenerateJWT(user.Email, user.UserName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	}
}
