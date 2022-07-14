package models

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ChatDB *gorm.DB

func InitDB() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := viper.GetString(`mysql.dns`)
	fmt.Println("dsn:", dsn)
	ChatDB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
