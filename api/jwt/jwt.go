package jwt

import (
	"github.com/gofiber/fiber/v2"
	"urlshort.ru/m/models"
)

// Register registers the API routes for JWT authentication.
//
// api - The fiber.Router to register the routes with.
//
// No return value.
func Register(api fiber.Router) {
	apiJWT := api.Group("/jwt")
	localDb = models.DATABASE

	apiJWT.Get("/refresh", refreshHandler)
	apiJWT.Get("/logout", logoutHandler)
	apiJWT.Post("/register", registerHandler)
	apiJWT.Post("/login", loginHandler)
	apiJWT.Post("/check", checkHandler)
	apiJWT.Delete("/delete", deleteHandler)
}
