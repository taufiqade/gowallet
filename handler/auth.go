package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taufiqade/gowallet/models"
	httpRequest "github.com/taufiqade/gowallet/models/http/request"
)

// AuthHandler godoc
type AuthHandler struct {
	authServ models.IAuthService
}

// NewAuthHandler godoc
func NewAuthHandler(r *gin.Engine, a models.IAuthService) {
	handler := &AuthHandler{authServ: a}
	r.POST("/login", handler.Login)
}

// Login godoc
func (a *AuthHandler) Login(c *gin.Context) {
	payload := httpRequest.Login{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := a.authServ.CreateToken(payload.Email, payload.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
