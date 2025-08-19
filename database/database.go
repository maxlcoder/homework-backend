package database

import (
	"fmt"
	"time"

	"github.com/maxlcoder/homework-backend/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	dsn := viper.Get("database.mysql.dsn").(string)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("数据库连接失败：%w", err)
	}

	DB = db

	sqlDB, _ := db.DB()
	if err != nil {
		return fmt.Errorf("获取数据库实例失败：%w", err)
	}

	// 配置连接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = model.AutoMigrate(db)
	if err != nil {
		return fmt.Errorf("数据库迁移失败：%w", err)
	}

	return nil
}
