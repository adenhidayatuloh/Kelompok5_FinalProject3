package entity

import "time"

type Task struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ID          uint   `gorm:"not null;primary_key;autoIncrement"`
	Title       string `gorm:"not null" `
	Description string `gorm:"not null"`
	CategoryID  uint
	Status      bool
	UserID      uint
	User        User
}
