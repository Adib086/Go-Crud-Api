package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID      uint     `gorm:"primaryKey"`
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []Course `gorm:"many2many:user_courses;" json:"courses"`
}
