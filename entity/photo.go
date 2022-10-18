package entity

import "time"

type Photo struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `gorm:"not null"`
	UserID    int       `json:"userid"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	User      User
}
