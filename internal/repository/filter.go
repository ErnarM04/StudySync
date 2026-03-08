package repository

import "time"

// TaskFilter describes list semantics for tasks: page and limit drive pagination;
// the other fields narrow the result set (status, subject, text search, sort, deadline window).
type TaskFilter struct {
	Page           int
	Limit          int
	Status         string
	SubjectID      *uint
	Search         string
	Sort           string // e.g. "created_at desc"
	DeadlineBefore *time.Time
	DeadlineAfter  *time.Time
}
