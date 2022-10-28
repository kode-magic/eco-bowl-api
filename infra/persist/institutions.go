package persist

import (
	"errors"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	infra "github.com/kode-magic/eco-bowl-api/infra/entities"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
)

type instituteRepo struct {
	db *gorm.DB
}

var _ core.InstitutionRepo = &instituteRepo{}

func NewInstituteRepo(db *gorm.DB) *instituteRepo {
	return &instituteRepo{db}
}

func toCentreDomain(model infra.Institution) *core.Institution {
	return &core.Institution{
		ID:            model.ID.String(),
		Name:          model.Name,
		Description:   model.Description,
		Address:       model.Address,
		ContactPerson: core.Contact(model.ContactPerson),
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
}

func toCentrePersistence(model core.Institution) *infra.Institution {
	return &infra.Institution{
		Name:          model.Name,
		Description:   model.Description,
		ContactPerson: infra.Contact(model.ContactPerson),
		Address:       model.Address,
	}
}

func (d instituteRepo) Create(department *core.Institution) (*core.Institution, map[string]string) {
	infraErr := map[string]string{}

	createCentre := toCentrePersistence(*department)

	err := d.db.Create(&createCentre).Error

	if err != nil {
		infraErr["db_error"] = err.Error()
		return nil, infraErr
	}

	return toCentreDomain(*createCentre), nil
}

func (d instituteRepo) List() (*[]core.Institution, error) {
	var dbCentres []infra.Institution
	err := d.db.Find(&dbCentres).Error

	if err != nil {
		return nil, err
	}

	centres := make([]core.Institution, len(dbCentres))

	for i, centre := range dbCentres {
		toDomain := toCentreDomain(centre)

		if err != nil {
			return nil, err
		}

		centres[i] = *toDomain
	}

	return &centres, nil

}

func (d instituteRepo) Get(id string) (*core.Institution, error) {
	var centre infra.Institution

	ID := ulids.ConvertToUUID(id)

	err := d.db.Where("id = ?", ID).Take(&centre).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("centre not found")
		} else {
			return nil, err
		}

	}

	return toCentreDomain(centre), nil
}

func (d instituteRepo) Update(centre *core.Institution) (string, error) {
	var model infra.Institution
	err := d.db.Model(&model).Where("id = ?", model.ID).Updates(infra.Institution{
		Name:          centre.Name,
		Description:   centre.Description,
		Address:       centre.Address,
		ContactPerson: infra.Contact(centre.ContactPerson),
	}).Error

	if err != nil {
		return "", err
	}

	return "Institution updated successful", nil
}

func (d instituteRepo) Delete(id string) (string, error) {
	var centre infra.Institution
	ID := ulids.ConvertToUUID(id)
	err := d.db.Delete(&centre, "id = ?", ID).Error
	if err != nil {
		return "", err
	}
	return "Institution deleted successfully", nil
}

func (d instituteRepo) GetByName(name string) (*core.Institution, error) {
	var centre infra.Institution

	err := d.db.Where("name = ?", name).Take(&centre).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("institution with name " + name + " not found")
		} else {
			return nil, err
		}
	}

	return toCentreDomain(centre), nil
}
