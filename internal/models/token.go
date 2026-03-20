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

func (t *Token) IsExpiredSoon() bool {
	// Calculamos el momento de expiración:
	// El momento en que llegó + la duración que nos dio Bungie (en segundos)
	expiryTime := t.ReceivedAt.Add(time.Duration(t.ExpiresIn) * time.Second)

	// Si "ahora + 60 segundos" es después del tiempo de expiración,
	// significa que al token le queda menos de un minuto de vida.
	return time.Now().Add(60 * time.Second).After(expiryTime)
}
