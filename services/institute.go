package services

import (
	"errors"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
)

type InstituteService struct {
	Repo core.InstitutionRepo
}

func (d *InstituteService) Create(centre *core.Institution) (*core.Institution, map[string]string) {
	_, err := d.Repo.GetByName(centre.Name)

	if err != nil {
		return d.Repo.Create(centre)
	}

	errMessage := make(map[string]string)
	errMessage["name"] = errors.New("centre with name " + centre.Name + " already exist").Error()

	return nil, errMessage
}

func (d *InstituteService) List() (*[]core.Institution, error) {
	return d.Repo.List()
}

func (d *InstituteService) Get(id string) (*core.Institution, error) {
	return d.Repo.Get(id)
}

func (d *InstituteService) Update(centre *core.Institution) (string, error) {
	_, err := d.Repo.Get(centre.ID)
	if err != nil {
		return "", err
	}
	return d.Repo.Update(centre)
}
