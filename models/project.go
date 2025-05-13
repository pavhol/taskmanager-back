package models

import "time"

type Project struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"not null" json:"name"`
	Description string     `json:"description"`
	OwnerID     uint       `json:"owner_id"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Status      string     `gorm:"type:varchar(50);default:'active'" json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
