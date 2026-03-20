// Package destiny ..
package destiny

import (
	"fmt"
	"net/http"

	"github.com/nicoki2004/d2-internal/internal/auth"
	"github.com/nicoki2004/d2-internal/internal/config"
	"github.com/nicoki2004/d2-internal/internal/logger"
	"github.com/nicoki2004/d2-internal/internal/models"
)

type Client struct {
	HTTPClient *http.Client
	Cfg        *config.Config
	Token      *models.Token
}

func NewClient(cfg *config.Config, token *models.Token) *Client {
	return &Client{
		HTTPClient: &http.Client{},
		Cfg:        cfg,
		Token:      token,
	}
}

func (c *Client) DoRequest(method, url string) (*http.Response, error) {
	if method == "" || url == "" {
		return nil, fmt.Errorf("method and URL cannot be empty")
	}

	// Si el token vence en menos de 60s, lo refrescamos antes de la llamada principal.
	if c.Token.IsExpiredSoon() {
		logger.GetLogger().Warn("Token por expirar, refrescando preventivamente...")
		newToken, err := auth.RefreshToken(c.Cfg, c.Token)
		if err != nil {
			return nil, fmt.Errorf("error en refresh preventivo: %w", err)
		}
		c.Token = newToken
	}

	makeReq := func() (*http.Request, error) {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("X-API-Key", c.Cfg.APIKey)
		req.Header.Add("Authorization", "Bearer "+c.Token.AccessToken)
		return req, nil
	}

	req, err := makeReq()
	if err != nil {
		return nil, err
	}

	logger.GetLogger().Debug("Calling Bungie: %s %s", method, url)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		resp.Body.Close()
		logger.GetLogger().Warn("Token rechazado (401), intentando refresh reactivo...")

		newToken, err := auth.RefreshToken(c.Cfg, c.Token)
		if err != nil {
			return nil, fmt.Errorf("sesión expirada: %w", err)
		}

		c.Token = newToken
		req, _ = makeReq()

		logger.GetLogger().Info("Token refrescado, reintentando request original")
		return c.HTTPClient.Do(req)
	}

	return resp, nil
}
