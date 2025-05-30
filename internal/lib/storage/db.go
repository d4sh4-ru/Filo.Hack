package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Filo.Hack/config"
)

// DBClient - структура для работы с БД
type DBClient struct {
	Db *gorm.DB
}

// NewDBClient - функция для инициализации нового клиента БД
func NewDBClient(cfg *config.Config) (*DBClient, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s search_path=%s port=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
		cfg.Postgres.Schema,
		cfg.Postgres.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	return &DBClient{Db: db}, nil
}

// CloseDBConnection - функция для закрытия соединения с базой данных
func CloseDBConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get DB instance: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close DB connection: %v", err)
	}

	return nil
}
