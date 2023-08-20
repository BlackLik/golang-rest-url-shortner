package jwt

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"urlshort.ru/m/schema"
	"urlshort.ru/m/utils"
)

// GenerateJWTRefresh generates a JWT refresh token for the given user ID.
//
// Parameters:
// - userId: the ID of the user (int).
//
// Returns:
// - token: the generated JWT refresh token (string).
// - err: an error if there was a problem generating the token (error).
func GenerateJWTRefresh(userId int) (string, error) {
	headerMarshal, err := GetHeaderJWTJson()
	if err != nil {
		return "", err
	}

	payload := PayloadJWTRefresh{
		UserID: userId,
		EXP:    time.Now().Add(REFRESH_TIME).Unix(),
	}

	payloadMarshal, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	resultToken := GenerateJWT(string(headerMarshal), string(payloadMarshal))
	return resultToken, nil
}

// GenerateJWTAccess generates a JWT access token for the given user ID, email, and role.
//
// Parameters:
// - userId: The ID of the user.
// - email: The email of the user.
// - role: The role of the user.
//
// Returns:
// - string: The generated JWT access token.
// - error: An error if the JWT access token generation fails.
func GenerateJWTAccess(userId int, email string, role string) (string, error) {
	headerMarshal, err := GetHeaderJWTJson()
	if err != nil {
		return "", err
	}

	payload := PayloadJWTAccess{
		UserID: userId,
		EXP:    time.Now().Add(ACCESS_TIME).Unix(),
		Email:  email,
		Role:   role,
	}
	payloadMarshal, err := json.Marshal(payload)

	if err != nil {
		return "", err
	}
	resultToken := GenerateJWT(string(headerMarshal), string(payloadMarshal))
	return resultToken, nil
}

// GenerateSignatureJWT generates a signature JWT.
//
// Parameters:
// - header: the header of the JWT.
// - payload: the payload of the JWT.
//
// Returns:
// - string: the signature JWT.
func GenerateSignatureJWT(header string, payload string) string {
	return utils.GenerateShortHashSHA256(header + payload + JWT_SECRET)
}

// GenerateJWT generates a JSON Web Token (JWT) using the provided header and payload.
//
// Parameters:
// - header: the header string for the JWT.
// - payload: the payload string for the JWT.
//
// Returns:
// - string: the JWT generated from the provided header and payload.
func GenerateJWT(header string, payload string) string {

	header = string(utils.Base64Encode(header))
	payload = string(utils.Base64Encode(payload))

	signature := GenerateSignatureJWT(header, payload)

	resultToken := header + "." + payload + "." + signature
	return resultToken
}

// GetHeaderJWTJson generates the JSON representation of the JWT header.
//
// It does not take any parameters.
// It returns a byte slice containing the JSON representation of the header and an error if there is any.
func GetHeaderJWTJson() ([]byte, error) {
	header := HeaderJWT{
		Algorithm: ALGORITHM_JWT,
		Protocol:  PROTOCOL_JWT,
	}

	headerMarshal, err := json.Marshal(header)
	if err != nil {
		return []byte{}, err
	}
	return headerMarshal, nil

}

// ExtractToken extracts the token from the authorization header.
//
// It takes a string parameter called authorization which represents the authorization header.
// It returns a string which represents the extracted token and an error if the token is invalid.
func ExtractToken(authorization string) (string, error) {
	parts := strings.SplitN(authorization, " ", 2)
	if len(parts) != 2 {
		return "", errors.New("invalid token")
	}
	return parts[1], nil
}

// CheckToken checks if the given token is valid.
//
// It takes a string parameter named "token" and returns a boolean value.
func CheckToken(token string) bool {
	if token == "" {
		return false
	}

	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 3 {
		return false
	}

	header, payload, signature := tokenParts[0], tokenParts[1], tokenParts[2]
	if header == "" || payload == "" || signature == "" {
		return false
	}

	if !CheckSignature(header, payload, signature) {
		return false
	}

	decodedHeader, err := utils.Base64Decode([]byte(header))
	if err != nil || !CheckProtocol(decodedHeader) {
		return false
	}

	decodedPayload, err := utils.Base64Decode([]byte(payload))
	if err != nil || !CheckPayload(decodedPayload) {
		return false
	}

	return true
}

// CheckSignature checks if the given header, payload, and signature match.
//
// Parameters:
// - header: the header string.
// - payload: the payload string.
// - signature: the signature string.
//
// Returns:
// - a boolean indicating whether the signature is valid.
func CheckSignature(header string, payload string, signature string) bool {
	return utils.GenerateShortHashSHA256(header+payload+JWT_SECRET) == signature
}

// CheckProtocol checks if the given header matches the PROTOCOL_JWT constant.
//
// header: a string representing the header to be checked.
// returns: a boolean indicating if the header matches the PROTOCOL_JWT constant.
func CheckProtocol(header string) bool {

	if header == "" {
		return false
	}
	var parseJson HeaderJWT
	err := json.Unmarshal([]byte(header), &parseJson)
	if err != nil {
		return false
	}

	return parseJson.Protocol == PROTOCOL_JWT
}

// CheckPayload checks the validity of a payload.
//
// It takes a payload string as a parameter and returns a boolean value indicating
// whether the payload is valid or not.
func CheckPayload(payload string) bool {
	dictonary := make(map[string]any)
	err := json.Unmarshal([]byte(payload), &dictonary)
	if err != nil {
		return false
	}
	if dictonary["exp"] == nil {
		return false
	}
	if int64(dictonary["exp"].(float64)) < time.Now().Unix() {
		return false
	}
	return true
}

// GetPayloadRefresh retrieves the payload from a JWT refresh token.
//
// The function takes a string parameter `payload` representing the JWT payload
// and returns a `PayloadJWTRefresh` object and an error.
func GetPayloadRefresh(payload string) (PayloadJWTRefresh, error) {
	if payload == "" {
		return PayloadJWTRefresh{}, errors.New("invalid payload")
	}
	if payload == "{}" {
		return PayloadJWTRefresh{}, nil
	}

	var dictonary PayloadJWTRefresh
	err := json.Unmarshal([]byte(payload), &dictonary)
	if err != nil {
		return PayloadJWTRefresh{}, err
	}

	if dictonary.UserID == 0 && dictonary.EXP == 0 {
		return PayloadJWTRefresh{}, errors.New("invalid payload")
	}

	return dictonary, nil
}

func GetPayloadAccess(payload string) (PayloadJWTAccess, error) {
	if payload == "" {
		return PayloadJWTAccess{}, errors.New("invalid payload")
	}
	if payload == "{}" {
		return PayloadJWTAccess{}, nil
	}
	var dictonary PayloadJWTAccess
	err := json.Unmarshal([]byte(payload), &dictonary)
	if err != nil {
		return PayloadJWTAccess{}, err
	}
	if dictonary.UserID == 0 && dictonary.EXP == 0 {
		return PayloadJWTAccess{}, errors.New("invalid payload")
	}
	return dictonary, nil
}

// ExtractTokenHandler extracts a token from the authorization header.
//
// The authorization parameter is a string representing the authorization header.
// It returns a string representing the extracted token and an error if any occurred.
func ExtractTokenHandler(c *fiber.Ctx) (string, error) {
	return GetExtractTokenHandler(c.GetReqHeaders()["Authorization"])
}

func checkErrorTokenHandler(token string, err error) (string, error) {
	if err != nil {
		return "", err
	}
	if !CheckToken(token) {
		return "", errors.New("invalid token")
	}
	return token, nil
}

func GetExtractTokenHandler(authorization string) (string, error) {
	token, err := ExtractToken(authorization)
	return checkErrorTokenHandler(token, err)
}

func checkErrorQueryDB(c *fiber.Ctx, result *gorm.DB) error {
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			utils.LoggerRequestUser(c, LOGGER_HANDLER, 404)
			return c.Status(404).JSON(schema.GetError404Response())
		}
		utils.LoggerRequestUser(c, LOGGER_HANDLER, 400)
		return c.Status(400).JSON(schema.GetError400Response())
	}
	return nil
}
