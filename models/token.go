package models

type RefreshTokenRequest struct {
	Username     string `json:"username"`
	Fullname     string `json:"full_name"`
	AccessToken  string
	RefreshToken string
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}
