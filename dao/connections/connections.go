package connections

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Mysql struct {
	User string
}

func NewMysqlConnection() *gorm.DB {
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"123456",
		"192.168.250.76:3306",
		"xm_test")

	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "fw_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,  // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
