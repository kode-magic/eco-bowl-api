package persist

import (
	"errors"
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
	err := d.db.Model(&model).Where("id = ?", model.ID).Updates(infra.Event{
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
