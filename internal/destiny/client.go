// Package apicfg ..
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
	HttpClient *http.Client
	Cfg        *config.Config
	Token      *models.Token
}

func NewClient(cfg *config.Config, token *models.Token) *Client {
	return &Client{
		HttpClient: &http.Client{},
		Cfg:        cfg,
		Token:      token,
	}
}

func (c *Client) DoRequest(method, url string) (*http.Response, error) {
	if method == "" {
		return nil, fmt.Errorf("HTTP method cannot be empty")
	}
	if url == "" {
		return nil, fmt.Errorf("URL cannot be empty")
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
	logger.GetLogger().Debug("Calling Bungie: %v", req)
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// If token expired (401)
	if resp.StatusCode == http.StatusUnauthorized {
		resp.Body.Close()

		logger.GetLogger().Warn("Token expired, attempting refresh")

		newToken, err := auth.RefreshToken(c.Cfg, c.Token)
		if err != nil {
			return nil, fmt.Errorf("session expired: %w", err)
		}

		c.Token = newToken

		req, err = makeReq()
		if err != nil {
			return nil, err
		}

		logger.GetLogger().Info("Token refreshed, retrying original request")
		return c.HttpClient.Do(req)
	}

	return resp, nil
}
