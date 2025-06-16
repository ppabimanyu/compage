package gormutils

import (
	"errors"
	"gorm.io/gorm"
	"log/slog"
)

type Model interface {
	TableName() string
}

func DBMigration(db *gorm.DB, models ...Model) error {
	if db == nil {
		return errors.New("db is nil")
	}
	if len(models) == 0 {
		return errors.New("no models provided for migration")
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			slog.Error("DBMigration: AutoMigrate failed", "model", model.TableName(), "error", err)
			return err
		}
		slog.Info("DBMigration: AutoMigrate successful", "model", model.TableName())
	}

	return nil
}
