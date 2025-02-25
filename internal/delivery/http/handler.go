package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"url-shortener/internal/usecase"
	"url-shortener/pkg/log"
	"url-shortener/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase usecase.Usecase
	logger  log.Logger
}

func NewRouter(usecase usecase.Usecase, logger log.Logger) *gin.Engine {
	h := Handler{
		usecase: usecase,
		logger:  logger,
	}

	r := gin.Default()

	r.GET("/:alias", h.GetURL)

	r.POST("/url", h.SaveURL)

	return r
}

func (h Handler) GetURL(c *gin.Context) {
	alias := c.Param("alias")
	if alias == "" {
		h.logger.Error("invalid request.")
		c.JSON(http.StatusBadRequest, response.Error("invalid request", 400))
		return
	}

	url, err := h.usecase.GetURL(alias)
	if errors.Is(err, response.ErrURLNotFound) {
		h.logger.Error("url not found. (Alias - '", alias, "').")
		c.JSON(http.StatusNotFound, response.Error("url not found", 404))
		return
	}
	if err != nil {
		h.logger.Error("internal error: ", err)
		c.JSON(http.StatusInternalServerError, response.Error("internal error", 500))
		return
	}

	h.logger.Info("url '", url, "' received successfully (Alias - '", alias, "').")
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
		h.logger.Error("failed to decode request.")
		c.JSON(http.StatusBadRequest, response.Error("failed to decode request", 400))
		return
	}

	url := req.URL

	alias, err := h.usecase.SaveURL(url)
	if errors.Is(err, response.ErrURLExists) {
		h.logger.Error("url already exists: '", url, "'.")
		c.JSON(http.StatusBadRequest, response.Error("url already exists", 400))
		return
	}
	if err != nil {
		h.logger.Error("failed to add url. err - ", err)
		c.JSON(http.StatusInternalServerError, response.Error("failed to add url", 500))
		return
	}

	h.logger.Info("url '", url, "' saved. Alias - '", alias, "'.")
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"alias":  alias,
	})
}
