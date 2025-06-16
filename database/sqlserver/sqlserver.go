package sqlserver

import (
	"errors"
	"fmt"
	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/sqlserver"
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
	if config.Username == "" {
		config.Username = "sa"
	}
	if config.Database == "" {
		return nil, errors.New("database name cannot be empty")
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

	var hostStr string
	if config.Port == 0 {
		hostStr = fmt.Sprintf("%s", config.Host)
	} else {
		hostStr = fmt.Sprintf("%s:%d", config.Host, config.Port)
	}

	dsn := fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s",
		config.Username,
		config.Password,
		hostStr,
		config.Database,
	)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
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

	slog.Info("SQLServerDB: Connection established")

	return db, nil
}
