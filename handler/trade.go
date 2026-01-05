package handler

import (
	"net/http"
	"strconv"

	"github.com/linlinbupt123-crypto/account_service/service"

	"github.com/gin-gonic/gin"
)

func FreezeFunds(c *gin.Context) {
	userIDStr := c.PostForm("user_id")
	currency := c.PostForm("currency")
	amountStr := c.PostForm("amount")
	txID := c.PostForm("tx_id")
	bizID := c.PostForm("biz_id")

	userID, _ := strconv.ParseInt(userIDStr, 10, 64)
	amount, _ := strconv.ParseInt(amountStr, 10, 64)
	accountService := service.AccountService{}
	err := accountService.FreezeFunds(userID, currency, amount, txID, bizID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "freeze success"})
}
