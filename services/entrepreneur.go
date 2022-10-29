package services

import (
	core "github.com/kode-magic/eco-bowl-api/core/entities"
)

type EntrepreneurService struct {
	Repo core.EntrepreneurRepo
}

func (t EntrepreneurService) Add(record *core.Entrepreneur) (*core.Entrepreneur, map[string]string) {
	errMsg := make(map[string]string)
	_, phoneErr := t.Repo.GetByPhoneEmail(record.Email, record.Phone)
	if phoneErr != nil {
		return t.Repo.Create(record)
	}

	errMsg["person"] = "person with email and(or) phone already exist"
	return nil, errMsg
}

func (t EntrepreneurService) List() (*[]core.Entrepreneur, error) {
	return t.Repo.List()
}

func (t EntrepreneurService) Get(id string) (*core.Entrepreneur, error) {
	return t.Repo.Get(id)
}
