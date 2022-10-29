package entities

import (
	"github.com/kode-magic/eco-bowl-api/core/commons"
	"time"
)

type (
	Growth struct {
		ID          string    `json:"id"`
		Year        int8      `json:"year"`
		NetWorth    float64   `json:"netWorth"`
		Income      float64   `json:"income"`
		Expenditure float64   `json:"expenditure"`
		Assets      []string  `json:"assets"`
		BusinessID  string    `json:"businessId"`
		Business    Business  `json:"business"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}

	Business struct {
		ID            string                 `json:"id"`
		Name          string                 `json:"name"`
		Description   string                 `json:"description"`
		Type          commons.BusinessTypes  `json:"type"`
		Level         commons.BusinessLevels `json:"level"`
		Founded       time.Time              `json:"founded"`
		Growths       []Growth               `json:"growths"`
		ContactPerson Contact                `json:"contactPerson"`
		Entrepreneurs []Entrepreneur         `json:"entrepreneurs"`
		CreatedAt     time.Time              `json:"createdAt"`
		UpdatedAt     time.Time              `json:"updatedAt"`
	}

	BusinessRepo interface {
		Create(business *Business) (*Business, map[string]string)
		List() (*[]Business, error)
		Get(id string) (*Business, error)
		GetByName(name string) (*Business, error)
		Update(business *Business) (string, error)
		AddGrowthToBusiness(business *Business)
	}
)
