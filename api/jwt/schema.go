package jwt

type UserJSON struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDelete struct {
	UserId int `json:"user_id"`
}

type HeaderJWT struct {
	Algorithm string `json:"alg"`
	Protocol  string `json:"typ"`
}

type PayloadJWTRefresh struct {
	UserID int   `json:"user_id"`
	EXP    int64 `json:"exp"`
}

type PayloadJWTAccess struct {
	UserID int    `json:"user_id"`
	EXP    int64  `json:"exp"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

type RefreshToken struct {
	Refresh string `json:"refresh"`
}

type AccessToken struct {
	Access string `json:"access"`
}

type CheckTokenJSON struct {
	Token string `json:"token"`
}

type RefreshAndAccessTokens struct {
	Refresh string `json:"refresh"`
	Access  string `json:"access"`
}
