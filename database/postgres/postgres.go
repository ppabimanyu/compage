package postgres

import (
	"errors"
	"fmt"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

type Config struct {
	Host            string
	Port            int
	Username        string
	Password        string
	Database        string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration
}

func NewConnection(config *Config) (*gorm.DB, error) {
	if config == nil {
		return nil, errors.New("config is nil")
	}
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == 0 {
		config.Port = 5432
	}
	if config.Username == "" {
		config.Username = "postgres"
	}
	if config.Password == "" {
		config.Password = "postgres"
	}
	if config.Database == "" {
		config.Database = "postgres"
	}
	if config.MaxIdleConn == 0 {
		config.MaxIdleConn = 10
	}
	if config.MaxOpenConn == 0 {
		config.MaxOpenConn = 100
	}
	if config.ConnMaxIdleTime == 0 {
		config.ConnMaxIdleTime = 5 * time.Minute
	}
	if config.ConnMaxLifetime == 0 {
		config.ConnMaxLifetime = 30 * time.Minute
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.Host, config.Username, config.Password, config.Database, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: slogGorm.New(),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConn)
	sqlDB.SetMaxOpenConns(config.MaxOpenConn)
	sqlDB.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)

	slog.Info("PostgresDB: Connection established")

	return db, nil
}
