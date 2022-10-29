package entities

import (
	"github.com/google/uuid"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
	"time"
)

type (
	Team struct {
		ID          uuid.UUID `gorm:"type:uuid;"`
		Name        string    `gorm:"not null;"`
		Description string    `gorm:"default:null;"`
		EventID     string    `gorm:"not null;"`
		Event       Event
		Trainees    []Trainee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}
)

func (t *Team) BeforeCreate(_ *gorm.DB) error {
	t.ID = ulids.GenerateUUID()
	return nil
}

func (t *Team) BeforeUpdate(_ *gorm.DB) error {
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Team) Prepare() {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}
