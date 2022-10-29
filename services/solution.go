package services

import (
	core "github.com/kode-magic/eco-bowl-api/core/entities"
)

type SolutionService struct {
	Repo core.SolutionRepo
}

func (r *SolutionService) Create(solution *core.Solution) (*core.Solution, map[string]string) {
	return r.Repo.Create(solution)
}

func (r *SolutionService) List(event string) (*[]core.Solution, error) {
	return r.Repo.List(event)
}

func (r *SolutionService) Get(id string) (*core.Solution, error) {
	return r.Repo.Get(id)
}

func (r *SolutionService) Update(solution *core.Solution) (string, error) {
	_, err := r.Repo.Get(solution.ID)
	if err != nil {
		return "", err
	}
	return r.Repo.Update(solution)
}
