package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "url-shortener/api/models"
)

func (h *Handler) Redirect(c *gin.Context) {
    code := c.Param("code")

    // Hot path: Redis lookup
    originalURL, err := h.redis.GetURL(code)
    if err != nil {
        // Cache miss: fall back to PostgreSQL
        originalURL, err = h.pg.GetURL(code)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
            return
        }
        // Repopulate cache
        h.redis.SetURL(code, originalURL)
    }

    // Async: publish click event to SQS (non-blocking)
    go h.sqs.PublishClickEvent(&models.ClickEvent{
        ShortCode:  code,
        UserAgent:  c.GetHeader("User-Agent"),
        IPAddress:  c.ClientIP(),
        Referrer:   c.GetHeader("Referer"),
    })

    c.Redirect(http.StatusMovedPermanently, originalURL)
}