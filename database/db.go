package database

import (
	"GoCrudApi/types"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	dsn := "sqlserver://@127.0.0.1:5752?database=GoCrudApi&encrypt=disable"
	var err error
	Db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection established")

	err = Db.AutoMigrate(&types.User{}, &types.Course{})
	if err != nil {
		panic("Database migration failed")
	}
}
