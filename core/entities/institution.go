package entities

import "time"

type (
	Contact struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	Institution struct {
		ID            string    `json:"id"`
		Name          string    `json:"name"`
		Description   string    `json:"description"`
		Address       string    `json:"address"`
		ContactPerson Contact   `json:"contactPerson"`
		CreatedAt     time.Time `json:"createdAt"`
		UpdatedAt     time.Time `json:"updatedAt"`
	}

	InstitutionRepo interface {
		Create(institute *Institution) (*Institution, map[string]string)
		List() (*[]Institution, error)
		Get(id string) (*Institution, error)
		GetByName(name string) (*Institution, error)
		Update(institute *Institution) (string, error)
	}
)

func (i *Institution) Validate() map[string]string {
	var errMessages = make(map[string]string)
	if i.Name == "" {
		errMessages["name"] = "centre name is required"
	}
	if i.Address == "" {
		errMessages["address"] = "address is required"
	}
	if i.ContactPerson.Name == "" {
		errMessages["contact_name"] = "Contact person name is required"
	}
	if i.ContactPerson.Phone == "" {
		errMessages["contact_phone"] = "Contact person phone number is required"
	}

	return errMessages
}
