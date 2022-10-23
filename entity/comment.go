package entity

import "time"

type Comment struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"userid" gorm:"type:integer;"`
	PhotoID   int       `json:"photoid" gorm:"type:integer;"`
	Message   string    `json:"message" gorm:"type:varchar; not null"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	User      User
	Photo     Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
