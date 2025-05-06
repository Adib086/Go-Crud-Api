package types

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey"`
	Title string `json:"title"`
}
