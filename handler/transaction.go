package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/taufiqade/gowallet/models"
	httpRequest "github.com/taufiqade/gowallet/models/http/request"
	authMiddleware "github.com/taufiqade/gowallet/utils/middleware"
)

// TransactionHandler godoc
type TransactionHandler struct {
	transactionServ models.ITransactionService
}

// NewTransactionHandler godoc
func NewTransactionHandler(r *gin.Engine, u models.ITransactionService) {
	handler := &TransactionHandler{transactionServ: u}
	transactionGroup := r.Group("transaction")
	midleware := authMiddleware.DefaultMiddleware{}

	transactionGroup.Use(midleware.JWTAuthMidlewareAdmin())
	{
		transactionGroup.POST("/topup", handler.TopUp)
		transactionGroup.POST("/transfer", handler.Transfer)
	}
}

// TopUp godoc
func (t *TransactionHandler) TopUp(c *gin.Context) {
	payload := httpRequest.TransactionRequest{}
	payload.IP = c.ClientIP()
	userAgent := c.Request.Header.Get("User-Agent")
	agent, _ := json.Marshal(userAgent)
	payload.UserAgent = string(agent)
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

	err := t.transactionServ.TopUp(payload.Email, &payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusCreated, "TopUp successfully executed")
}

// Transfer godoc
func (t *TransactionHandler) Transfer(c *gin.Context) {
	payload := httpRequest.TransactionRequest{}
	payload.IP = c.ClientIP()
	userAgent := c.Request.Header.Get("User-Agent")
	agent, _ := json.Marshal(userAgent)
	payload.UserAgent = string(agent)
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

	userID, _ := strconv.Atoi(c.Request.Header["Wallet-Uid"][0])
	if err := t.transactionServ.Transfer(userID, payload.Email, &payload); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusCreated, "transfer successfully executed")
}
