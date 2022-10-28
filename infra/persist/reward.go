package persist

import (
	"errors"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	infra "github.com/kode-magic/eco-bowl-api/infra/entities"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
)

type rewardRepo struct {
	db *gorm.DB
}

var _ core.RewardRepo = &rewardRepo{}

func NewRewardRepo(db *gorm.DB) *rewardRepo {
	return &rewardRepo{db}
}

func toRewardDomain(model infra.Reward) *core.Reward {
	return &core.Reward{
		ID:          model.ID.String(),
		Name:        model.Name,
		Description: model.Description,
		Offers:      model.Offers,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func toRewardPersistence(model core.Reward) *infra.Reward {
	return &infra.Reward{
		Name:        model.Name,
		Description: model.Description,
		Offers:      model.Offers,
		EventID:     model.Event.ID,
	}
}

func (d rewardRepo) Create(reward *core.Reward) (*core.Reward, map[string]string) {
	infraErr := map[string]string{}

	createCentre := toRewardPersistence(*reward)

	err := d.db.Create(&createCentre).Error

	if err != nil {
		infraErr["db_error"] = err.Error()
		return nil, infraErr
	}

	return toRewardDomain(*createCentre), nil
}

func (d rewardRepo) List() (*[]core.Reward, error) {
	var dbRewards []infra.Reward
	err := d.db.Find(&dbRewards).Error

	if err != nil {
		return nil, err
	}

	rewards := make([]core.Reward, len(dbRewards))

	for i, reward := range dbRewards {
		toDomain := toRewardDomain(reward)

		if err != nil {
			return nil, err
		}

		rewards[i] = *toDomain
	}

	return &rewards, nil

}

func (d rewardRepo) Get(id string) (*core.Reward, error) {
	var reward infra.Reward

	ID := ulids.ConvertToUUID(id)

	err := d.db.Where("id = ?", ID).Take(&reward).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("reward not found")
		} else {
			return nil, err
		}

	}

	return toRewardDomain(reward), nil
}

func (d rewardRepo) Update(reward *core.Reward) (string, error) {
	var model infra.Reward
	err := d.db.Model(&model).Where("id = ?", model.ID).Updates(infra.Reward{
		Name:        reward.Name,
		Description: reward.Description,
		Offers:      reward.Offers,
		EventID:     reward.Event.ID,
	}).Error

	if err != nil {
		return "", err
	}

	return "Reward updated successful", nil
}

func (d rewardRepo) Delete(id string) (string, error) {
	var reward infra.Reward
	ID := ulids.ConvertToUUID(id)
	err := d.db.Delete(&reward, "id = ?", ID).Error
	if err != nil {
		return "", err
	}
	return "Reward deleted successfully", nil
}
