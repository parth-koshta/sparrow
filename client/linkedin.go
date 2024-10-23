package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const REDIRECT_URI = "https://oauth.pstmn.io/v1/callback"
const GET_TOKEN_URL = "https://www.linkedin.com/oauth/v2/accessToken"
const GET_USER_INFO_URL = "https://api.linkedin.com/v2/userinfo"

type LinkedinClient struct {
	ClientID     string
	ClientSecret string
}

func NewLinkedInClient(clientID, clientSecret string) *LinkedinClient {
	return &LinkedinClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

type TokenInfo struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (lc *LinkedinClient) GetAccessToken(code string) (*TokenInfo, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", REDIRECT_URI)
	data.Set("client_id", lc.ClientID)
	data.Set("client_secret", lc.ClientSecret)

	req, err := http.NewRequest("POST", GET_TOKEN_URL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get access token: %s", body)
	}

	var tokenInfo TokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &tokenInfo, nil
}

type UserInfo struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (c *LinkedinClient) GetUserInfo(accessToken string) (*UserInfo, error) {
	req, err := http.NewRequest("GET", GET_USER_INFO_URL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("X-Restli-Protocol-Version", "2.0.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("LinkedIn API returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var userInfo UserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user info: %w", err)
	}

	return &userInfo, nil
}
