package router

import (
	"github.com/linlinbupt123-crypto/account_service/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	account := r.Group("/account")
	{
		account.POST("/create", handler.CreateAccount)
		account.GET("/balance/:userId/:asset", handler.GetBalance)
	}

	transaction := r.Group("/transaction")
	{
		transaction.POST("/deposit", handler.Deposit)
		transaction.POST("/withdraw", handler.Withdraw)
	}

	trade := r.Group("/trade")
	{
		trade.POST("/freeze", handler.FreezeFunds)
	}

	return r
}
