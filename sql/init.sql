-- 账户余额表
CREATE TABLE account_balance (
    user_id BIGINT NOT NULL,
    currency VARCHAR(16) NOT NULL,
    available BIGINT NOT NULL DEFAULT 0, -- 可用余额
    frozen BIGINT NOT NULL DEFAULT 0,    -- 冻结余额
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, currency)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- 账变流水表，记录所有资金变动（充值、提现、冻结、解冻、成交、内部转账），保证幂等性
CREATE TABLE account_ledger (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_id VARCHAR(64) NOT NULL,             
    user_id BIGINT NOT NULL,
    currency VARCHAR(16) NOT NULL,
    account_change BIGINT NOT NULL,                
    frozen_change BIGINT NOT NULL DEFAULT 0,
    biz_type ENUM('deposit','withdraw','freeze','unfreeze','trade','transfer') NOT NULL, 
    biz_id VARCHAR(64) NOT NULL,            
    balance_after BIGINT NOT NULL,         
    frozen_after BIGINT NOT NULL,           
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_tx (tx_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



CREATE TABLE withdraws (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT NOT NULL,
  currency VARCHAR(16) NOT NULL,
  amount BIGINT NOT NULL,
  address VARCHAR(128) NOT NULL,

  status VARCHAR(32) NOT NULL,
  tx_hash VARCHAR(128),

  created_at DATETIME,
  updated_at DATETIME,

  INDEX idx_user_currency (user_id, currency),
  INDEX idx_status (status)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
