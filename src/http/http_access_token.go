package http

import (
	"net/http"

	"github.com/Delaram-Gholampoor-Sagha/bookstore_oauth_api/src/domain/access_token"
	"github.com/Delaram-Gholampoor-Sagha/bookstore_oauth_api/src/utils/errors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

// this handler needs access token to work
func NewHandler(service access_token.Service) AccessTokenHandler {

	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {

	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)

}

// creating an access token based on on giving an access token
// so far we have created an access token base on the existance of the user in the database
func (handler *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	// this function takes the json request and attempts to use that request fill this access token based on this json configuration
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequessrError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := handler.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, at)

}
