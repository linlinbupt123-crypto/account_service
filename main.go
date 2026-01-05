package main

import (
	"fmt"
	"log"

	"github.com/linlinbupt123-crypto/account_service/config"
	"github.com/linlinbupt123-crypto/account_service/router"
)

func main() {
	// 加载配置
	err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := config.LoadConfig(); err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	r := router.SetupRouter()

	addr := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	log.Printf("server started on %s", addr)

	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
