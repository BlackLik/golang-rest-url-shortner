package urls

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
	"urlshort.ru/m/api/jwt"
	"urlshort.ru/m/models"
	"urlshort.ru/m/schema"
	"urlshort.ru/m/utils"
)

// getURLWithShort обрабатывает HTTP-запрос для получения параметров URL.
//
// @Summary Получить параметры URL
// @Description Обрабатывает HTTP-запрос для получения параметров URL.
// @Tags Параметры URL
// @Accept json
// @Produce json
// @Param shorturl path string true "Короткий URL"
// @Success 200 {object} URLResponse
// @Failure 400 {object} schema.Response
// @Failure 404 {object} schema.Response
// @Router /api/urls/{shorturl} [get]
//
// Parameters:
// - c: Указатель на объект fiber.Ctx, представляющий контекст HTTP-запроса.
// Return type: error. Объект ошибки, если произошла ошибка при обработке запроса, в противном случае nil.
func getURLWithShort(c *fiber.Ctx) error {
	c.Accepts("application/json")
	// TODO Logger handler ip address response url path
	var url models.URL
	result := localDb.First(&url, "short_url = ?", c.Params("shorturl"))
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			utils.LoggerRequestUser(c, LOGGER_HANDLER, 404)
			return c.Status(404).JSON(GetError404Response())
		}
		slog.Debug(LOGGER_HANDLER, result.Error)
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(GetError400Response())
	}
	utils.LoggerRequestUser(c, LOGGER_HANDLER, 200)
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
// @Failure 400 {object} schema.Response
// @Router /api/urls/ [post]
//
// Parameters:
// - c: Указатель на объект fiber.Ctx, представляющий контекст запроса.
// Return type: error. Объект ошибки, если произошла ошибка при обработке запроса, в противном случае nil.
func createURLWithOriginal(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("UTF-8", "UTF-16")
	inputJson := new(CreateURLBody)
	if err := c.BodyParser(inputJson); err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(GetError400Response())
	}

	var url models.URL
	newShortUrl := utils.GenerateShortHashMD5(inputJson.OriginalURL)
	url.OriginalURL = inputJson.OriginalURL
	url.ShortURL = newShortUrl
	url.CreatedAt = time.Now()

	result := localDb.Create(&url)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "UNIQUE constraint failed") {
			localDb.First(&url, "original_url = ?", inputJson.OriginalURL)
			utils.LoggerRequestUser(c, LOGGER_HANDLER, 200)
			return c.JSON(GetURLResponse(url))
		}
		slog.Debug(LOGGER_HANDLER, result.Error)
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(GetError400Response())
	}

	url.ShortURL = utils.Conver10IntTo32String(int64(url.ID))
	localDb.Save(&url)

	slog.Debug(LOGGER_HANDLER, url)
	utils.LoggerRequestUser(c, LOGGER_HANDLER, 200)
	return c.JSON(GetURLResponse(url))
}

// deleteURLWithShort deletes a URL with a given short code.
//
// This function handles the DELETE /url endpoint and deletes the URL associated
// with the specified short code. It requires the request body to be in JSON format
// and the "Content-Type" header to be set to "application/json".
// @Summary Удалить URL
// @Description Удалить URL с предоставленным коротким URL.
// @Tags Параметры URL
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param shorturl path string true "Короткий URL"
// @Success 200 {object} URLResponse
// @Failure 400 {object} schema.Response
// @Failure 401 {object} schema.Response
// @Failure 404 {object} schema.Response
// @Router /api/urls/{shorturl} [delete]
//
// Parameters:
// - c (*fiber.Ctx): the Fiber context object representing the HTTP request and response.
//
// Returns:
// - error: an error if any occurred during the processing of the request.
func deleteURLWithShort(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("UTF-8", "UTF-16")

	if c.Params("shorturl") == "" {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	token, err := jwt.ExtractTokenHandler(c)
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 401)
		return c.Status(401).JSON(schema.GetError401Response())
	}

	resultPayload, err := jwt.GetPayloadHandlerAccess(token)
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 401)
		return c.Status(401).JSON(schema.GetError401Response())
	}
	if resultPayload.Role != models.ROLE_ADMIN {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	var user models.User

	result := localDb.First(&user, "id = ?", resultPayload.UserID)
	err = jwt.CheckErrorQueryDB(c, result)
	if err != nil {
		return err
	}

	var url models.URL

	result = localDb.First(&url, "short_url = ? and deleted_at IS NULL", c.Params("shorturl"))
	if result.Error != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 404)
		return c.Status(404).JSON(schema.GetError404Response())
	}
	result = localDb.Where("short_url = ? AND deleted_at IS NULL", c.Params("shorturl")).Delete(&url)

	slog.Debug(LOGGER_HANDLER, url)

	if result.Error != nil {
		slog.Debug(LOGGER_HANDLER, result.Error)
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	utils.LoggerRequestUser(c, LOGGER_HANDLER, 200)
	return c.JSON(schema.GetSuccess200Response())
}

// updateURLWithShort deletes a URL with a given short code.
//
// This function handles the DELETE /url endpoint and deletes the URL associated
// with the specified short code. It requires the request body to be in JSON format
// and the "Content-Type" header to be set to "application/json".
// @Summary Обновить URL
// @Description Обновить URL с предоставленным коротким и существующим URL.
// @Tags Параметры URL
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param shorturl path string true "Короткий URL"
// @Param bodyJson body ShortURLBody true "Original URL"
// @Success 200 {object} URLResponse
// @Failure 400 {object} schema.Response
// @Failure 401 {object} schema.Response
// @Failure 404 {object} schema.Response
// @Router /api/urls/{shorturl} [patch]
//
// Parameters:
// - c (*fiber.Ctx): the Fiber context object representing the HTTP request and response.
//
// Returns:
// - error: an error if any occurred during the processing of the request.
func updateURLWithShort(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("UTF-8", "UTF-16")

	if c.Params("shorturl") == "" {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	bodyJson := new(ShortURLBody)
	if err := c.BodyParser(bodyJson); err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	if bodyJson.OriginalURL == "" {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	token, err := jwt.ExtractTokenHandler(c)
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 401)
		return c.Status(401).JSON(schema.GetError401Response())
	}

	resultPayload, err := jwt.GetPayloadHandlerAccess(token)
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 401)
		return c.Status(401).JSON(schema.GetError401Response())
	}
	if resultPayload.Role != models.ROLE_ADMIN {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	var user models.User

	result := localDb.First(&user, "id = ?", resultPayload.UserID)
	err = jwt.CheckErrorQueryDB(c, result)
	if err != nil {
		return err
	}

	var url models.URL

	result = localDb.First(&url, "short_url = ?", c.Params("shorturl"))
	if result.Error != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 404)
		return c.Status(404).JSON(schema.GetError404Response())
	}

	url.OriginalURL = bodyJson.OriginalURL

	result = localDb.Save(&url)
	if result.Error != nil {
		slog.Debug(LOGGER_HANDLER, result.Error)
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	utils.LoggerRequestUser(c, LOGGER_HANDLER, 200)
	return c.JSON(schema.GetSuccess200Response())
}
