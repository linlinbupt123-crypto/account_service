package handler

import (
	"net/http"
	"strconv"

	"github.com/linlinbupt123-crypto/account_service/service"

	"github.com/gin-gonic/gin"
)

var accountService = &service.AccountService{}

func CreateAccount(c *gin.Context) {
	userIDStr := c.PostForm("user_id")
	currency := c.PostForm("currency")
	userID, _ := strconv.ParseInt(userIDStr, 10, 64)

	err := accountService.CreateAccount(userID, currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "account created"})
}

func GetBalance(c *gin.Context) {
	userIDStr := c.Param("userId")
	currency := c.Param("asset")
	userID, _ := strconv.ParseInt(userIDStr, 10, 64)

	acc, err := accountService.GetBalance(userID, currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, acc)
}
