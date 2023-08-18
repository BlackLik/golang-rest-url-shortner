package urls

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
	"urlshort.ru/m/models"
	"urlshort.ru/m/utils"
)

var localDb *gorm.DB
var loggerHandler string = "api.urls"

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
	apiUrls.Post("/", createURLWithOriginal)
	// TODO api.Delete("/:shorturl", deleteURLWithShort)
	// TODO api.Patch("/:shorturl", updateURLWithShort)
}

// getURLWithShort обрабатывает HTTP-запрос для получения параметров URL.
//
// @Summary Получить параметры URL
// @Description Обрабатывает HTTP-запрос для получения параметров URL.
// @Tags Параметры URL
// @Accept json
// @Produce json
// @Param shorturl path string true "Короткий URL"
// @Success 200 {object} URLResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/urls/{shorturl} [get]
//
// Parameters:
// - c: Указатель на объект fiber.Ctx, представляющий контекст HTTP-запроса.
// Return type: error. Объект ошибки, если произошла ошибка при обработке запроса, в противном случае nil.
func getURLWithShort(c *fiber.Ctx) error {
	c.Accepts("application/json")
	// TODO Logger handler ip address response url path
	utils.LoggerRequestUser(c, loggerHandler)
	var url models.URL
	result := localDb.First(&url, "short_url = ?", c.Params("shorturl"))
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(GetError404Response())
		}
		slog.Debug(loggerHandler, result.Error)
		return c.Status(400).JSON(GetError400Response())
	}
	return c.JSON(GetURLResponse(url))
}

// createURLWithOriginal создает URL с предоставленным исходным URL.
//
// @Summary Создать URL
// @Description Создает URL с предоставленным исходным URL.
// @Tags Параметры URL
// @Accept json
// @Produce json
// @Param c body CreateURLBody true "Тело запроса"
// @Success 200 {object} URLResponse
// @Failure 400 {object} ErrorResponse
// @Router /api/urls/ [post]
//
// Parameters:
// - c: Указатель на объект fiber.Ctx, представляющий контекст запроса.
// Return type: error. Объект ошибки, если произошла ошибка при обработке запроса, в противном случае nil.
func createURLWithOriginal(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("UTF-8", "UTF-16")
	utils.LoggerRequestUser(c, loggerHandler)

	inputJson := new(CreateURLBody)

	if err := c.BodyParser(inputJson); err != nil {
		return c.Status(400).JSON(GetError400Response())
	}

	var url models.URL

	newShortUrl := utils.GenerateShortHash(inputJson.OriginalURL)
	url.OriginalURL = inputJson.OriginalURL
	url.ShortURL = newShortUrl
	url.CreatedAt = time.Now()

	result := localDb.Create(&url)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "UNIQUE constraint failed") {
			localDb.First(&url, "original_url = ?", inputJson.OriginalURL)
			return c.JSON(GetURLResponse(url))
		}
		slog.Debug(loggerHandler, result.Error)
		return c.Status(400).JSON(GetError400Response())
	}
	slog.Debug(loggerHandler, result)
	return c.JSON(GetURLResponse(url))
}
