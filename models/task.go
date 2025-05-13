package models

import "time"

type Task struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	ProjectID   uint       `json:"project_id"`
	Title       string     `gorm:"not null" json:"title"`
	Description string     `json:"description"`
	Priority    string     `gorm:"type:varchar(20);default:'medium'" json:"priority"`
	Status      string     `gorm:"type:varchar(20);default:'todo'" json:"status"`
	ReporterID  uint       `json:"reporter_id"`
	AssigneeID  *uint      `json:"assignee_id"`
	StartDate   *time.Time `json:"start_date"`
	DueDate     *time.Time `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
