package gobudins

import "time"

const (
	RouteAccessToken   = "/auth/token/access"
	RouteUsers         = "/users"
	RouteAuthTokenCode = "/auth/token/code"
)

type ErrorResponse struct {
	Code        string `json:"error"`
	Description string `json:"error_description"`
}

type ConnectCallbackData struct {
	Code         string `json:"code"`
	ConnectionID string `json:"connectionID"`
}

type AskForToken struct {
	Code         string `json:"code"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type GetAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

const UserMe = "me"

type User struct {
	ID       int       `json:"id"`
	Signin   time.Time `json:"signin"`
	Platform int       `json:"platform"`
}
