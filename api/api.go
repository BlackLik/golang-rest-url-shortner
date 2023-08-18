package api

import (
	"github.com/gofiber/fiber/v2"
	"urlshort.ru/m/api/jwt"
	"urlshort.ru/m/api/urls"
)

// Register registers the API routes for the given Fiber app and Gorm DB.
//
// app: The Fiber app instance.
// db: The Gorm DB instance.
func Register(app *fiber.App) {
	api := app.Group("/api")
	urls.Register(api)
	jwt.Register(api)
}
