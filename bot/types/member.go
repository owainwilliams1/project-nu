package types

type Member struct {
	MemberID         string `firestore:"-"`
	Team             string `firestore:"team"`
	Username         string `firestore:"username"`
	ValorantUsername string `firestore:"valorant_username,omitempty"`
	ApexUsername     string `firestore:"apex_username,omitempty"`
	Icon             string `firestore:"icon,omitempty"`
}
