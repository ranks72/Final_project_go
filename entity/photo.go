package entity

import "time"

type Photo struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Caption   string
	PhotoUrl  string `gorm:"not null"`
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
}

type User_photo struct {
	ID       int `gorm:"primaryKey"`
	Email    string
	Username string
}
