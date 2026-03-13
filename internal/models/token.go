// Package models...
package models

import "time"

type Token struct {
	AccessToken      string    `json:"access_token"`
	TokenType        string    `json:"token_type"`
	ExpiresIn        int       `json:"expires_in"`
	RefreshToken     string    `json:"refresh_token"`
	RefreshExpiresIn int       `json:"refresh_expires_in"`
	MembershipID     string    `json:"membership_id"`
	DisplayName      string    `json:"display_name,omitempty"`
	ReceivedAt       time.Time `json:"-"`
}
