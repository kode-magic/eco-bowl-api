package entities

import "time"

type (
	Team struct {
		ID          string    `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		EventID     string    `json:"eventId"`
		Event       Event     `json:"event"`
		Trainees    []Trainee `json:"trainees"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}

	TeamRepo interface {
		Create(team *Team) (*Team, map[string]string)
		List() (*[]Team, error)
		Get(id string) (*Team, error)
		Update(team *Team) (string, error)
	}
)

func (r *Team) Validate() map[string]string {
	var errMessages = make(map[string]string)
	if r.Name == "" {
		errMessages["name"] = "team name is required"
	}
	if r.Description == "" {
		errMessages["description"] = "description is required"
	}
	if r.EventID == "" {
		errMessages["event"] = "event is required"
	}

	return errMessages
}
