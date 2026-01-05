package repository

import (
	"fmt"

	"github.com/linlinbupt123-crypto/account_service/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(user, password, host string, port int, dbname string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	// 自动迁移表
	err = DB.AutoMigrate(&model.Account{}, &model.Transaction{})
	return err
}
