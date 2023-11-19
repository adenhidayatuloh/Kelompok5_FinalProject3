package entity

import "time"

type Category struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uint   `gorm:"not null;primary_key;autoIncrement"`
	Type      string `gorm:"not null"`
	Task      []Task
}
