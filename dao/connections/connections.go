package connections

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	MysqlUserName           = "root"
	MysqlPassWordHome       = "123456"
	MysqlPassWordCompany    = "1234"
	MysqlHostAndPortHome    = "192.168.250.76:3306"
	MysqlHostAndPortCompany = "localhost:3306"
	MysqlDatabaseHome       = "xm_test"
	MysqlDatabaseCompany    = "test"
)

type Mysql struct {
	User string
}

func NewMysqlConnection() *gorm.DB {
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MysqlUserName,
		MysqlPassWordCompany,
		MysqlHostAndPortCompany,
		MysqlDatabaseCompany)

	db, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "fw_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,  // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		CreateBatchSize: 100, // 自己本地docker玩玩,就别弄太多数据了。
		Logger:          logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

