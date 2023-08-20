package urls

import (
	"github.com/gofiber/fiber/v2"
	"urlshort.ru/m/models"
)

// Register registers the given API with the provided fiber.Router.
//
// api: The fiber.Router instance to register.
//
// This function registers several URL endpoints for handling short URLs.
// It creates a group of endpoints under "/urls" and adds GET and POST
// handlers for retrieving and creating short URLs, respectively.
//
// TODO: Uncomment the lines for DELETE and PATCH handlers once implemented.
//
// Return type: None.
func Register(api fiber.Router) {
	apiUrls := api.Group("/urls")
	localDb = models.DATABASE
	apiUrls.Get("/:shorturl", getURLWithShort)
	apiUrls.Delete("/:shorturl", deleteURLWithShort)
	// TODO api.Patch("/:shorturl", updateURLWithShort)
	apiUrls.Patch("/:shorturl", updateURLWithShort)
	apiUrls.Post("/", createURLWithOriginal)
}
