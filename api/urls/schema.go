package urls

import "time"

type CreateURLBody struct {
	OriginalURL string `json:"original_url"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type URLResponse struct {
	ID          uint      `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
