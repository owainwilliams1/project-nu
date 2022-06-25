package responses

type Success string

const (
	AcceptInvite            Success = "Successfully joined team `%s`."
	CreateTeam              Success = "Team `%s` created!"
	DeleteTeam              Success = "Successfully deleted your team."
	InviteMember            Success = "Successfully invited <@%s>. They must now run `/accept-invite team-id:%s` to join."
	Leave                   Success = "Successfully left the team."
	MakeMemberCoach         Success = "Successfully made <@%s> a coach."
	MakeMemberPlayer        Success = "Successfully made <@%s> a player."
	MakeMemberSubstitute    Success = "Successfully made <@%s> a substitute."
	MakeMemberNotCoach      Success = "Successfully made <@%s> a coach."
	MakeMemberNotPlayer     Success = "Successfully made <@%s> a player."
	MakeMemberNotSubstitute Success = "Successfully made <@%s> a substitute."
	Register                Success = "Welcome, <@%s>, you have been registered!"
	RemoveMember            Success = "Successfully removed <@%s> from the team."
	SetTeamIcon             Success = "Successfully updated team icon."
	TransferOwnership       Success = "Successfully made <@%s> the new Manager."
)
