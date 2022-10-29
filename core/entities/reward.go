package entities

import "time"

type (
	Reward struct {
		ID          string    `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		EventID     string    `json:"eventId"`
		Event       Event     `json:"event"`
		Position    int       `json:"position"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}

	RewardRepo interface {
		Create(reward *Reward) (*Reward, map[string]string)
		List() (*[]Reward, error)
		Get(id string) (*Reward, error)
		Update(reward *Reward) (string, error)
	}
)

func (r *Reward) Validate() map[string]string {
	var errMessages = make(map[string]string)
	if r.Name == "" {
		errMessages["name"] = "name is required"
	}
	if r.Description == "" {
		errMessages["description"] = "description is required"
	}
	if r.EventID == "" {
		errMessages["event"] = "event is required"
	}
	if r.Position < 1 {
		errMessages["position"] = "position is required"
	}

	return errMessages
}
