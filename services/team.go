package services

import (
	"errors"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
)

type TeamService struct {
	Team        core.Team
	Trainee     core.Trainee
	Repo        core.TeamRepo
	TraineeRepo core.TraineeRepo
}

func (t TeamService) Create(team *core.Team) (*core.Team, map[string]string) {
	_, err := t.Repo.GetByName(team.EventID, team.Name)

	if err != nil {
		_, teamErr := t.Repo.Create(team)
		if teamErr != nil {
			return nil, teamErr
		}

		return t.Repo.Create(team)
	}

	errMessage := make(map[string]string)
	errMessage["name"] = errors.New("team with name " + team.Name + " already exist").Error()

	return nil, errMessage
}
