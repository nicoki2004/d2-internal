// Package auth ...
package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/nicoki2004/d2-internal/internal/config"
	"github.com/nicoki2004/d2-internal/internal/models"
)

// AuthURL builds the Bungie OAuth authorization URL.
func AuthURL(cfg config.Config) (string, error) {
	v := url.Values{}
	v.Set("client_id", cfg.ClientID)
	v.Set("response_type", "code")
	// v.Set("state", state)
	v.Set("redirect_uri", cfg.RedirectURL)

	return models.AUTH_URL_PREFIX + v.Encode(), nil
}

// ExchangeCode Change a code for a Token
func ExchangeCode(cfg *config.Config, code string) (*models.Token, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}
	if code == "" {
		return nil, fmt.Errorf("authorization code cannot be empty")
	}
	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)
	form.Set("client_id", cfg.ClientID)
	form.Set("client_secret", cfg.Secret)
	if cfg.RedirectURL != "" {
		form.Set("redirect_uri", cfg.RedirectURL)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", models.AUTH_TOKEN_URL_PREFIX, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if cfg.APIKey != "" {
		req.Header.Set("X-API-Key", cfg.APIKey)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		var errorDetail any
		if err := json.NewDecoder(resp.Body).Decode(&errorDetail); err != nil {
			errorDetail = "unable to parse error details"
		}
		return nil, fmt.Errorf("token exchange failed: %s - Detail: %v", resp.Status, errorDetail)
	}

	t, err := DecodeToken(resp)
	if err != nil {
		return nil, err
	}
	t.ReceivedAt = time.Now()
	return t, nil
}

// RefreshToken ... Refresh a token and return a new token
func RefreshToken(cfg *config.Config, oldToken *models.Token) (*models.Token, error) {
	if cfg == nil {
		return &models.Token{}, fmt.Errorf("config cannot be nil")
	}
	if oldToken == nil {
		return &models.Token{}, fmt.Errorf("oldToken cannot be nil")
	}
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", oldToken.RefreshToken)
	data.Set("client_id", cfg.ClientID)
	data.Set("client_secret", cfg.Secret)

	resp, err := http.PostForm("https://www.bungie.net/platform/app/oauth/token/", data)
	if err != nil {
		return &models.Token{}, err
	}

	newToken, err := DecodeToken(resp)
	if err != nil {
		return &models.Token{}, err
	}

	newToken.ReceivedAt = time.Now()

	return newToken, nil
}

// DecodeToken ...Decodes a Token
func DecodeToken(resp *http.Response) (*models.Token, error) {
	defer resp.Body.Close()

	var token models.Token

	err := json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar el token de Bungie: %w", err)
	}

	// Seteamos el momento exacto en que recibimos el JSON.
	// Esto hace que t.ReceivedAt.Add(t.ExpiresIn) sea un cálculo real.
	token.ReceivedAt = time.Now()

	return &token, nil
}
