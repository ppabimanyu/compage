package configloader

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log/slog"
	"reflect"
)

func Load(conf any) error {
	val := reflect.ValueOf(conf)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return errors.New("conf must be a pointer to struct")
	}

	if err := godotenv.Load(); err != nil {
		slog.Warn("ConfigLoader: Failed to load config from file: .env")
		slog.Info("ConfigLoader: Trying to load from environment variables")
	}

	if err := envconfig.Process("", conf); err != nil {
		return err
	}

	slog.Info("ConfigLoader: Configuration loaded successfully", "config", conf)
	return nil
}
