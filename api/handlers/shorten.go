package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
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

	code, err := h.createUniqueCode(c, req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to shorten URL"})
		return
	}

	if err := h.redis.SetLongURL(c.Request.Context(), code, req.URL); err != nil {
		log.Printf("redis cache set failed for code=%s: %v", code, err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"short_url": h.baseShortURL + "/" + code,
		"code":      code,
	})
}

func (h *Handler) createUniqueCode(c *gin.Context, longURL string) (string, error) {
	for attempts := 0; attempts < 5; attempts++ {
		code, err := h.shortener.Generate()
		if err != nil {
			return "", err
		}

		err = h.pg.InsertURL(c.Request.Context(), &models.URL{
			Code:    code,
			LongURL: longURL,
		})
		if err == nil {
			return code, nil
		}

		var pgErr *pgconn.PgError
		if !errors.As(err, &pgErr) || pgErr.Code != "23505" {
			return "", err
		}
	}

	return "", errors.New("short code collision retries exhausted")
}