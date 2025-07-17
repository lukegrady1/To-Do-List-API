package models

import (
	"time"
)

type Todo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(100);not null" json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
