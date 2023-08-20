package jwt

// GetJWTRefreshAndAccessTokens returns a RefreshAndAccessTokens struct with the provided refresh token and access token.
//
// Parameters:
// - refreshToken: a string representing the refresh token.
// - accessToken: a string representing the access token.
//
// Return:
// - RefreshAndAccessTokens: a struct containing the refresh token and access token.
func GetJWTRefreshAndAccessTokens(refreshToken string, accessToken string) RefreshAndAccessTokens {
	return RefreshAndAccessTokens{
		Refresh: refreshToken,
		Access:  accessToken,
	}
}

func GetJWTAccessTokens(accessToken string) AccessToken {
	return AccessToken{
		Access: accessToken,
	}
}
