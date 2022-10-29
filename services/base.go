package services

import core "github.com/kode-magic/eco-bowl-api/core/entities"

type BaseService struct {
	User         core.UserRepository
	Institute    core.InstitutionRepo
	Event        core.EventRepo
	Reward       core.RewardRepo
	Team         core.TeamRepo
	Trainee      core.TraineeRepo
	Solution     core.SolutionRepo
	Entrepreneur core.EntrepreneurRepo
}
