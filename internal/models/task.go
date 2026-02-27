package models

import "time"

// Task is a work item in "tasks": it may belong to a Subject and optionally a Sprint;
// deadline is the task-level target time; status drives workflow.
type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" example:"Finish Go backend"`
	Description string    `json:"description" example:"Implement CRUD with JWT"`
	Status      string    `json:"status" example:"in-progress"`
	Deadline    time.Time `json:"deadline" example:"2025-12-01T12:00:00Z"`
	SubjectID   uint      `json:"subject_id" example:"1"`
	Subject     Subject   `json:"subject" gorm:"foreignKey:SubjectID"`
	CreatedAt   time.Time `json:"created_at"`
	SprintID    uint      `json:"sprint_id" gorm:"default:null"`
}
