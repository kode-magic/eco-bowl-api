package infra

import (
	"fmt"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	infra "github.com/kode-magic/eco-bowl-api/infra/entities"
	"github.com/kode-magic/eco-bowl-api/infra/persist"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repositories struct {
	User      core.UserRepository
	Institute core.InstitutionRepo
	Event     core.EventRepo
	Reward    core.RewardRepo
	db        *gorm.DB
}

func DBConfiguration() (*Repositories, error) {
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	DbDSN := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	db, err := gorm.Open(postgres.Open(DbDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	return &Repositories{
		User:      persist.NewUserRepo(db),
		Institute: persist.NewInstituteRepo(db),
		Event:     persist.NewEventRepo(db),
		Reward:    persist.NewRewardRepo(db),
		db:        db,
	}, nil
}

// Close closes the  database connection
func (s *Repositories) Close() error {
	sqlDb, _ := s.db.DB()
	return sqlDb.Close()
}

// AutoMigrate all tables to the database
func (s *Repositories) AutoMigrate() error {
	return s.db.AutoMigrate(
		&infra.User{},
		&infra.Institution{},
		&infra.Event{},
		&infra.Reward{},
	)
}
