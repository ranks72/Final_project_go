package entity

import "time"

type SocialMedia struct {
	ID             int       `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" gorm:"type:varchar; unique;not null"`
	SocialMediaUrl string    `json:"social_media_url" gorm:"unique;not null"`
	UserID         int       `json:"userid"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
