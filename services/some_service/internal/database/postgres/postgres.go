package postgres

import (
	"fmt"
	"log"

	"github.com/tonitaga/some_service/internal/config"

	pq_driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorage(configuration config.PostgresConfig) (*postgresStorage, error) {
	connectionURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configuration.Host,
		configuration.Port,
		configuration.User,
		configuration.Password,
		configuration.Name,
	)

	db, err := gorm.Open(pq_driver.Open(connectionURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("[NewPostgresStorage] Error on opening postgres DB. Cause: %v", err)
	}

	return &postgresStorage{
		db: db,
	}, nil
}

func (p *postgresStorage) AutoMigrate() error {
	log.Println("[PostgresStorage] Making auto migrations")
	return nil
}
