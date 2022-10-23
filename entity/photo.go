package entity

import "time"

type Photo struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:" unique;not null"`
	Caption   string    `json:"caption" gorm:"type:varchar(255);"`
	PhotoUrl  string    `json:"photourl" gorm:"type:varchar(255); unique;not null"`
	UserID    int       `json:"userid"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
