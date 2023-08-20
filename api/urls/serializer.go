package urls

import (
	"urlshort.ru/m/models"

	"urlshort.ru/m/schema"
)

// GetURLResponse returns a URLResponse struct based on the provided URL.
//
// It takes a parameter "url" of type models.URL and returns a URLResponse struct.
func GetURLResponse(url models.URL) URLResponse {
	return URLResponse{
		ID:          url.ID,
		OriginalURL: url.OriginalURL,
		ShortURL:    url.ShortURL,
		CreatedAt:   url.CreatedAt,
	}
}

// GetError404Response generates an error response with a 404 status code.
//
// Parameters:
// - err: the error that caused the response.
// Return:
// - ErrorResponse: the error response with the 404 status code.
func GetError404Response() schema.Response {
	return schema.Response{
		Code:    404,
		Message: "Not Found",
	}
}

// GetError400Response returns an ErrorResponse with code 400 and the error message.
//
// Parameter(s):
// - err: the error that will be used as the message for the ErrorResponse.
// Return type(s):
// - ErrorResponse: the error response object with code 400 and the error message.
func GetError400Response() schema.Response {
	return schema.Response{
		Code:    400,
		Message: "Bad Request",
	}
}
