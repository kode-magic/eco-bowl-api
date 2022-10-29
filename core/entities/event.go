package entities

import "time"

type (
	Event struct {
		ID            string      `json:"id"`
		Name          string      `json:"name"`
		Description   string      `json:"description"`
		StartDate     time.Time   `json:"startDate"`
		EndDate       time.Time   `json:"endDate"`
		InstitutionID string      `json:"institutionId"`
		Institution   Institution `json:"institution"`
		Trainees      []Trainee   `json:"trainees"`
		Teams         []Team      `json:"teams"`
		Rewards       []Reward    `json:"rewards"`
		CreatedAt     time.Time   `json:"createdAt"`
		UpdatedAt     time.Time   `json:"updatedAt"`
	}

	EventRepo interface {
		Create(event *Event) (*Event, map[string]string)
		List() (*[]Event, error)
		Get(id string) (*Event, error)
		GetByName(name string) (*Event, error)
		Update(event *Event) (string, error)
	}
)

func (i *Event) Validate() map[string]string {
	var date time.Time
	var errMessages = make(map[string]string)
	if i.Name == "" {
		errMessages["name"] = "centre name is required"
	}
	if i.Institution.ID == "" {
		errMessages["institution"] = "address is required"
	}
	if i.Description == "" {
		errMessages["description"] = "description is required"
	}
	if i.StartDate == date {
		errMessages["start_date"] = "start_date is required"
	}
	if i.EndDate == date {
		errMessages["end_date"] = "end_date is required"
	}

	return errMessages
}
