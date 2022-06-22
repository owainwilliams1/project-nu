package types

import (
	"hushclan.com/pkg/utils"
)

type Team struct {
	TeamID      string   `firestore:"-"`
	OwnerID     string   `firestore:"owner_id"`
	Color       int      `firestore:"color"`
	Members     []string `firestore:"members"`
	Players     []string `firestore:"players"`
	Substitutes []string `firestore:"substitutes"`
	Coaches     []string `firestore:"coaches"`
	TeamName    string   `firestore:"team_name"`
	Description string   `firestore:"description"`
	Game        string   `firestore:"game"`
	Icon        string   `firestore:"icon"`
	Region      string   `firestore:"region"`
	Sex         string   `firestore:"sex"`
}

func (t Team) GetMemberType(memberID string) ([]string, bool) {
	memberType := []string{}
	hasType := false

	if memberID == t.OwnerID {
		hasType = true
		memberType = append(memberType, "Manager")
	}

	if utils.ContainsString(t.Players, memberID) {
		hasType = true
		memberType = append(memberType, "Player")
	}

	if utils.ContainsString(t.Substitutes, memberID) {
		hasType = true
		memberType = append(memberType, "Sub")
	}

	if utils.ContainsString(t.Coaches, memberID) {
		hasType = true
		memberType = append(memberType, "Coach")
	}

	return memberType, hasType
}
