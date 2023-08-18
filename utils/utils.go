package utils

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

// GenerateShortHash генерирует очень короткий хеш из строки.
//
// Параметры:
// - input: Исходная строка, для которой нужно сгенерировать хеш.
//
// Возвращаемый тип:
// - string: Сгенерированный очень короткий хеш в виде 32-ричного кода.
func GenerateShortHash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

// LoggerRequestUser logs the user request.
//
// Parameters:
//   - c (*fiber.Ctx): The fiber context object.
//   - loggerHandler (string): The logger handler.
func LoggerRequestUser(c *fiber.Ctx, loggerHandler string) {
	slog.Info(loggerHandler, "ip", c.IP(), "path", c.Path())
}
