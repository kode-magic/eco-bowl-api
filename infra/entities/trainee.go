package entities

import (
	"github.com/google/uuid"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
	"time"
)

type Trainee struct {
	ID            uuid.UUID `gorm:"type:uuid;"`
	Forename      string    `gorm:"not null;"`
	Surname       string    `gorm:"not null;"`
	Gender        string    `gorm:"not null;"`
	BirthDate     time.Time `gorm:"not null;"`
	Phone         string    `gorm:"not null;"`
	Email         string    `gorm:"not null;"`
	Qualification string    `gorm:"default:null;"`
	EventID       string    `gorm:"not null;"`
	Event         Event
	TeamID        string `gorm:"default:null;"`
	Team          Team
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (t *Trainee) BeforeCreate(_ *gorm.DB) error {
	t.ID = ulids.GenerateUUID()
	return nil
}

func (t *Trainee) BeforeUpdate(_ *gorm.DB) error {
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Trainee) Prepare() {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}
