package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	ROLE_USER  = "user"
	ROLE_ADMIN = "admin"
)

type URL struct {
	gorm.Model
	OriginalURL string    `gorm:"uniqueIndex"`
	ShortURL    string    `gorm:"uniqueIndex"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
}

type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex; not null"`
	Password     string `gorm:"not null"`
	RefreshToken string `gorm:"index"`
	Role         string `gorm:"default:'user'"`
}
