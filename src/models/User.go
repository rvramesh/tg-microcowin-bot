package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;<-:create"`
	CreatedAt    time.Time `gorm:"<-:create"`
	UpdatedAt    time.Time
	FirstName    string `gorm:"<-:create"`
	LastName     string `gorm:"<-:create"`
	ChatId       int64  `gorm:"<-:create"`
	UserName     string `gorm:"uniqueindex;<-:create"`
	CurrentState string
	Data         string
	IsBlocked    bool
	IsAdmin      bool
}
