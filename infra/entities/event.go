package entities

import (
	"github.com/google/uuid"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
	"time"
)

type Event struct {
	ID            uuid.UUID `gorm:"type:uuid;"`
	Name          string    `gorm:"not null;"`
	Description   string    `gorm:"not null;"`
	StartDate     time.Time `gorm:"not null;"`
	EndDate       time.Time `gorm:"not null;"`
	InstitutionID string    `gorm:"not null;"`
	Institution   Institution
	Trainees      []Trainee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Teams         []Team    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (e *Event) BeforeCreate(_ *gorm.DB) error {
	e.ID = ulids.GenerateUUID()
	return nil
}

func (e *Event) BeforeUpdate(_ *gorm.DB) error {
	e.UpdatedAt = time.Now()
	return nil
}

func (e *Event) Prepare() {
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
}
