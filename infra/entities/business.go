package entities

import (
	"github.com/google/uuid"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
	"time"
)

type (
	Growth struct {
		ID          uuid.UUID `gorm:"type:uuid;"`
		Year        int64     `gorm:"not null;"`
		NetWorth    float64   `gorm:"not null;"`
		Income      float64   `gorm:"not null;"`
		Expenditure float64   `gorm:"not null;"`
		Assets      string    `gorm:"default:null;"`
		BusinessID  string    `gorm:"not null;"`
		Business    Business
		CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	Business struct {
		ID            uuid.UUID      `gorm:"type:uuid;"`
		Name          string         `gorm:"not null;"`
		Description   string         `gorm:"not null;"`
		Type          string         `gorm:"not null;"`
		Level         string         `gorm:"not null;"`
		Founded       time.Time      `gorm:"not null;"`
		Growths       []Growth       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		ContactPerson Contact        `gorm:"type:jsonb;"`
		Entrepreneurs []Entrepreneur `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		CreatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	}
)

func (b *Business) BeforeCreate(_ *gorm.DB) error {
	b.ID = ulids.GenerateUUID()
	return nil
}

func (b *Business) BeforeUpdate(_ *gorm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Business) Prepare() {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}

func (g *Growth) BeforeCreate(_ *gorm.DB) error {
	g.ID = ulids.GenerateUUID()
	return nil
}

func (g *Growth) BeforeUpdate(_ *gorm.DB) error {
	g.UpdatedAt = time.Now()
	return nil
}

func (g *Growth) Prepare() {
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
}
