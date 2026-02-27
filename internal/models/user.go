package models

import "time"

// User is the persisted account row in table "users": identity, login email (unique),
// bcrypt hash (never exposed as JSON), role for authorization, and creation time.
type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Email        string    `json:"email" gorm:"uniqueIndex"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}
