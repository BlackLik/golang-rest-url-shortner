package models

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	gorm.Model
	OriginalURL string    `gorm:"uniqueIndex"`
	ShortURL    string    `gorm:"uniqueIndex"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
}
