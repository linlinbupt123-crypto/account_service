package handler

import (
	"net/http"
	"strconv"

	"github.com/linlinbupt123-crypto/account_service/service"

	"github.com/gin-gonic/gin"
)

type FreezeFundsRequest struct {
	Currency string `json:"currency" binding:"required"`
	Amount   string `json:"amount" binding:"required"`
	TxId     string `json:"tx_id" binding:"required"`
	BizId    string `json:"biz_id" binding:"required"`
}

func FreezeFunds(c *gin.Context) {
	userIDStr := c.Param("userID")
	var req FreezeFundsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := strconv.ParseInt(userIDStr, 10, 64)
	amount, _ := strconv.ParseInt(req.Amount, 10, 64)
	accountService := service.AccountService{}
	err := accountService.FreezeFunds(userID, req.Currency, amount, req.TxId, req.BizId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "freeze success"})
}
