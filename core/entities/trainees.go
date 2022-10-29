package entities

import (
	"github.com/kode-magic/eco-bowl-api/core/commons"
	"time"
)

type (
	Trainee struct {
		ID            string          `json:"id"`
		Forename      string          `json:"forename"`
		Surname       string          `json:"surname"`
		Gender        commons.Genders `json:"gender"`
		BirthDate     time.Time       `json:"birthDate"`
		Phone         string          `json:"phone"`
		Email         string          `json:"email"`
		Qualification string          `json:"qualification"`
		EventID       string          `json:"eventId"`
		Event         Event           `json:"event"`
		TeamID        string          `json:"teamId"`
		Team          Team            `json:"team"`
		CreatedAt     time.Time       `json:"createdAt"`
		UpdatedAt     time.Time       `json:"updatedAt"`
	}

	TraineeRepo interface {
		Create(trainee *Trainee) (*Trainee, map[string]string)
		List(event string) (*[]Trainee, error)
		Get(id string) (*Trainee, error)
		GetByPhoneEmail(event, phoneEmail string) (*Trainee, error)
		GetByEmail(event, email string) (*Trainee, error)
		AddToTeam(trainee *Trainee) (*Trainee, error)
	}
)

func (trainee *Trainee) Validate() map[string]string {
	var date time.Time
	var errorMessages = make(map[string]string)
	if trainee.EventID == "" {
		errorMessages["event"] = "event is required"
	}
	if trainee.Forename == "" {
		errorMessages["forename"] = "forename is required"
	}
	if trainee.Surname == "" {
		errorMessages["surname"] = "surname is required"
	}
	if trainee.Phone == "" {
		errorMessages["phone"] = "phone is required"
	}
	if trainee.Email == "" {
		errorMessages["email"] = "email is required"
	}
	if trainee.Gender != commons.Male && trainee.Gender != commons.Female {
		errorMessages["gender"] = "gender is required and must be a valid gender value"
	}
	if trainee.BirthDate == date {
		errorMessages["birth_date"] = "birth date is required"
	}
	if trainee.Qualification == "" {
		errorMessages["qualification"] = "qualification is required"
	}

	return errorMessages
}
