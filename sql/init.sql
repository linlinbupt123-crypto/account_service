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
    tx_id VARCHAR(64) NOT NULL,             -- 幂等交易ID
    user_id BIGINT NOT NULL,
    currency VARCHAR(16) NOT NULL,
    change BIGINT NOT NULL,                 -- 可用余额变动
    frozen_change BIGINT NOT NULL DEFAULT 0,-- 冻结余额变动
    biz_type ENUM('deposit','withdraw','freeze','unfreeze','trade','transfer') NOT NULL, 
    biz_id VARCHAR(64) NOT NULL,            -- 业务ID，如订单号
    balance_after BIGINT NOT NULL,          -- 可用余额变动后的余额
    frozen_after BIGINT NOT NULL,           -- 冻结余额变动后的余额
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_tx (tx_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



