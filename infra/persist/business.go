package persist

import (
	"errors"
	enum "github.com/kode-magic/eco-bowl-api/core/commons"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	infra "github.com/kode-magic/eco-bowl-api/infra/entities"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
)

type businessRepo struct {
	db *gorm.DB
}

var _ core.BusinessRepo = &businessRepo{}

func NewBusinessRepo(db *gorm.DB) *businessRepo {
	return &businessRepo{db}
}

func toBusinessDomain(model infra.Business) *core.Business {
	growths := make([]core.Growth, len(model.Growths))
	for i, growth := range model.Growths {
		growths[i] = core.Growth{
			ID:          growth.ID.String(),
			Year:        growth.Year,
			NetWorth:    growth.NetWorth,
			Income:      growth.Income,
			Expenditure: growth.Expenditure,
			Assets:      growth.Assets,
			CreatedAt:   growth.CreatedAt,
			UpdatedAt:   growth.UpdatedAt,
		}
	}

	entrepreneurs := make([]core.Entrepreneur, len(model.Entrepreneurs))
	for i, entrepreneur := range model.Entrepreneurs {
		entrepreneurs[i] = core.Entrepreneur{
			ID:        entrepreneur.ID.String(),
			Forename:  entrepreneur.Forename,
			Surname:   entrepreneur.Surname,
			Gender:    enum.Genders(entrepreneur.Gender),
			Phone:     entrepreneur.Phone,
			Email:     entrepreneur.Email,
			BirthDate: entrepreneur.BirthDate,
			CreatedAt: entrepreneur.CreatedAt,
			UpdatedAt: entrepreneur.UpdatedAt,
		}
	}
	return &core.Business{
		ID:            model.ID.String(),
		Name:          model.Name,
		Description:   model.Description,
		Founded:       model.Founded,
		Type:          enum.BusinessTypes(model.Type),
		Level:         enum.BusinessLevels(model.Level),
		ContactPerson: core.Contact(model.ContactPerson),
		Growths:       growths,
		Entrepreneurs: entrepreneurs,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
}

func toBusinessPersistence(model core.Business) *infra.Business {
	return &infra.Business{
		Name:          model.Name,
		Description:   model.Description,
		Founded:       model.Founded,
		Type:          string(model.Type),
		Level:         string(model.Level),
		ContactPerson: infra.Contact(model.ContactPerson),
	}
}

func (d businessRepo) Create(business *core.Business) (*core.Business, map[string]string) {
	infraErr := map[string]string{}

	createCentre := toBusinessPersistence(*business)

	err := d.db.Create(&createCentre).Error

	if err != nil {
		infraErr["db_error"] = err.Error()
		return nil, infraErr
	}

	return toBusinessDomain(*createCentre), nil
}

func (d businessRepo) List() (*[]core.Business, error) {
	var dbRewards []infra.Business
	err := d.db.Preload("Entrepreneurs").Preload("Growths").Find(&dbRewards).Error

	if err != nil {
		return nil, err
	}

	businesses := make([]core.Business, len(dbRewards))

	for i, reward := range dbRewards {
		toDomain := toBusinessDomain(reward)

		if err != nil {
			return nil, err
		}

		businesses[i] = *toDomain
	}

	return &businesses, nil

}

func (d businessRepo) Get(id string) (*core.Business, error) {
	var business infra.Business

	ID := ulids.ConvertToUUID(id)

	err := d.db.Where("id = ?", ID).Take(&business).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("business not found")
		} else {
			return nil, err
		}

	}

	return toBusinessDomain(business), nil
}

func (d businessRepo) Update(reward *core.Business) (string, error) {
	var model infra.Business
	err := d.db.Model(&model).Where("id = ?", reward.ID).Updates(infra.Business{
		Name:          reward.Name,
		Description:   reward.Description,
		Founded:       reward.Founded,
		Type:          string(reward.Type),
		Level:         string(reward.Level),
		ContactPerson: infra.Contact(reward.ContactPerson),
	}).Error

	if err != nil {
		return "", err
	}

	return "Business updated successful", nil
}

func (d businessRepo) Delete(id string) (string, error) {
	var business infra.Business
	ID := ulids.ConvertToUUID(id)
	err := d.db.Delete(&business, "id = ?", ID).Error
	if err != nil {
		return "", err
	}
	return "Business deleted successfully", nil
}

func (d businessRepo) GetByName(name string) (*core.Business, error) {
	var business infra.Business

	err := d.db.Where("name = ?", name).Take(&business).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("business with name " + name + " not found")
		} else {
			return nil, err
		}
	}

	return toBusinessDomain(business), nil
}

func (d businessRepo) AddGrowthToBusiness(growth *core.Growth) (string, error) {
	createGrowth := infra.Growth{
		Year:        growth.Year,
		NetWorth:    growth.NetWorth,
		Income:      growth.Income,
		Expenditure: growth.Expenditure,
		Assets:      growth.Assets,
		BusinessID:  growth.BusinessID,
	}
	err := d.db.Create(createGrowth).Error

	if err != nil {
		return "", err
	}

	return "Growth successfully added to business", nil
}
