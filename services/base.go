package services

import core "github.com/kode-magic/eco-bowl-api/core/entities"

type BaseService struct {
	User      core.UserRepository
	Institute core.InstitutionRepo
	Event     core.EventRepo
}
