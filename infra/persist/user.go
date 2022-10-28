package persist

import (
	"fmt"
	enum "github.com/kode-magic/eco-bowl-api/core/commons"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	infra "github.com/kode-magic/eco-bowl-api/infra/entities"
	"github.com/kode-magic/eco-bowl-api/utils"
	"gorm.io/gorm"
	"strings"
)

type UserRepo struct {
	db *gorm.DB
}

func ToUserDomain(model infra.User) *core.User {

	return &core.User{
		ID:           model.ID.String(),
		FirstName:    model.FirstName,
		LastName:     model.LastName,
		Phone:        model.Phone,
		Password:     model.Password,
		BasePassword: model.BasePassword,
		Status:       model.Status,
		Role:         enum.Roles(model.Role),
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
	}
}

func (u UserRepo) Add(record *core.User) (*core.User, map[string]string) {
	infraErr := map[string]string{}

	createUser := infra.User{
		FirstName: record.FirstName,
		LastName:  record.LastName,
		Phone:     record.Phone,
		Role:      string(record.Role),
	}

	err := u.db.Create(&createUser).Error

	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			infraErr["email"] = "phone number already exist"
			return nil, infraErr
		}
		infraErr["db_error"] = "database error"
		return nil, infraErr
	}

	return ToUserDomain(createUser), nil
}

func (u UserRepo) Edit(record *core.User) (*core.User, map[string]string) {
	infraErr := map[string]string{}
	var model *infra.User

	err := u.db.Model(&model).Where("id = ?", record.ID).Updates(infra.User{
		LastName:  record.LastName,
		FirstName: record.FirstName,
		Phone:     record.Phone,
		Role:      string(record.Role),
	}).Error

	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			infraErr["email"] = "Phone already exist"
			return nil, infraErr
		}
		infraErr["db_error"] = "database error"
		return nil, infraErr
	}

	return record, nil
}

func (u UserRepo) Users() (*[]core.User, error) {
	var dbUsers []infra.User

	err := u.db.Find(&dbUsers).Error

	if err != nil {
		return nil, err
	}

	users := make([]core.User, len(dbUsers))

	for i, pUsers := range dbUsers {
		toDomain := ToUserDomain(pUsers)

		if err != nil {
			return nil, err
		}

		users[i] = *toDomain
	}

	return &users, nil
}

func (u UserRepo) User(id string) (*core.User, error) {
	var user *infra.User

	err := u.db.Where("id = ?", id).Take(&user).Error

	if err != nil {
		return nil, err
	}

	return ToUserDomain(*user), nil
}

func (u UserRepo) Remove(id string) (string, error) {
	var user *infra.User

	err := u.db.Delete(&user, "id = ?", id).Error

	if err != nil {
		return "", err
	}

	return "User deleted successfully", nil
}

func (u UserRepo) ResetPassword(phone string) (*core.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) Disable(id string) (*core.User, error) {
	var user *infra.User

	err := u.db.Where("id = ?", id).Take(&user).Error

	if err != nil {
		return nil, fmt.Errorf("no user found for this id")
	}

	err = u.db.Model(&user).Where("id = ?", user.ID).Updates(infra.User{
		Status: string(enum.Disable),
	}).Error

	if err != nil {
		return nil, err
	}

	return ToUserDomain(*user), nil
}

func (u UserRepo) Enable(id string) (*core.User, error) {
	var user *infra.User

	err := u.db.Where("id = ?", id).Take(&user).Error

	if err != nil {
		return nil, fmt.Errorf("no user found for this id")
	}

	err = u.db.Model(&user).Where("id = ?", user.ID).Updates(infra.User{
		Status: string(enum.Active),
	}).Error

	if err != nil {
		return nil, err
	}

	return ToUserDomain(*user), nil
}

func (u UserRepo) Login(phone string) (*core.User, error) {
	var user *infra.User

	err := u.db.Where("phone=?", phone).Take(&user).Error

	if err != nil {
		return nil, fmt.Errorf("no account found for this user")
	}

	return ToUserDomain(*user), nil
}

func (u UserRepo) ChangePassword(record core.User) (*core.User, error) {
	var user *infra.User

	record.BasePassword = utils.EncodeString(record.Password)
	record.Password, _ = utils.HashPassword(record.Password)

	err := u.db.Where("id = ?", record.ID).Take(&user).Error

	if err != nil {
		return nil, fmt.Errorf("no user found for this id")
	}

	err = u.db.Model(&user).Where("id = ?", user.ID).Updates(infra.User{
		Password: record.Password,
	}).Error

	if err != nil {
		return nil, err
	}

	return ToUserDomain(*user), nil
}

func (u UserRepo) CreatePassword(record core.User) (*core.User, error) {
	var user *infra.User

	record.Password, _ = utils.HashPassword(record.Password)

	err := u.db.Where("id = ?", record.ID).Take(&user).Error

	if err != nil {
		return nil, fmt.Errorf("no user found for this id")
	}

	err = u.db.Model(&user).Where("id = ?", user.ID).Updates(infra.User{
		Password: record.Password,
		Status:   string(enum.Active),
	}).Error

	if err != nil {
		return nil, err
	}

	return ToUserDomain(*user), nil
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

var _ core.UserRepository = &UserRepo{}
