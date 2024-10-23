package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const BASE_URL = "https://api.openai.com/v1"

type OpenAIClient struct {
	APIKey string
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{
		APIKey: apiKey,
	}
}

type ChatCompletionRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature,omitempty"`
	TopP        float64   `json:"top_p,omitempty"`
}

type OpenAIResponse struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	SystemFingerprint string   `json:"system_fingerprint"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	Logprobs     *int    `json:"logprobs,omitempty"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens            int                     `json:"prompt_tokens"`
	CompletionTokens        int                     `json:"completion_tokens"`
	TotalTokens             int                     `json:"total_tokens"`
	CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"`
}

type CompletionTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"`
}

func (c *OpenAIClient) GenerateLinkedInPosts(topic string, numPosts int) ([]string, error) {
	url := fmt.Sprintf("%s/chat/completions", BASE_URL)

	// Create the request body with a custom prompt for LinkedIn posts
	requestBody := ChatCompletionRequest{
		Model: "gpt-4o-mini", // Replace with the model you are using
		Messages: []Message{
			{Role: "system", Content: "You are a content generation assistant specialized in crafting engaging social media posts."},
			{Role: "user", Content: fmt.Sprintf("Write %d LinkedIn posts about %s. Each post should be professional, informative, and provide value to a technical audience.", numPosts, topic)},
		},
		Temperature: 0.8, // Higher value for more creativity
		TopP:        0.9, // Adjust to increase diversity
	}

	// Marshal the body to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// Create an HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Add headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and return the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response JSON into the OpenAIResponse struct
	var response OpenAIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	// Extract the generated text from the response
	if len(response.Choices) == 0 {
		return nil, fmt.Errorf("no choices received in response")
	}
	message := response.Choices[0].Message.Content

	// Split the message into individual posts based on newlines (or a different delimiter if needed)
	posts := splitIntoPosts(message, numPosts)

	return posts, nil
}

func splitIntoPosts(text string, numPosts int) []string {
	// Split the response into individual posts by double newlines
	posts := strings.Split(text, "\n\n")

	// If we have more posts than needed, trim the slice
	if len(posts) > numPosts {
		posts = posts[:numPosts]
	}

	return posts
}
