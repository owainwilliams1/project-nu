package responses

type Failure string

const (
	RequireRegistration           Failure = "You need to use `/register` first."
	ForbiddenAlreadyMember        Failure = "You are already in a team."
	ForbiddenNoInvite             Failure = "You are not invited to team `%s`."
	ForbiddenTeamExists           Failure = "The team `%s` already exists."
	ForbiddenNotOwner             Failure = "You do not manage a team."
	ForbiddenAlreadyInvited       Failure = "User <@%s> has already been invited to the team."
	ForbiddenNotMember            Failure = "You are not in a team."
	ForbiddenOwnerAction          Failure = "Please transfer ownership first with `/transfer-owner` or delete your team with `/delete-team`."
	ForbiddenUserNotMember        Failure = "Member is not in your team."
	ForbiddenMaxCoaches           Failure = "You have the maximum number of `1` coaches."
	ForbiddenMaxPlayers           Failure = "You have the maximum number of `5` players."
	ForbiddenMaxSubstitutes       Failure = "You have the maximum number of `3` substitutes."
	ForbiddenAlreadyCoach         Failure = "Member is already a coach."
	ForbiddenAlreadyPlayer        Failure = "Member is already a player."
	ForbiddenAlreadySubstitute    Failure = "Member is already a substitute."
	ForbiddenAlreadyNotCoach      Failure = "Member is already not a coach."
	ForbiddenAlreadyNotPlayer     Failure = "Member is already not a player."
	ForbiddenAlreadyNotSubstitute Failure = "Member is already not a substitute."
	ForbiddenAlreadyRegistered    Failure = "You have already registered."
	NotFoundTeam                  Failure = "Team `%s` does not exist."
	NotFoundTeams                 Failure = "There are currently no teams."
	Unexpected                    Failure = "An unexpected error has occured, please try again later. If the issue persists, please contact a developer."
	ValidationTeamName            Failure = "Team names must only contain letters and spaces. They cannot be longer than 24 characters."
	ValidationHex                 Failure = "`%s` is not a valid hex code. A hex code looks like `#00ff66`."
	ValidationSex                 Failure = "`%s` is not a valid sex. Choose from `male/female/mixed`."
	ValidationRegion              Failure = "`%s` is not a valid region. Choose from `eu/na`."
	ValidationURL                 Failure = "That is not a valid URL."
)
