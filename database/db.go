package database

import (
	"GoCrudApi/types"
	"fmt"
	"gorm.io/driver/mysql"
	"os"

	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	err = Db.AutoMigrate(&types.User{}, &types.Course{})
	if err != nil {
		panic("Database migration failed")
	}
}
