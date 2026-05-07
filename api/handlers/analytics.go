package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Analytics returns aggregate click counts for a short code.
func (h *Handler) Analytics(c *gin.Context) {
	code := c.Param("code")

	count, err := h.pg.GetClickCount(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch analytics"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":         code,
		"total_clicks": count,
	})
}
