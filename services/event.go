package services

import (
	"errors"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
)

type EventService struct {
	Repo core.EventRepo
}

func (d *EventService) Create(event *core.Event) (*core.Event, map[string]string) {
	_, err := d.Repo.GetByName(event.Name)

	if err != nil {
		return d.Repo.Create(event)
	}

	errMessage := make(map[string]string)
	errMessage["name"] = errors.New("event with name " + event.Name + " already exist").Error()

	return nil, errMessage
}

func (d *EventService) List() (*[]core.Event, error) {
	return d.Repo.List()
}

func (d *EventService) Get(id string) (*core.Event, error) {
	return d.Repo.Get(id)
}

func (d *EventService) Update(event *core.Event) (string, error) {
	_, err := d.Repo.Get(event.ID)
	if err != nil {
		return "", err
	}
	return d.Repo.Update(event)
}
