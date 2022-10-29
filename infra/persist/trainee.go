package persist

import (
	enum "github.com/kode-magic/eco-bowl-api/core/commons"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	infra "github.com/kode-magic/eco-bowl-api/infra/entities"
	"gorm.io/gorm"
)

type traineeRepo struct {
	db *gorm.DB
}

func ToTraineeDomain(model infra.Trainee) *core.Trainee {

	return &core.Trainee{
		ID:        model.ID.String(),
		Forename:  model.Forename,
		Surname:   model.Surname,
		Gender:    enum.Genders(model.Gender),
		BirthDate: model.BirthDate,
		Phone:     model.Phone,
		Email:     model.Email,
		Team: core.Team{
			ID:          model.Team.ID.String(),
			Name:        model.Team.Name,
			Description: model.Team.Description,
		},
		Qualification: model.Qualification,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
}

func ToTraineePersistence(model core.Trainee) *infra.Trainee {
	return &infra.Trainee{
		Forename:      model.Forename,
		Surname:       model.Surname,
		Gender:        string(model.Gender),
		BirthDate:     model.BirthDate,
		Phone:         model.Phone,
		Email:         model.Email,
		EventID:       model.EventID,
		Qualification: model.Qualification,
	}
}

func (t traineeRepo) Create(trainee *core.Trainee) (*core.Trainee, map[string]string) {
	infraErr := map[string]string{}

	createTrainee := ToTraineePersistence(*trainee)

	err := t.db.Create(&createTrainee).Error

	if err != nil {
		infraErr["db_error"] = err.Error()
		return nil, infraErr
	}

	return ToTraineeDomain(*createTrainee), nil
}

func (t traineeRepo) List(event string) (*[]core.Trainee, error) {
	var dbTrainees []infra.Trainee

	err := t.db.Preload("Team").Where("event_id = ?", event).Find(&dbTrainees).Error

	if err != nil {
		return nil, err
	}

	trainees := make([]core.Trainee, len(dbTrainees))

	for i, trainee := range dbTrainees {
		toDomain := ToTraineeDomain(trainee)

		if err != nil {
			return nil, err
		}

		trainees[i] = *toDomain
	}

	return &trainees, nil
}

func (t traineeRepo) Get(id string) (*core.Trainee, error) {
	var trainee *infra.Trainee

	err := t.db.Preload("Team").Where("id = ?", id).Take(&trainee).Error

	if err != nil {
		return nil, err
	}

	return ToTraineeDomain(*trainee), nil
}

func (t traineeRepo) GetByPhoneEmail(event, phoneEmail string) (*core.Trainee, error) {
	var trainee *infra.Trainee

	err := t.db.Where("event = ? AND phone = ? OR email = ?", event, phoneEmail, phoneEmail).Take(&trainee).Error

	if err != nil {
		return nil, err
	}

	return ToTraineeDomain(*trainee), nil
}

func (t traineeRepo) GetByEmail(event, email string) (*core.Trainee, error) {
	var trainee *infra.Trainee

	err := t.db.Where("event_id = ? AND email = ?", event, email).Take(&trainee).Error

	if err != nil {
		return nil, err
	}

	return ToTraineeDomain(*trainee), nil
}

func (t traineeRepo) AddToTeam(trainee *core.Trainee) (*core.Trainee, error) {
	var model infra.Trainee
	err := t.db.Model(&model).Where("id = ?", model.ID).Updates(infra.Trainee{
		TeamID: trainee.TeamID,
	}).Error

	if err != nil {
		return nil, err
	}

	return trainee, nil
}

func NewTraineeRepo(db *gorm.DB) *traineeRepo {
	return &traineeRepo{db}
}

var _ core.TraineeRepo = &traineeRepo{}
