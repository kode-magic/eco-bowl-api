package entities

import (
	"github.com/google/uuid"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
	"time"
)

type Reward struct {
	ID          uuid.UUID `gorm:"type:uuid;"`
	Name        string    `gorm:"not null;"`
	Description string    `gorm:"default:null;"`
	Offers      []string  `gorm:"type:text[];not null;"`
	EventID     string    `gorm:"not null;"`
	Event       Event
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (r *Reward) BeforeCreate(_ *gorm.DB) error {
	r.ID = ulids.GenerateUUID()
	return nil
}

func (r *Reward) BeforeUpdate(_ *gorm.DB) error {
	r.UpdatedAt = time.Now()
	return nil
}

func (r *Reward) Prepare() {
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}