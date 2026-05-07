package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "url-shortener/api/models"
)

func (h *Handler) Shorten(c *gin.Context) {
    var req struct {
        URL string `json:"url" binding:"required,url"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    code, err := h.shortener.GenerateCode(req.URL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to shorten URL"})
        return
    }

    url := &models.URL{
        OriginalURL: req.URL,
        ShortCode:   code,
    }

    if err := h.pg.SaveURL(url); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save URL"})
        return
    }

    if err := h.redis.SetURL(code, req.URL); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to cache URL"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "short_url": "https://short.ly/" + code,
        "code":      code,
    })
}