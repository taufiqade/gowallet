package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/taufiqade/gowallet/models"
	httpRequest "github.com/taufiqade/gowallet/models/http/request"
)

// UserHandler godoc
type UserHandler struct {
	userServ models.IUserService
}

// NewUserHandler godoc
func NewUserHandler(r *gin.Engine, u models.IUserService) {
	handler := &UserHandler{userServ: u}
	r.GET("/user/:id", handler.GetUserByID)
	r.POST("/user", handler.Create)
}

// GetUserByID godoc
func (u *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

	res, err := u.userServ.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Create godoc
func (u *UserHandler) Create(c *gin.Context) {
	payload := httpRequest.UserRequest{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.userServ.Create(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusCreated, user)
}
