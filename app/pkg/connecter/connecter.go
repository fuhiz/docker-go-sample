package connecter

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	// APP_MODEからデータベース名を決める
	dbName := os.Getenv("DB_NAME")
	if os.Getenv("APP_MODE") == "test" {
		dbName = os.Getenv("DB_NAME_TEST")
	}

	// DB接続 (https://gorm.io/docs/connecting_to_the_database.html)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		dbName,
		os.Getenv("DB_LOC"))
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = gormDB
}

func DB() *gorm.DB {
	return db
}
