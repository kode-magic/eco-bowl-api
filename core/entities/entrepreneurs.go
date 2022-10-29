package entities

import (
	"github.com/kode-magic/eco-bowl-api/core/commons"
	"time"
)

type (
	Entrepreneur struct {
		ID         string          `json:"id"`
		Forename   string          `json:"forename"`
		Surname    string          `json:"surname"`
		Gender     commons.Genders `json:"gender"`
		BirthDate  time.Time       `json:"birthDate"`
		Phone      string          `json:"phone"`
		Email      string          `json:"email"`
		BusinessID string          `json:"businessId"`
		Business   Business        `json:"business"`
		CreatedAt  time.Time       `json:"createdAt"`
		UpdatedAt  time.Time       `json:"updatedAt"`
	}

	EntrepreneurRepo interface {
		Create(person *Entrepreneur) (*Entrepreneur, map[string]string)
		List() (*[]Entrepreneur, error)
		Get(id string) (*Entrepreneur, error)
		GetByPhoneEmail(phone, email string) (*Entrepreneur, error)
		AddToBusiness(person *Entrepreneur) (string, error)
	}
)

func (e *Entrepreneur) Validate() map[string]string {
	var date time.Time
	var errorMessages = make(map[string]string)
	if e.Forename == "" {
		errorMessages["forename"] = "forename is required"
	}
	if e.Surname == "" {
		errorMessages["surname"] = "surname is required"
	}
	if e.Phone == "" {
		errorMessages["phone"] = "phone is required"
	}
	if e.Email == "" {
		errorMessages["email"] = "email is required"
	}
	if e.Gender != commons.Male && e.Gender != commons.Female {
		errorMessages["gender"] = "gender is required and must be a valid gender value"
	}
	if e.BirthDate == date {
		errorMessages["birth_date"] = "birth date is required"
	}

	return errorMessages
}
