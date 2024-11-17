package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

const (
	REDIRECT_URI      = "https://oauth.pstmn.io/v1/callback"
	GET_TOKEN_URL     = "https://www.linkedin.com/oauth/v2/accessToken"
	GET_USER_INFO_URL = "https://api.linkedin.com/v2/userinfo"
	PUBLISH_POST_URL  = "https://api.linkedin.com/v2/ugcPosts"
)

// LinkedinClientInterface defines the methods for interacting with the LinkedIn API.
type LinkedinAPIClient interface {
	GetAccessToken(code string) (*TokenInfo, error)
	GetUserInfo(accessToken string) (*UserInfo, error)
	PublishPost(accessToken, sub string, postPayload PayloadPublishPost) error
}

// LinkedinClient provides an implementation for interacting with the LinkedIn API.
type LinkedinClient struct {
	ClientID     string
	ClientSecret string
	HttpClient   *http.Client
}

// NewLinkedInClient initializes and returns a new LinkedinClient.
func NewLinkedInClient(clientID, clientSecret string) LinkedinAPIClient {
	return &LinkedinClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		HttpClient:   &http.Client{},
	}
}

// TokenInfo represents the access token response from LinkedIn.
type TokenInfo struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetAccessToken exchanges an authorization code for an access token.
func (lc *LinkedinClient) GetAccessToken(code string) (*TokenInfo, error) {
	data := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {REDIRECT_URI},
		"client_id":     {lc.ClientID},
		"client_secret": {lc.ClientSecret},
	}

	req, err := http.NewRequest("POST", GET_TOKEN_URL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create access token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := lc.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send access token request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("access token request failed: %s", body)
	}

	var tokenInfo TokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return nil, fmt.Errorf("failed to decode access token response: %w", err)
	}

	return &tokenInfo, nil
}

// UserInfo represents LinkedIn user information.
type UserInfo struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// GetUserInfo fetches user information for the given access token.
func (lc *LinkedinClient) GetUserInfo(accessToken string) (*UserInfo, error) {
	req, err := http.NewRequest("GET", GET_USER_INFO_URL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create user info request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("X-Restli-Protocol-Version", "2.0.0")

	resp, err := lc.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send user info request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("user info request failed with status %d: %s", resp.StatusCode, body)
	}

	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info response: %w", err)
	}

	return &userInfo, nil
}

// PayloadPublishPost represents the payload for publishing a LinkedIn post.
type PayloadPublishPost struct {
	PostID pgtype.UUID `json:"post_id"`
	Text   string      `json:"text"`
}

// PublishPost publishes a post on LinkedIn.
func (lc *LinkedinClient) PublishPost(accessToken, sub string, postPayload PayloadPublishPost) error {
	author := fmt.Sprintf("urn:li:person:%s", sub)
	post := map[string]interface{}{
		"author":         author,
		"lifecycleState": "PUBLISHED",
		"specificContent": map[string]interface{}{
			"com.linkedin.ugc.ShareContent": map[string]interface{}{
				"shareCommentary": map[string]string{
					"text": postPayload.Text,
				},
				"shareMediaCategory": "NONE",
			},
		},
		"visibility": map[string]string{
			"com.linkedin.ugc.MemberNetworkVisibility": "PUBLIC",
		},
	}

	payload, err := json.Marshal(post)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal post payload")
		return fmt.Errorf("failed to marshal post payload: %w", err)
	}

	req, err := http.NewRequest("POST", PUBLISH_POST_URL, bytes.NewBuffer(payload))
	if err != nil {
		log.Error().Err(err).Msg("failed to create publish post request")
		return fmt.Errorf("failed to create publish post request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := lc.HttpClient.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("failed to send publish post request")
		return fmt.Errorf("failed to send publish post request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		log.Error().Msgf("failed to publish post: %s", body)
		return fmt.Errorf("failed to publish post: %s", body)
	}

	log.Info().Msgf("post published successfully, post id: %s", uuid.UUID(postPayload.PostID.Bytes).String())
	return nil
}
