package core

import (
	"fmt"
	"mark-blog-backend/global"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 连接数据库操作
func InitGorm() *gorm.DB {
	MysqlInfo := global.Config.Mysql
	if MysqlInfo.Host == "" {
		global.Log.Warning("未配置mysql，取消gorm的连接")
	}
	var MysqlLogger logger.Interface
	//获取数据库信息
	dsn := MysqlInfo.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: MysqlLogger,
	})

	if err != nil {
		global.Log.Fatal(fmt.Printf("[%s] mysql连接失败", dsn))
	}

	sqlDB, _ := db.DB()

	//判断连接是否成功
	if err := sqlDB.Ping(); err != nil {
		fmt.Println("open database fail")
	}
	fmt.Println("open database success")
	return db
}
