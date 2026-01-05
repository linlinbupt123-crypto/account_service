account_service/ 

├── main.go # 程序入口 

├── router/ 

    │ └── router.go # Gin 路由注册 

├── handler/ 

    │ └── account.go # 用户账户接口 

    │ └── transaction.go # 资金操作接口 

    │ └── trade.go # 撮合调用接口 

├── service/ 

    │ └── account_service.go # 账户业务逻辑 

    │ └── transaction_service.go # 流水业务逻辑 

├── model/ 

    │ └── account.go 

    │ └── transaction.go 

├── repository/ 

    │ └── account_repo.go 

    │ └── transaction_repo.go 

├── config/ 

    │ └── config.yml # 数据库配置等 

├── utils/ 

    │ └── id_generator.go # txId 或 orderId 生成 

└── go.mod