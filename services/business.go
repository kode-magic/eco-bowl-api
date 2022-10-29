package services

import core "github.com/kode-magic/eco-bowl-api/core/entities"

type BusinessService struct {
	Repo core.BusinessRepo
}

func (s BusinessService) Create(business *core.Business) (*core.Business, map[string]string) {
	//_, err := s.Repo.GetByName(business.Name)
	//if err != nil {
	//	return s.Repo.Create(business)
	//}

	return s.Repo.Create(business)

}

func (s BusinessService) List() {

}
