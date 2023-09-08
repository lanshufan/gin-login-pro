package modules

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		fmt.Println("open app.ini failure", err)
		return
	}

	// rw
	user := cfg.Section("mysql").Key("username").String()
	pass := cfg.Section("mysql").Key("password").String()
	ip := cfg.Section("mysql").Key("ip").String()
	port := cfg.Section("mysql").Key("port").String()
	database := cfg.Section("mysql").Key("database").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, ip, port, database)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("conn to mysql failure", err)
		return
	}

	// 连接池
	sqlDB, _ := DB.DB()
	sqlDB.SetConnMaxIdleTime(time.Hour)
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetConnMaxLifetime(24 * time.Hour)
	sqlDB.SetMaxOpenConns(20)
}
