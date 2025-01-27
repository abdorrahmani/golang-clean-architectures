package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"vertical-slice/internal/modules/products/models"
)

func InitTestDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"test_user",
		"test_password",
		"127.0.0.1",
		"3307",
		"test_db",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database: " + err.Error())
	}
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		panic("failed to migrate test database: " + err.Error())
	}
	return db
}
