package services

import (
	"errors"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
)

type TeamRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Event       string   `json:"event"`
	Members     []string `json:"members" binding:"required"`
}

type TeamService struct {
	Repo        core.TeamRepo
	TraineeRepo core.TraineeRepo
}

func (t TeamService) Create(param *TeamRequest) (*core.Team, map[string]string) {
	_, err := t.Repo.GetByName(param.Event, param.Name)

	if err != nil {
		teamCreate := core.Team{
			Name:        param.Name,
			Description: param.Description,
			EventID:     param.Event,
			Event: core.Event{
				ID: param.Event,
			},
		}
		team, teamErr := t.Repo.Create(&teamCreate)
		if teamErr != nil {
			return nil, teamErr
		}

		members := make([]core.Trainee, len(param.Members))

		for i, member := range param.Members {
			addTrainees := core.Trainee{
				ID:     member,
				TeamID: team.ID,
			}

			trainee, _ := t.TraineeRepo.AddToTeam(&addTrainees)
			members[i] = *trainee
		}

		team.Trainees = members

		return team, nil
	}

	errMessage := make(map[string]string)
	errMessage["name"] = errors.New("team with name " + param.Name + " already exist").Error()

	return nil, errMessage
}

func (t TeamService) List(event string) (*[]core.Team, error) {
	return t.Repo.List(event)
}
