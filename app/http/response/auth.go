package response

type LoginResponse struct {
	UserID       uint   `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	APIKey       string `json:"api_key"`
}
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}
