package entities

import (
	"github.com/kode-magic/eco-bowl-api/core/commons"
	"github.com/kode-magic/eco-bowl-api/utils"
	"time"
)

type (
	User struct {
		ID           string        `json:"id"`
		FirstName    string        `json:"firstName"`
		LastName     string        `json:"lastName"`
		Phone        string        `json:"phone"`
		Password     string        `json:"password"`
		BasePassword string        `json:"basePassword"`
		Status       string        `json:"status"`
		Role         commons.Roles `json:"role"`
		CreatedAt    time.Time     `json:"createdAt"`
		UpdatedAt    time.Time     `json:"updatedAt"`
	}

	PublicUser struct {
		ID        string        `json:"id"`
		FirstName string        `json:"firstName" binding:"required"`
		LastName  string        `json:"lastName" binding:"required"`
		Phone     string        `json:"phone" binding:"required"`
		Role      commons.Roles `json:"role"`
		Status    string        `json:"status"`
	}

	UserWithPassword struct {
		ID           string `json:"id"`
		BasePassword string `json:"basePassword"`
		Status       string `json:"status"`
	}

	UserRepository interface {
		Add(record *User) (*User, map[string]string)
		Login(phone string) (*User, error)
		Edit(record *User) (*User, map[string]string)
		ResetPassword(phone string) (*User, error)
		ChangePassword(user User) (*User, error)
		CreatePassword(user User) (*User, error)
		Disable(id string) (*User, error)
		Enable(id string) (*User, error)
		Users() (*[]User, error)
		User(id string) (*User, error)
		Remove(id string) (string, error)
	}
)

func (user *User) RetrievePassword() *UserWithPassword {
	return &UserWithPassword{
		ID:           user.ID,
		BasePassword: user.BasePassword,
		Status:       user.Status,
	}
}

func (user *User) PublicUser() *PublicUser {
	return &PublicUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Role:      user.Role,
		Status:    user.Status,
	}
}

func (user *User) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)

	switch action {
	case "update":
		if user.ID == "" {
			errorMessages["id"] = "id is required"
		}
	case "login":
		if user.Phone == "" {
			errorMessages["phone"] = "phone is required"
		}
		if user.Password == "" {
			errorMessages["password"] = "password is required"
		}
	default:
		if user.FirstName == "" {
			errorMessages["first_name"] = "first name is required"
		}

		if user.LastName == "" {
			errorMessages["last_name"] = "last name is required"
		}

		if user.Role == "" {
			errorMessages["role"] = "role is required"
		}

		if user.Role != commons.Admin && user.Role != commons.Normal {
			errorMessages["role"] = "role must either be 'Admin' or 'Normal'"
		}

		if user.Phone == "" {
			errorMessages["phone"] = "phone is required"
		} else {
			if !utils.IsPhone(user.Phone) {
				errorMessages["phone"] = "please enter a valid phone"
			}
		}
	}

	return errorMessages
}

func (user *User) HashPassword() {
	password, _ := utils.HashPassword(user.Password)
	user.Password = password
}

func (user *User) CheckPassword(password string) bool {
	return utils.CheckPasswordHash(user.Password, password)
}
