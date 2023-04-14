package handlers

import (
	"gin-api/db"
	"gin-api/ent"
	"gin-api/internal/paging"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAllStudent(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var page paging.Paging

		if err := c.ShouldBind(&page); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		page.Process()

		students, err := db.GetAllStudent(c.Request.Context(), client, &page)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  students,
			"page":  page.Page,
			"limit": page.Limit,
			"total": page.Total,
		})
	}
}
