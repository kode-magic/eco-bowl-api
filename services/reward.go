package services

import (
	core "github.com/kode-magic/eco-bowl-api/core/entities"
)

type RewardService struct {
	Repo core.RewardRepo
}

func (r *RewardService) Create(centre *core.Reward) (*core.Reward, map[string]string) {
	return r.Repo.Create(centre)
}

func (r *RewardService) List() (*[]core.Reward, error) {
	return r.Repo.List()
}

func (r *RewardService) Get(id string) (*core.Reward, error) {
	return r.Repo.Get(id)
}

func (r *RewardService) Update(reward *core.Reward) (string, error) {
	_, err := r.Repo.Get(reward.ID)
	if err != nil {
		return "", err
	}
	return r.Repo.Update(reward)
}
