package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	FullName string `json:"full_name"`
	Password string `json:"password" gorm:"column:password"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

func MigrateUsers(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

func ToUserResponse(u User) UserResponse {
	return UserResponse{
		ID:       u.ID,
		Username: u.Username,
		FullName: u.FullName,
	}
}
