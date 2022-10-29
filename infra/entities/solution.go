package entities

import (
	"github.com/google/uuid"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
	"time"
)

type Solution struct {
	ID          uuid.UUID `gorm:"type:uuid;"`
	Title       string    `gorm:"not null;"`
	Description string    `gorm:"default:null;"`
	EventID     string    `gorm:"not null;"`
	Event       Event
	TeamID      string `gorm:"not null;"`
	Team        Team
	RewardID    string `gorm:"default:null;"`
	Reward      Reward
	Position    int       `gorm:"default:0"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (r *Solution) BeforeCreate(_ *gorm.DB) error {
	r.ID = ulids.GenerateUUID()
	return nil
}

func (r *Solution) BeforeUpdate(_ *gorm.DB) error {
	r.UpdatedAt = time.Now()
	return nil
}

func (r *Solution) Prepare() {
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}
