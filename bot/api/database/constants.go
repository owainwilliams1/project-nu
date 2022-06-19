package database

type MemberType string

const (
	Player     MemberType = "player"
	Substitute MemberType = "substitute"
	Coach      MemberType = "coach"
)
