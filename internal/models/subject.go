package models

import "time"

// Subject is a study area stored in "subjects"; tasks reference it via subject_id.
type Subject struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
