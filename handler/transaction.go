package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taufiqade/gowallet/models"
	httpRequest "github.com/taufiqade/gowallet/models/http/request"
)

// TransactionHandler godoc
type TransactionHandler struct {
	transactionServ models.ITransactionService
}

// NewTransactionHandler godoc
func NewTransactionHandler(r *gin.Engine, u models.ITransactionService) {
	handler := &TransactionHandler{transactionServ: u}
	r.POST("/topup", handler.TopUp)
}

// TopUp godoc
func (t *TransactionHandler) TopUp(c *gin.Context) {
	payload := httpRequest.TopUpRequest{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
	}

	err := t.transactionServ.TopUp(payload.BeneficiaryId, &payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusCreated, err)
}
