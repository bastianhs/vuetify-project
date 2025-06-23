package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	FullName string `json:"full_name"`
	Password string `json:"-" gorm:"column:password"`
}

func MigrateUsers(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
