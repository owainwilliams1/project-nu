package types

type Member struct {
	MemberID string `firestore:"-"`
	Team     string `firestore:"team"`
	Username string `firestore:"username,omitempty"`
	Icon     string `firestore:"icon,omitempty"`
}
