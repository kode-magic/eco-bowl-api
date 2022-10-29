package entities

import (
	"github.com/google/uuid"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
	"time"
)

type Entrepreneur struct {
	ID         uuid.UUID `gorm:"type:uuid;"`
	Forename   string    `gorm:"not null;"`
	Surname    string    `gorm:"not null;"`
	Gender     string    `gorm:"not null;"`
	BirthDate  time.Time `gorm:"not null;"`
	Phone      string    `gorm:"not null;"`
	Email      string    `gorm:"not null;"`
	BusinessID string    `gorm:"default:null;"`
	Business   Business
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (t *Entrepreneur) BeforeCreate(_ *gorm.DB) error {
	t.ID = ulids.GenerateUUID()
	return nil
}

func (t *Entrepreneur) BeforeUpdate(_ *gorm.DB) error {
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Entrepreneur) Prepare() {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}
