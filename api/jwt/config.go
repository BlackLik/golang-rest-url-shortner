package jwt

import (
	"time"

	"gorm.io/gorm"
	"urlshort.ru/m/config"
	"urlshort.ru/m/utils"
)

const LOGGER_HANDLER = "api.jwt"

var JWT_SECRET = utils.GenerateShortHashSHA256(config.ConfigAll.SECRET_KEY_JWT)

const REFRESH_TIME = time.Hour * 24 * 31
const ACCESS_TIME = time.Hour * 24
const TYPE_CHECK_PROTOCOL = "Bearer"
const ALGORITHM_JWT = "HS256"
const PROTOCOL_JWT = "JWT"

var localDb *gorm.DB
