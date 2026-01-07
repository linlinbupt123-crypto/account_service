package utils

const (
	StatusInit      = "INIT"       //  刚创建，还没冻结钱
	StatusFrozen    = "FROZEN"     //    钱已冻结
	StatusChainSent = "CHAIN_SENT" //已发送链上交易
	StatusConfirmed = "CONFIRMED"  //链上确认成功
	StatusFailed    = "CONFIRMED"  //中途失败（可回滚）
)
