package persist

import (
	enum "github.com/kode-magic/eco-bowl-api/core/commons"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	infra "github.com/kode-magic/eco-bowl-api/infra/entities"
	"gorm.io/gorm"
)

type entrepreneurRepo struct {
	db *gorm.DB
}

func ToEntrepreneurDomain(model infra.Entrepreneur) *core.Entrepreneur {
	return &core.Entrepreneur{
		ID:        model.ID.String(),
		Forename:  model.Forename,
		Surname:   model.Surname,
		Gender:    enum.Genders(model.Gender),
		BirthDate: model.BirthDate,
		Phone:     model.Phone,
		Email:     model.Email,
		Business: core.Business{
			ID:   model.BusinessID,
			Name: model.Business.Name,
		},
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func ToEntrepreneurPersistence(model core.Entrepreneur) *infra.Entrepreneur {
	return &infra.Entrepreneur{
		Forename:  model.Forename,
		Surname:   model.Surname,
		Gender:    string(model.Gender),
		BirthDate: model.BirthDate,
		Phone:     model.Phone,
		Email:     model.Email,
	}
}

func (t entrepreneurRepo) Create(trainee *core.Entrepreneur) (*core.Entrepreneur, map[string]string) {
	infraErr := map[string]string{}

	createEntrepreneur := ToEntrepreneurPersistence(*trainee)

	err := t.db.Create(&createEntrepreneur).Error

	if err != nil {
		infraErr["db_error"] = err.Error()
		return nil, infraErr
	}

	return ToEntrepreneurDomain(*createEntrepreneur), nil
}

func (t entrepreneurRepo) List() (*[]core.Entrepreneur, error) {
	var dbTrainees []infra.Entrepreneur

	err := t.db.Preload("Business").Find(&dbTrainees).Error

	if err != nil {
		return nil, err
	}

	entrepreneurs := make([]core.Entrepreneur, len(dbTrainees))

	for i, trainee := range dbTrainees {
		toDomain := ToEntrepreneurDomain(trainee)

		if err != nil {
			return nil, err
		}

		entrepreneurs[i] = *toDomain
	}

	return &entrepreneurs, nil
}

func (t entrepreneurRepo) Get(id string) (*core.Entrepreneur, error) {
	var entrepreneur *infra.Entrepreneur

	err := t.db.Preload("Business").Where("id = ?", id).Take(&entrepreneur).Error

	if err != nil {
		return nil, err
	}

	return ToEntrepreneurDomain(*entrepreneur), nil
}

func (t entrepreneurRepo) GetByPhoneEmail(phone, email string) (*core.Entrepreneur, error) {
	var entrepreneur *infra.Entrepreneur

	err := t.db.Where("phone = ? OR email = ?", phone, email).Take(&entrepreneur).Error

	if err != nil {
		return nil, err
	}

	return ToEntrepreneurDomain(*entrepreneur), nil
}

func (t entrepreneurRepo) AddToBusiness(person *core.Entrepreneur) (string, error) {
	var model infra.Entrepreneur
	err := t.db.Model(&model).Where("id = ?", person.ID).Updates(infra.Entrepreneur{
		BusinessID: person.BusinessID,
	}).Error

	if err != nil {
		return "", err
	}

	return "Entrepreneur successfully added to business", nil
}

func NewEntrepreneurRepo(db *gorm.DB) *entrepreneurRepo {
	return &entrepreneurRepo{db}
}

var _ core.EntrepreneurRepo = &entrepreneurRepo{}
