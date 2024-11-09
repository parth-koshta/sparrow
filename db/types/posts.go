package db

type PostStatus string

const (
	PostStatusDraft     PostStatus = "draft"
	PostStatusScheduled PostStatus = "scheduled"
	PostStatusEnqueued  PostStatus = "enqueued"
	PostStatusPublished PostStatus = "published"
)
