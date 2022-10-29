package persist

import (
	"errors"
	enum "github.com/kode-magic/eco-bowl-api/core/commons"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	infra "github.com/kode-magic/eco-bowl-api/infra/entities"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
)

type eventRepo struct {
	db *gorm.DB
}

var _ core.EventRepo = &eventRepo{}

func NewEventRepo(db *gorm.DB) *eventRepo {
	return &eventRepo{db}
}

func toEventDomain(model infra.Event) *core.Event {
	rewards := make([]core.Reward, len(model.Rewards))
	for i, reward := range model.Rewards {
		rewards[i] = core.Reward{
			ID:          reward.ID.String(),
			Name:        reward.Name,
			Description: reward.Description,
			CreatedAt:   reward.CreatedAt,
			UpdatedAt:   reward.UpdatedAt,
		}
	}

	teams := make([]core.Team, len(model.Teams))
	for i, team := range model.Teams {
		teams[i] = core.Team{
			ID:          team.ID.String(),
			Name:        team.Name,
			Description: team.Description,
			CreatedAt:   team.CreatedAt,
			UpdatedAt:   team.UpdatedAt,
		}
	}

	trainees := make([]core.Trainee, len(model.Trainees))
	for i, trainee := range model.Trainees {
		trainees[i] = core.Trainee{
			ID:            trainee.ID.String(),
			Forename:      trainee.Forename,
			Surname:       trainee.Surname,
			Gender:        enum.Genders(trainee.Gender),
			Phone:         trainee.Phone,
			Email:         trainee.Email,
			Qualification: trainee.Qualification,
			BirthDate:     trainee.BirthDate,
			Team: core.Team{
				ID:   trainee.TeamID,
				Name: trainee.Team.Name,
			},
			CreatedAt: trainee.CreatedAt,
			UpdatedAt: trainee.UpdatedAt,
		}
	}

	return &core.Event{
		ID:          model.ID.String(),
		Name:        model.Name,
		Description: model.Description,
		StartDate:   model.StartDate,
		EndDate:     model.EndDate,
		Institution: core.Institution{
			ID:            model.InstitutionID,
			Name:          model.Institution.Name,
			Description:   model.Institution.Description,
			Address:       model.Institution.Address,
			ContactPerson: core.Contact(model.Institution.ContactPerson),
		},
		Rewards:   rewards,
		Teams:     teams,
		Trainees:  trainees,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func toEventPersistence(model core.Event) *infra.Event {
	return &infra.Event{
		Name:          model.Name,
		Description:   model.Description,
		StartDate:     model.StartDate,
		EndDate:       model.EndDate,
		InstitutionID: model.Institution.ID,
	}
}

func (d eventRepo) Create(event *core.Event) (*core.Event, map[string]string) {
	infraErr := map[string]string{}

	createEvent := toEventPersistence(*event)

	err := d.db.Create(&createEvent).Error

	if err != nil {
		infraErr["db_error"] = err.Error()
		return nil, infraErr
	}

	return toEventDomain(*createEvent), nil
}

func (d eventRepo) List() (*[]core.Event, error) {
	var dbEvents []infra.Event
	err := d.db.Preload("Institution").Preload("Trainees").Preload("Rewards").Preload("Teams").Find(&dbEvents).Error

	if err != nil {
		return nil, err
	}

	events := make([]core.Event, len(dbEvents))

	for i, event := range dbEvents {
		toDomain := toEventDomain(event)

		if err != nil {
			return nil, err
		}

		events[i] = *toDomain
	}

	return &events, nil

}

func (d eventRepo) Get(id string) (*core.Event, error) {
	var event infra.Event

	ID := ulids.ConvertToUUID(id)

	err := d.db.Preload("Institution").Preload("Trainees").Preload("Teams").Preload("Rewards").Where("id = ?", ID).Take(&event).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("event not found")
		} else {
			return nil, err
		}
	}

	return toEventDomain(event), nil
}

func (d eventRepo) Update(event *core.Event) (string, error) {
	var model infra.Event
	err := d.db.Model(&model).Where("id = ?", event.ID).Updates(infra.Event{
		Name:          event.Name,
		Description:   event.Description,
		StartDate:     event.StartDate,
		EndDate:       event.EndDate,
		InstitutionID: event.Institution.ID,
	}).Error

	if err != nil {
		return "", err
	}

	return "Event updated successful", nil
}

func (d eventRepo) Delete(id string) (string, error) {
	var centre infra.Event
	ID := ulids.ConvertToUUID(id)
	err := d.db.Delete(&centre, "id = ?", ID).Error
	if err != nil {
		return "", err
	}
	return "Event deleted successfully", nil
}

func (d eventRepo) GetByName(name string) (*core.Event, error) {
	var event infra.Event

	err := d.db.Where("name = ?", name).Take(&event).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("event with name " + name + " not found")
		} else {
			return nil, err
		}
	}

	return toEventDomain(event), nil
}
