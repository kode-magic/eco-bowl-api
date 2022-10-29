package persist

import (
	"errors"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	infra "github.com/kode-magic/eco-bowl-api/infra/entities"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
)

type teamRepo struct {
	db *gorm.DB
}

var _ core.TeamRepo = &teamRepo{}

func NewTeamRepo(db *gorm.DB) *teamRepo {
	return &teamRepo{db}
}

func toTeamDomain(model infra.Team) *core.Team {
	return &core.Team{
		ID:          model.ID.String(),
		Name:        model.Name,
		Description: model.Description,
		Event: core.Event{
			ID:          model.EventID,
			Name:        model.Event.Name,
			Description: model.Event.Name,
			StartDate:   model.Event.StartDate,
			EndDate:     model.Event.EndDate,
		},
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func toTeamPersistence(model core.Team) *infra.Team {
	return &infra.Team{
		Name:        model.Name,
		Description: model.Description,
		EventID:     model.EventID,
	}
}

func (d teamRepo) Create(team *core.Team) (*core.Team, map[string]string) {
	infraErr := map[string]string{}

	createCentre := toTeamPersistence(*team)

	err := d.db.Create(&createCentre).Error

	if err != nil {
		infraErr["db_error"] = err.Error()
		return nil, infraErr
	}

	return toTeamDomain(*createCentre), nil
}

func (d teamRepo) List() (*[]core.Team, error) {
	var dbTeams []infra.Team
	err := d.db.Find(&dbTeams).Error

	if err != nil {
		return nil, err
	}

	teams := make([]core.Team, len(dbTeams))

	for i, team := range dbTeams {
		toDomain := toTeamDomain(team)

		if err != nil {
			return nil, err
		}

		teams[i] = *toDomain
	}

	return &teams, nil

}

func (d teamRepo) Get(id string) (*core.Team, error) {
	var team infra.Team

	ID := ulids.ConvertToUUID(id)

	err := d.db.Where("id = ?", ID).Take(&team).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("team not found")
		} else {
			return nil, err
		}
	}

	return toTeamDomain(team), nil
}

func (d teamRepo) Update(centre *core.Team) (string, error) {
	var model infra.Team
	err := d.db.Model(&model).Where("id = ?", centre.ID).Updates(infra.Team{
		Name:        centre.Name,
		Description: centre.Description,
	}).Error

	if err != nil {
		return "", err
	}

	return "Team updated successful", nil
}

func (d teamRepo) Delete(id string) (string, error) {
	var centre infra.Team
	ID := ulids.ConvertToUUID(id)
	err := d.db.Delete(&centre, "id = ?", ID).Error
	if err != nil {
		return "", err
	}
	return "Team deleted successfully", nil
}

func (d teamRepo) GetByName(name string) (*core.Team, error) {
	var centre infra.Team

	err := d.db.Where("name = ?", name).Take(&centre).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("institution with name " + name + " not found")
		} else {
			return nil, err
		}
	}

	return toTeamDomain(centre), nil
}
