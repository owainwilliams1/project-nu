package database

type MemberType string

const (
	Player     MemberType = "player"
	Substitute MemberType = "substitute"
	Coach      MemberType = "coach"
)

type UsernameType string

const (
	ProjectNu   UsernameType = "project-nu"
	Valorant    UsernameType = "valorant"
	ApexLegends UsernameType = "apex-legends"
)
