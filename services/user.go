package services

import (
	"errors"
	enum "github.com/kode-magic/eco-bowl-api/core/commons"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	"github.com/kode-magic/eco-bowl-api/utils"
)

type User struct {
	Repo core.UserRepository
}

func (u User) Add(record *core.User) (*core.User, map[string]string) {
	user, err := u.Repo.Add(record)
	if err != nil {
		return nil, err
	}

	user.BasePassword = utils.DecodeString(user.BasePassword)

	return user, nil
}

func (u User) Edit(record *core.User) (*core.User, map[string]string) {
	return u.Repo.Edit(record)
}

func (u User) Users() (*[]core.PublicUser, error) {
	users, err := u.Repo.Users()

	if err != nil {
		return nil, err
	}

	var serialAccount []core.PublicUser
	for _, account := range *users {
		serialAccount = append(serialAccount, *account.PublicUser())
	}

	return &serialAccount, nil

}

func (u User) User(id string) (*core.PublicUser, error) {
	user, err := u.Repo.User(id)

	if err != nil {
		return nil, err
	}

	return user.PublicUser(), nil
}

func (u User) Remove(id string) (string, error) {
	return u.Repo.Remove(id)
}

func (u User) ResetPassword(phone string) (*core.User, error) {
	return u.Repo.ResetPassword(phone)
}

func (u User) Disable(id string) (*core.User, error) {
	return u.Repo.Disable(id)
}

func (u User) Enable(id string) (*core.User, error) {
	return u.Repo.Enable(id)
}

func (u User) Login(phone, password string) (*core.User, error) {
	userData, err := u.Repo.Login(phone)

	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, userData.Password) {
		return nil, errors.New("invalid login credentials")
	}

	return userData, nil
}

func (u User) ChangePassword(user core.User) (*core.User, error) {
	userData, err := u.Repo.User(user.ID)
	if err != nil {
		return nil, err
	}

	if userData.Status == string(enum.Pending) {
		return u.Repo.CreatePassword(user)
	}

	return u.Repo.ChangePassword(user)
}

func (u User) GetPassword(id string) (*core.UserWithPassword, error) {
	user, err := u.Repo.User(id)

	if err != nil {
		return nil, err
	}

	user.BasePassword = utils.DecodeString(user.BasePassword)

	return user.RetrievePassword(), nil
}
