package gobudins

import "time"

type WebhookUser struct {
	ID       int       `json:"id"`
	Signin   time.Time `json:"signin"`
	Platform int       `json:"platform"`
}
