package urls

import (
	"time"
)

type CreateURLBody struct {
	OriginalURL string `json:"original_url"`
}

type URLResponse struct {
	ID          uint      `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type ShortURLBody struct {
	OriginalURL string `json:"original_url"`
}
