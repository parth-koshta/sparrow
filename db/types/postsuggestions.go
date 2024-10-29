package db

type PostSuggestionStatus string

const (
	PostSuggestionStatusPending  PostSuggestionStatus = "pending"
	PostSuggestionStatusAccepted PostSuggestionStatus = "accepted"
)
