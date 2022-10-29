package services

import (
	core "github.com/kode-magic/eco-bowl-api/core/entities"
)

type SolutionRewardRequest struct {
	Solution string `json:"solution"`
	Reward   string `json:"reward"`
}

type SolutionService struct {
	Repo       core.SolutionRepo
	RewardRepo core.RewardRepo
}

func (s *SolutionService) Create(solution *core.Solution) (*core.Solution, map[string]string) {
	return s.Repo.Create(solution)
}

func (s *SolutionService) List(event string) (*[]core.Solution, error) {
	return s.Repo.List(event)
}

func (s *SolutionService) Get(id string) (*core.Solution, error) {
	return s.Repo.Get(id)
}

func (s *SolutionService) Update(solution *core.Solution) (string, error) {
	_, err := s.Repo.Get(solution.ID)
	if err != nil {
		return "", err
	}
	return s.Repo.Update(solution)
}

func (s *SolutionService) AddedReward(request *SolutionRewardRequest) (string, error) {
	reward, err := s.RewardRepo.Get(request.Reward)
	if err != nil {
		return "", err
	}
	solution, solErr := s.Repo.Get(request.Solution)
	if solErr != nil {
		return "", solErr
	}
	solution.RewardID = request.Reward
	solution.Reward = *reward
	solution.Position = reward.Position

	return s.Repo.AddReward(solution)
}
