package entities

import "time"

type (
	Solution struct {
		ID          string    `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		EventID     string    `json:"eventId"`
		Event       Event     `json:"event"`
		TeamID      string    `json:"teamId"`
		Team        Team      `json:"team"`
		RewardID    string    `json:"RewardId"`
		Reward      Reward    `json:"reward"`
		Position    int       `json:"position"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}

	SolutionRepo interface {
		Create(solution *Solution) (*Solution, map[string]string)
		List(event string) (*[]Solution, error)
		Get(id string) (*Solution, error)
		Update(solution *Solution) (string, error)
		AddReward(solution *Solution) (string, error)
	}
)

func (r *Solution) Validate() map[string]string {
	var errMessages = make(map[string]string)
	if r.Title == "" {
		errMessages["name"] = "name is required"
	}
	if r.Description == "" {
		errMessages["description"] = "description is required"
	}
	if r.EventID == "" {
		errMessages["event"] = "event is required"
	}
	if r.TeamID == "" {
		errMessages["team"] = "team is required"
	}

	return errMessages
}
