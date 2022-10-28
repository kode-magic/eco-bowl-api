package entities

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/kode-magic/go-bowl/ulids"
	"gorm.io/gorm"
	"time"
)

type (
	Contact struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	Institution struct {
		ID            uuid.UUID `gorm:"type:uuid;"`
		Name          string    `gorm:"not null;"`
		Description   string    `gorm:"default:null;"`
		Address       string    `gorm:"not null;"`
		ContactPerson Contact   `gorm:"type:jsonb;"`
		CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}
)

func (i *Institution) BeforeCreate(_ *gorm.DB) error {
	i.ID = ulids.GenerateUUID()
	return nil
}

func (i *Institution) BeforeUpdate(_ *gorm.DB) error {
	i.UpdatedAt = time.Now()
	return nil
}

func (i *Institution) Prepare() {
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()
}

func (c Contact) Value() (driver.Value, error) {
	valueString, err := json.Marshal(c)
	return string(valueString), err
}

func (c *Contact) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &c); err != nil {
		return err
	}
	return nil
}
