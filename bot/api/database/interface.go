package database

import (
	"errors"

	"cloud.google.com/go/firestore"
	"hushclan.com/pkg/utils"
	"hushclan.com/types"
)

func (d *Database) CreateTeam(team *types.Team) (err error) {
	_, err = d.client.
		Collection("teams").
		Doc(team.TeamID).
		Create(d.ctx, team)

	if err != nil {
		return
	}

	err = d.AddTeamMember(team.TeamID, team.OwnerID)
	err = d.AddMemberTeam(team.OwnerID, team.TeamID)

	return
}

func (d *Database) DeleteTeam(teamName string) (err error) {
	team, err := d.GetTeam(teamName)

	for _, member := range team.Members {
		err = d.RemoveMemberTeam(member)
		return
	}

	_, err = d.client.
		Collection("teams").
		Doc(teamName).
		Delete(d.ctx)

	return
}

func (d *Database) CreateMember(member *types.Member) (err error) {
	_, err = d.client.
		Collection("members").
		Doc(member.MemberID).
		Create(d.ctx, member)

	return
}

func (d *Database) AddPlayerType(teamName, memberID string, memberType MemberType) (err error) {
	team, err := d.GetTeam(teamName)
	if err != nil {
		return
	}

	teamDoc := d.client.
		Collection("teams").
		Doc(teamName)

	switch memberType {
	case Player:
		_, err = teamDoc.Update(d.ctx, []firestore.Update{
			{
				Path:  "players",
				Value: append(team.Players, memberID),
			},
		})
	case Substitute:
		_, err = teamDoc.Update(d.ctx, []firestore.Update{
			{
				Path:  "substitutes",
				Value: append(team.Players, memberID),
			},
		})
	case Coach:
		_, err = teamDoc.Update(d.ctx, []firestore.Update{
			{
				Path:  "coaches",
				Value: append(team.Players, memberID),
			},
		})
	}

	return
}

func (d *Database) RemovePlayerType(teamName, memberID string, memberType MemberType) (err error) {
	team, err := d.GetTeam(teamName)
	if err != nil {
		return
	}

	teamDoc := d.client.
		Collection("teams").
		Doc(teamName)

	switch memberType {
	case Player:
		_, err = teamDoc.Update(d.ctx, []firestore.Update{
			{
				Path:  "players",
				Value: utils.RemoveArrayString(team.Players, memberID),
			},
		})
	case Substitute:
		_, err = teamDoc.Update(d.ctx, []firestore.Update{
			{
				Path:  "substitutes",
				Value: utils.RemoveArrayString(team.Substitutes, memberID),
			},
		})
	case Coach:
		_, err = teamDoc.Update(d.ctx, []firestore.Update{
			{
				Path:  "coaches",
				Value: utils.RemoveArrayString(team.Coaches, memberID),
			},
		})
	}

	return
}

func (d *Database) AddTeamMember(teamName string, memberID string) (err error) {
	teamDoc, err := d.client.
		Collection("teams").
		Doc(teamName).
		Get(d.ctx)

	if err != nil {
		return
	}

	team := types.Team{}
	err = teamDoc.DataTo(&team)

	if err != nil {
		return
	}

	_, err = d.client.
		Collection("teams").
		Doc(teamName).
		Update(d.ctx, []firestore.Update{
			{
				Path:  "members",
				Value: append(team.Members, memberID),
			},
		})

	return
}

func (d *Database) RemoveTeamMember(teamName string, memberID string) (err error) {
	teamDoc, err := d.client.
		Collection("teams").
		Doc(teamName).
		Get(d.ctx)

	if err != nil {
		return
	}

	team := types.Team{}
	err = teamDoc.DataTo(&team)

	if err != nil {
		return
	}

	_, err = d.client.
		Collection("teams").
		Doc(teamName).
		Update(d.ctx, []firestore.Update{
			{
				Path:  "members",
				Value: utils.RemoveArrayString(team.Members, memberID),
			},
			{
				Path:  "players",
				Value: utils.RemoveArrayString(team.Players, memberID),
			},
			{
				Path:  "substitutes",
				Value: utils.RemoveArrayString(team.Substitutes, memberID),
			},
			{
				Path:  "coaches",
				Value: utils.RemoveArrayString(team.Coaches, memberID),
			},
		})

	if err != nil {
		return
	}

	err = d.RemoveMemberTeam(memberID)

	return
}

func (d *Database) AddMemberTeam(memberID string, teamID string) (err error) {
	memberDoc, err := d.client.
		Collection("members").
		Doc(memberID).
		Get(d.ctx)

	if err != nil {
		return
	}

	member := types.Member{}
	err = memberDoc.DataTo(&member)

	if err != nil {
		return
	}

	_, err = d.client.
		Collection("members").
		Doc(memberID).
		Update(d.ctx, []firestore.Update{
			{
				Path:  "team",
				Value: teamID,
			},
		})

	return
}

func (d *Database) RemoveMemberTeam(memberID string) (err error) {
	memberDoc, err := d.client.
		Collection("teams").
		Doc(memberID).
		Get(d.ctx)

	if err != nil {
		return
	}

	member := types.Member{}
	err = memberDoc.DataTo(&member)

	if err != nil {
		return
	}

	_, err = d.client.
		Collection("members").
		Doc(memberID).
		Update(d.ctx, []firestore.Update{
			{
				Path:  "team",
				Value: "",
			},
		})

	return
}

func (d *Database) GetTeam(teamName string) (team types.Team, err error) {
	teamDoc, err := d.client.
		Collection("teams").
		Doc(teamName).
		Get(d.ctx)

	if err != nil {
		return
	}

	team = types.Team{}
	err = teamDoc.DataTo(&team)
	team.TeamID = teamDoc.Ref.ID

	return
}

func (d *Database) GetTeamByOwner(ownerID string) (team types.Team, err error) {
	teamDocs, err := d.client.
		Collection("teams").
		Where("owner_id", "==", ownerID).
		Documents(d.ctx).
		GetAll()

	if err != nil {
		return
	}

	if teamDocs == nil {
		err = errors.New("no documents were found")
		return
	}

	teamDoc := teamDocs[0]

	team = types.Team{}
	err = teamDoc.DataTo(&team)
	team.TeamID = teamDoc.Ref.ID

	return
}

func (d *Database) GetTeams(limit int, page int) (teams []types.Team, err error) {
	teamDocs, err := d.client.
		Collection("teams").
		Documents(d.ctx).
		GetAll()

	if err != nil {
		return
	}

	if teamDocs == nil {
		err = errors.New("no documents were found")
		return
	}

	skip := page * limit
	if skip > len(teamDocs) {
		skip = len(teamDocs)
	}

	end := skip + limit
	if end > len(teamDocs) {
		end = len(teamDocs)
	}

	teamDocs = teamDocs[skip:end]

	for _, doc := range teamDocs {
		team := types.Team{}
		err = doc.DataTo(&team)
		team.TeamID = doc.Ref.ID

		if err != nil {
			return
		}

		teams = append(teams, team)
	}

	return
}

func (d *Database) GetMember(memberID string) (member types.Member, err error) {
	memberDoc, err := d.client.
		Collection("members").
		Doc(memberID).
		Get(d.ctx)

	if err != nil {
		return
	}

	member = types.Member{}
	err = memberDoc.DataTo(&member)
	member.MemberID = memberDoc.Ref.ID

	return
}
