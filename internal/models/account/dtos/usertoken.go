package dtos

type UserToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Role        string `json:"role"`
	ExpiresIn   string `json:"expires_in"`
	IssuedAt    string `json:"issued_at"`
}
