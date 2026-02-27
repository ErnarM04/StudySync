package models

import "time"

// Sprint is a time-boxed iteration in "sprints"; Tasks link back with sprint_id.
type Sprint struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" example:"Sprint 1"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status" example:"active"` // planned | active | completed
	Tasks     []Task    `json:"tasks" gorm:"foreignKey:SprintID"`
	CreatedAt time.Time `json:"created_at"`
}
