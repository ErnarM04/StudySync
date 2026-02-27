package models

import "time"

// DeadlineRequest is request-only input (no DB table); used to create or update deadline payloads.
type DeadlineRequest struct {
	TaskID  uint      `json:"task_id" binding:"required"`
	DueDate time.Time `json:"due_date" binding:"required"`
}
