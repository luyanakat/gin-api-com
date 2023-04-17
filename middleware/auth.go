package middleware

import (
	"gin-api/token"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{
				"error": "request doesn't contain access token",
			})
			return
		}

		if err := token.ValidateToken(tokenString); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Next()
	}
}
