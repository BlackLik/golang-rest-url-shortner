package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

// GenerateShortHashMD5 генерирует очень короткий хеш из строки.
//
// Параметры:
// - input: Исходная строка, для которой нужно сгенерировать хеш.
//
// Возвращаемый тип:
// - string: Сгенерированный очень короткий хеш в виде 32-ричного кода.
func GenerateShortHashMD5(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

// GenerateShortHashSHA256 generates a short SHA256 hash for the given input string.
//
// input: the string to be hashed.
// returns: the generated short hash.
func GenerateShortHashSHA256(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

// LoggerRequestUser logs the user request.
//
// Parameters:
//   - c (*fiber.Ctx): The fiber context object.
//   - loggerHandler (string): The logger handler.
func LoggerRequestUser(c *fiber.Ctx, loggerHandler string, status int) {
	slog.Info(loggerHandler, "ip", c.IP(), "path", c.Path(), "method", c.Method(), "status", status)
}

// Base64Encode encodes the given input string to base64.
//
// Parameters:
// - input: the string to be encoded.
//
// Returns:
// - encoded: the base64 encoded string as a byte slice.
func Base64Encode(input string) []byte {
	// encoded := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	// base64.StdEncoding.Encode(encoded, []byte(input))
	// return encoded
	return []byte(base64.StdEncoding.EncodeToString([]byte(input)))
}

// Base64Decode decodes a base64 encoded byte array and returns the decoded string.
//
// The function takes a parameter named "input" of type []byte, which represents the base64 encoded byte array.
// The function returns a string and an error. The string is the decoded version of the input byte array, and the error is nil if the decoding is successful, otherwise it contains the error encountered during decoding.
func Base64Decode(input []byte) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

// Conver10IntTo32String converts the given 10-based integer to a 32-based string.
//
// It takes an int64 as input parameter and returns a string.
func Conver10IntTo32String(input int64) string {
	return strconv.FormatInt(input, 32)
}
