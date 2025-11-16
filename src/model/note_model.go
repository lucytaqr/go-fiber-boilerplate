package model

import "time"

type Note struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"userId" gorm:"type:uuid;index"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ImageURL  *string   `json:"imageUrl,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
