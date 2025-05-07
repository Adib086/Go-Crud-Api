package database

import (
	"GoCrudApi/types"
	"fmt"
	"gorm.io/driver/mysql"
	"log"

	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	dsn := "root:123@tcp(localhost:3306)/GoCurdApi?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		fmt.Println("Database connection established")
	}

	err = Db.AutoMigrate(&types.User{}, &types.Course{})
	if err != nil {
		panic("Database migration failed")
	}
}
