package services

import (
	core "github.com/kode-magic/eco-bowl-api/core/entities"
)

type TraineeService struct {
	Repo core.TraineeRepo
}

func (t TraineeService) Add(record *core.Trainee) (*core.Trainee, map[string]string) {
	errMsg := make(map[string]string)
	_, phoneErr := t.Repo.GetByPhoneEmail(record.EventID, record.Phone)
	if phoneErr != nil {
		return t.Repo.Create(record)
	}

	errMsg["person"] = "person with email and(or) phone already registered for the event"
	return nil, errMsg
}

func (t TraineeService) List(event string) (*[]core.Trainee, error) {
	return t.Repo.List(event)
}

func (t TraineeService) Get(id string) (*core.Trainee, error) {
	return t.Repo.Get(id)
}
