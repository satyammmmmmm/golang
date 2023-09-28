package http

import (
	"net/http"
	"oauth/domain/access_token"
	"oauth/utils/errors"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
	}
	if err := handler.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return

	}
	c.JSON(http.StatusCreated, at)
}
