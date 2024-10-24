package client

import (
	"context"
	"fmt"
	"strings"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type OpenAIClient struct {
	Client *openai.Client
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	// Initialize the OpenAI SDK client
	return &OpenAIClient{
		Client: openai.NewClient(option.WithAPIKey(apiKey)),
	}
}

func (c *OpenAIClient) GenerateLinkedInPosts(topic string, numPosts int) ([]string, error) {
	// Create a custom prompt for LinkedIn posts
	prompt := fmt.Sprintf("You are an experienced software engineer and technical writer for a software company. Response should only be requested content string separated by '$$$$$'. This will be used to separate the posts. Strings should not have any other delimiter in the start or end or serial number, no summary at start or end, only postable content, it should be well formatted for posting on LinkedIn. Write brief, concise and professional, %d LinkedIn posts for description %s. Each post should be professional and informative with examples where relevant, and provide value to a technical audience. Each post should be self-explanatory and should not have past context linked. Do not repeat past responses.", numPosts, topic)

	// Call the OpenAI ChatCompletion API
	resp, err := c.Client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Model: openai.F(openai.ChatModelGPT4oMini),
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		Temperature: openai.Float(0.9),  // Adjust for creativity
		TopP:        openai.Float(0.95), // Adjust for response diversity
	})
	if err != nil {
		return nil, err
	}

	// Check if there are any responses
	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no choices received in response")
	}

	// Extract the content of the response
	message := resp.Choices[0].Message.Content

	// Split the message into individual posts based on the delimiter "$$$$$"
	posts := splitIntoPosts(message, numPosts)

	return posts, nil
}

// Helper function to split posts by the specified delimiter
func splitIntoPosts(text string, numPosts int) []string {
	var filteredPosts []string
	posts := strings.Split(text, "$$$$$")
	for _, suggestion := range posts {
		content := strings.ReplaceAll(suggestion, "\n", "")
		content = strings.ReplaceAll(content, "\r", "")
		content = strings.TrimSpace(content)
		if content != "" {
			filteredPosts = append(filteredPosts, content)
		}
	}

	// Trim to the requested number of posts
	if len(filteredPosts) > numPosts {
		filteredPosts = filteredPosts[:numPosts]
	}

	return filteredPosts
}
