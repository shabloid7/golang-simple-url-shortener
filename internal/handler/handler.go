package handler

import (
	"net/http"
	"github.com/shabloid7/golang-simple-url-shortener/internal/errors"
	"github.com/shabloid7/golang-simple-url-shortener/internal/model"
	"github.com/shabloid7/golang-simple-url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.URLService
}

func NewHandler(service service.URLService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/shorten", h.shorten)
	r.GET("/:code", h.redirect)
}

func (h *Handler) shorten(c *gin.Context) {
	var request model.ShortenRequest
	ctx := c.Request.Context()

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	
	originalURL := request.URL
	shortURL, err := h.service.Shorten(ctx, originalURL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not shorten url",
		})
		return
	}

	c.JSON(http.StatusCreated, model.ShortenResponse{
		ShortURL: shortURL,
	})
}

func (h *Handler) redirect(c *gin.Context) {
	code := c.Param("code")
	ctx := c.Request.Context()
	originalURL, err := h.service.Resolve(ctx, code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.ErrURLNotFound.Error(),
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, originalURL)
}