package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"golang.org/x/exp/slog"
	"urlshort.ru/m/api"
	"urlshort.ru/m/config"
	"urlshort.ru/m/docs"
	_ "urlshort.ru/m/models"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email info@mangi.org
// @license.name MIT License
// @license.url https://opensource.org/license/mit/
// @host localhost:8080
// @BasePath /
func main() {
	config := config.ConfigAll
	slog.Debug("config", config)

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample swagger for Fiber"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// TODO init start fiber
	app := fiber.New()
	app.Use(cors.New())
	api.Register(app)

	app.Get("/docs/*", swagger.New(swagger.Config{
		Title:        "Swagger Example API",
		DocExpansion: "full",
	}))

	slog.Error("Error", app.Listen(":8080"))

	// TODO init routes

}
