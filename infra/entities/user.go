package entities

import (
	"github.com/kode-magic/eco-bowl-api/core/commons"
	"github.com/kode-magic/eco-bowl-api/utils"
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
)

type (
	User struct {
		ID           uuid.UUID `gorm:"type:uuid;" json:"id"`
		FirstName    string    `gorm:"size:100;not null;"`
		LastName     string    `gorm:"size:100;not null;"`
		Phone        string    `gorm:"size:100;not null;unique;"`
		Password     string    `gorm:"size:100;"`
		BasePassword string    `gorm:"size:100;"`
		Status       string    `gorm:"size:100;"`
		Role         string    `gorm:"not null;"`
		CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}
)

func (user *User) BeforeCreate(_ *gorm.DB) error {
	password := utils.NumberGenerator(6)
	user.ID = ulids.GenerateUUID()
	user.Status = string(commons.Pending)
	user.Password, _ = utils.HashPassword(password)
	user.BasePassword = utils.EncodeString(password)
	return nil
}

func (user *User) BeforeUpdate(_ *gorm.DB) error {
	user.UpdatedAt = time.Now()
	return nil
}

func (user *User) Prepare() {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Phone = html.EscapeString(strings.TrimSpace(user.Phone))
}
