package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"url-shortener/api/models"
)

func (h *Handler) Redirect(c *gin.Context) {
	code := c.Param("code")
	ctx := c.Request.Context()

	longURL, err := h.redis.GetLongURL(ctx, code)
	if err != nil {
		log.Printf("redis lookup failed for code=%s: %v", code, err)
	}

	if longURL == "" {
		urlRecord, err := h.pg.GetURL(ctx, code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to lookup URL"})
			return
		}
		if urlRecord == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}

		longURL = urlRecord.LongURL
		if err := h.redis.SetLongURL(ctx, code, longURL); err != nil {
			log.Printf("redis cache set failed for code=%s: %v", code, err)
		}
	}

	click := &models.Click{
		Code:      code,
		Timestamp: time.Now().UTC(),
		IP:        c.ClientIP(),
		UserAgent: c.GetHeader("User-Agent"),
		Referrer:  c.GetHeader("Referer"),
	}

	go func() {
		publishCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if err := h.sqs.PublishClick(publishCtx, click); err != nil {
			log.Printf("publish click failed for code=%s: %v", code, err)
		}
	}()

	c.Redirect(http.StatusFound, longURL)
}