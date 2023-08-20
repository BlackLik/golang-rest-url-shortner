package jwt

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"urlshort.ru/m/models"
	"urlshort.ru/m/schema"
	"urlshort.ru/m/utils"
)

// @Summary Refresh JWT token
// @Description Refreshes the JWT token using the refresh token
// @Tags JWT
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {refresh_token}"
// @Success 200 {object} AccessToken
// @Failure 400 {object} schema.Response
// @Failure 401 {object} schema.Response
// @Failure 404 {object} schema.Response
// @Router /api/jwt/refresh [get]
//
// Parameters:
// - c: Указатель на объект fiber.Ctx, представляющий контекст запроса.
// Return type: error. Объект ошибки, если произошла ошибка при обработке запроса, в противном случае nil.
func refreshHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("UTF-8", "UTF-16")
	token, err := ExtractTokenHandler(c)
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	tokenSplit := strings.Split(token, ".")
	decodeToken, err := utils.Base64Decode([]byte(tokenSplit[1]))
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}
	payloadRefresh, err := GetPayloadRefresh(decodeToken)

	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	var user models.User
	result := localDb.Where("id = ? AND refresh_token = ?", payloadRefresh.UserID, utils.GenerateShortHashSHA256(token)).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			utils.LoggerRequestUser(c, LOGGER_HANDLER, 404)
			return c.Status(404).JSON(schema.GetError404Response())
		}
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 401)
		return c.Status(401).JSON(schema.GetError401Response())
	}

	accessToken, err := GenerateJWTAccess(int(user.ID), user.Email, user.Role)
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	utils.LoggerRequestUser(c, LOGGER_HANDLER, 200)
	return c.JSON(GetJWTAccessTokens(accessToken))

}

// @Summary Register user
// @Description Registers a new user
// @Tags JWT
// @Accept json
// @Produce json
// @Param requestBody body UserJSON true "User object"
// @Success 200 {object} RefreshAndAccessTokens
// @Failure 400 {object} schema.Response
// @Router /api/jwt/register [post]
//
// Parameters:
// - c: Указатель на объект fiber.Ctx, представляющий контекст запроса.
// Return type: error. Объект ошибки, если произошла ошибка при обработке запроса, в противном случае nil.
func registerHandler(c *fiber.Ctx) error {

	c.Accepts("application/json")
	c.AcceptsCharsets("UTF-8", "UTF-16")
	inputUserJson := new(UserJSON)
	if err := c.BodyParser(inputUserJson); err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	var user models.User
	user.Email = inputUserJson.Email
	user.Password = utils.GenerateShortHashSHA256(inputUserJson.Email + inputUserJson.Password)
	user.Role = "user"
	if err := localDb.Create(&user).Error; err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	refreshToken, err := GenerateJWTRefresh(int(user.ID))
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	user.RefreshToken = utils.GenerateShortHashSHA256(refreshToken)
	if err := localDb.Save(&user).Error; err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	accessToken, err := GenerateJWTAccess(int(user.ID), user.Email, user.Role)
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	utils.LoggerRequestUser(c, LOGGER_HANDLER, 200)
	return c.JSON(GetJWTRefreshAndAccessTokens(refreshToken, accessToken))
}

// @Summary User login
// @Description Logs in a user
// @Tags JWT
// @Accept json
// @Produce json
// @Param requestBody body UserJSON true "User object"
// @Success 200 {object} RefreshAndAccessTokens
// @Failure 400 {object} schema.Response
// @Failure 401 {object} schema.Response
// @Router /api/jwt/login [post]
//
// Parameters:
// - c: Указатель на объект fiber.Ctx, представляющий контекст запроса.
// Return type: error. Объект ошибки, если произошла ошибка при обработке запроса, в противном случае nil.
func loginHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("UTF-8", "UTF-16")
	inputUserJson := new(UserJSON)
	if err := c.BodyParser(inputUserJson); err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}
	var user models.User
	passwordDB := utils.GenerateShortHashSHA256(inputUserJson.Email + inputUserJson.Password)
	if err := localDb.Where("email = ? AND password = ?", inputUserJson.Email, passwordDB).First(&user).Error; err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 401)
		return c.Status(401).JSON(schema.GetError401Response())
	}
	refreshToken, err := GenerateJWTRefresh(int(user.ID))
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	user.RefreshToken = utils.GenerateShortHashSHA256(refreshToken)
	if err := localDb.Save(&user).Error; err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	accessToken, err := GenerateJWTAccess(int(user.ID), user.Email, user.Role)
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}
	utils.LoggerRequestUser(c, LOGGER_HANDLER, 200)
	return c.JSON(GetJWTRefreshAndAccessTokens(refreshToken, accessToken))
}

// @Summary User logout
// @Description Logs out a user
// @Tags JWT
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} RefreshAndAccessTokens
// @Failure 400 {object} schema.Response
// @Failure 401 {object} schema.Response
// @Router /api/jwt/logout [get]
//
// Parameters:
// - c: Указатель на объект fiber.Ctx, представляющий контекст запроса.
// Return type: error. Объект ошибки, если произошла ошибка при обработке запроса, в противном случае nil.
func logoutHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("UTF-8", "UTF-16")
	token, err := ExtractTokenHandler(c)
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	tokenSplit := strings.Split(token, ".")
	decodeToken, err := utils.Base64Decode([]byte(tokenSplit[1]))
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}
	payloadRefresh, err := GetPayloadRefresh(decodeToken)
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	var user models.User
	result := localDb.Where("id = ? AND refresh_token = ?", payloadRefresh.UserID, utils.GenerateShortHashSHA256(token)).First(&user)

	err = CheckErrorQueryDB(c, result)
	if err != nil {
		return err
	}

	user.RefreshToken = ""
	if err := localDb.Save(&user).Error; err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	utils.LoggerRequestUser(c, LOGGER_HANDLER, 200)
	return c.JSON(GetJWTRefreshAndAccessTokens("", ""))
}

// @Summary Check token
// @Description Checks the validity of a token
// @Tags JWT
// @Accept json
// @Produce json
// @Param requestBody body CheckTokenJSON true "Token object"
// @Success 200 {object} string
// @Failure 400 {object} schema.Response
// @Router /api/jwt/check [post]
//
// Parameters:
// - c: Указатель на объект fiber.Ctx, представляющий контекст запроса.
// Return type: error. Объект ошибки, если произошла ошибка при обработке запроса, в противном случае nil.
func checkHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("UTF-8", "UTF-16")

	body := new(CheckTokenJSON)
	if err := c.BodyParser(body); err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}
	if body.Token == "" {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	_, err := GetExtractTokenHandler(body.Token)
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}
	utils.LoggerRequestUser(c, LOGGER_HANDLER, 200)
	return c.SendStatus(200)
}

// @Summary Delete user
// @Description Deletes a user
// @Tags JWT
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} string
// @Failure 400 {object} schema.Response
// @Failure 400 {object} schema.Response
// @Failure 404 {object} schema.Response
// @Router /api/jwt/delete [delete]
//
// Parameters:
// - c: Указатель на объект fiber.Ctx, представляющий контекст запроса.
// Return type: error. Объект ошибки, если произошла ошибка при обработке запроса, в противном случае nil.
func deleteHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")
	c.AcceptsCharsets("UTF-8", "UTF-16")

	body := new(UserDelete)
	if err := c.BodyParser(body); err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}
	if body.UserId == 0 {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	token, err := ExtractTokenHandler(c)
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 401)
		return c.Status(401).JSON(schema.GetError401Response())
	}

	tokenSplit := strings.Split(token, ".")
	payloadAccess, err := GetPayloadAccess(tokenSplit[1])
	if err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	var user models.User
	result := localDb.Where("id = ?", payloadAccess.UserID, utils.GenerateShortHashSHA256(token)).First(&user)
	err = CheckErrorQueryDB(c, result)
	if err != nil {
		return err
	}

	if user.Role != "admin" {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	result = localDb.First(&user, "id = ?", body.UserId)
	err = CheckErrorQueryDB(c, result)
	if err != nil {
		return err
	}
	if err = localDb.Delete(&user, "id = ?", body.UserId).Error; err != nil {
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}

	utils.LoggerRequestUser(c, LOGGER_HANDLER, 200)
	return c.SendStatus(200)
}
