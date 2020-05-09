package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/models"
	httpRequest "test/models/http/request"
)

type TransactionHandler struct {
	transactionServ 		models.ITransactionService
}


func NewTransactionHandler(r *gin.Engine, u models.ITransactionService) {
	handler := &TransactionHandler{transactionServ: u}
	r.POST("/topup", handler.TopUp)
}

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