package persist

import (
	"errors"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	infra "github.com/kode-magic/eco-bowl-api/infra/entities"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
)

type solutionRepo struct {
	db *gorm.DB
}

var _ core.SolutionRepo = &solutionRepo{}

func NewSolutionRepo(db *gorm.DB) *solutionRepo {
	return &solutionRepo{db}
}

func toSolutionDomain(model infra.Solution) *core.Solution {
	return &core.Solution{
		ID:          model.ID.String(),
		Title:       model.Title,
		Description: model.Description,
		Event: core.Event{
			ID:   model.EventID,
			Name: model.Event.Name,
		},
		Team: core.Team{
			ID:          model.TeamID,
			Name:        model.Team.Name,
			Description: model.Team.Description,
		},
		Reward: core.Reward{
			ID:          model.RewardID,
			Name:        model.Reward.Name,
			Description: model.Reward.Description,
			Position:    model.Position,
		},
		Position:  model.Position,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func toSolutionPersistence(model core.Solution) *infra.Solution {
	return &infra.Solution{
		Title:       model.Title,
		Description: model.Description,
		EventID:     model.EventID,
		TeamID:      model.TeamID,
	}
}

func (d solutionRepo) Create(solution *core.Solution) (*core.Solution, map[string]string) {
	infraErr := map[string]string{}

	createSolution := toSolutionPersistence(*solution)

	err := d.db.Create(&createSolution).Error

	if err != nil {
		infraErr["db_error"] = err.Error()
		return nil, infraErr
	}

	return toSolutionDomain(*createSolution), nil
}

func (d solutionRepo) List(event string) (*[]core.Solution, error) {
	var dbSolutions []infra.Solution
	err := d.db.Preload("Team").Where("event_id = ?", event).Order("position asc").Find(&dbSolutions).Error

	if err != nil {
		return nil, err
	}

	solutions := make([]core.Solution, len(dbSolutions))

	for i, solution := range dbSolutions {
		toDomain := toSolutionDomain(solution)

		if err != nil {
			return nil, err
		}

		solutions[i] = *toDomain
	}

	return &solutions, nil

}

func (d solutionRepo) Get(id string) (*core.Solution, error) {
	var solution infra.Solution

	ID := ulids.ConvertToUUID(id)

	err := d.db.Preload("Team").Where("id = ?", ID).Take(&solution).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("solution not found")
		} else {
			return nil, err
		}

	}

	return toSolutionDomain(solution), nil
}

func (d solutionRepo) Update(solution *core.Solution) (string, error) {
	var model infra.Solution
	err := d.db.Model(&model).Where("id = ?", solution.ID).Updates(infra.Solution{
		Title:       solution.Title,
		Description: solution.Description,
		EventID:     solution.Event.ID,
		TeamID:      solution.TeamID,
	}).Error

	if err != nil {
		return "", err
	}

	return "Solution updated successful", nil
}

func (d solutionRepo) AddReward(solution *core.Solution) (string, error) {
	var model infra.Solution
	err := d.db.Model(&model).Where("id = ?", model.ID).Updates(infra.Solution{
		RewardID: solution.RewardID,
		Position: solution.Reward.Position,
	}).Error

	if err != nil {
		return "", err
	}

	return "Solution reward added successfully ", nil
}

func (d solutionRepo) Delete(id string) (string, error) {
	var reward infra.Solution
	ID := ulids.ConvertToUUID(id)
	err := d.db.Delete(&reward, "id = ?", ID).Error
	if err != nil {
		return "", err
	}
	return "Solution deleted successfully", nil
}
