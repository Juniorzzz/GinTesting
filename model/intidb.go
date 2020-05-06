package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

// DB 数据库链接单例
var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("mysql", "root:root@/db_gin?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	DB = db
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetConnMaxLifetime(time.Second * 20)

	migration()
}

//执行数据迁移
func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Video{})
}
