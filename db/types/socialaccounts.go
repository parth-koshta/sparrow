package db

type SocialPlatform string

const (
	SocialPlatformLinkedIn SocialPlatform = "linkedin"
)

func (s SocialPlatform) String() string {
	return string(s)
}
