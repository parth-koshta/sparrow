package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const LINKEDIN_TOKEN_URL = "https://www.linkedin.com/oauth/v2/accessToken"

type LinkedInClient struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

func NewLinkedInClient(clientID, clientSecret, redirectURI string) *LinkedInClient {
	return &LinkedInClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
	}
}

func (lc *LinkedInClient) GetAccessToken(code string) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", lc.RedirectURI)
	data.Set("client_id", lc.ClientID)
	data.Set("client_secret", lc.ClientSecret)

	req, err := http.NewRequest("POST", LINKEDIN_TOKEN_URL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to get access token: %s", body)
	}

	var response struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return response.AccessToken, nil
}
