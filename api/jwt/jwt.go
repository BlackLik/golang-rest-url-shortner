package jwt

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"urlshort.ru/m/models"
)

var localDb *gorm.DB

// Register registers the API routes for JWT authentication.
//
// api - The fiber.Router to register the routes with.
//
// No return value.
func Register(api fiber.Router) {
	apiJWT := api.Group("/jwt")
	localDb = models.DATABASE
	_ = apiJWT
	_ = localDb
	// TODO refresh handlers apiJWT.Post("/refresh", refresh)
	// TODO logout handlers apiJWT.Post("/logout", logout)
	// TODO register handlers apiJWT.Post("/register", register)
	// TODO login handlers apiJWT.Post("/login", login)
	// TODO check handlers apiJWT.Post("/check", check)
}
