package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"url-shortener/pkg/response"

	"github.com/gin-gonic/gin"
)

type usecase interface {
	GetURL(alias string) (string, error)
	SaveURL(url string) (string, error)
}

type Handler struct {
	usecase usecase
}

func NewRouter(usecase usecase) *gin.Engine {
	h := Handler{
		usecase,
	}

	r := gin.Default()

	r.GET("/:alias", h.GetURL)

	r.POST("/url", h.SaveURL)

	return r
}

func (h Handler) GetURL(c *gin.Context) {
	alias := c.Param("alias")
	if alias == "" {
		c.JSON(http.StatusBadRequest, response.Error("invalid request", 400))
		return
	}

	url, err := h.usecase.GetURL(alias)
	if errors.Is(err, response.ErrURLNotFound) {
		c.JSON(http.StatusNotFound, response.Error("url not found", 404))
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error("internal error", 500))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"url":    url,
	})
}

type request struct {
	URL string `json:"url" validate:"required,url"`
}

func (h Handler) SaveURL(c *gin.Context) {
	var req request

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("failed to decode request", 400))
		return
	}

	url := req.URL

	alias, err := h.usecase.SaveURL(url)
	if errors.Is(err, response.ErrURLExists) {
		c.JSON(http.StatusBadRequest, response.Error("url already exists", 400))
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error("failed to add url", 500))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"alias":  alias,
	})
}
