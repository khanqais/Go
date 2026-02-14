package config

import (
	"fmt"
	"os"
	"strings"	
	"github.com/joho/godotenv"
)

type Config struct {
	MongoURL    string
	MongoDBName string
	JWT_SECRET  string
}

func Load() (Config, error) {
	_ = godotenv.Load()
	cfg := Config{
		MongoURL:    strings.TrimSpace(os.Getenv("MONGO_URI")),
		MongoDBName: strings.TrimSpace(os.Getenv("MONGO_DB_NAME")),
		JWT_SECRET:  strings.TrimSpace(os.Getenv("JWT_SECRET")),
	}
	if cfg.MongoURL == "" {
		return Config{}, fmt.Errorf("missing mongo URL")
	}
	if cfg.MongoDBName == "" {
		return Config{}, fmt.Errorf("missing mongo DBNAME")
	}
	if cfg.JWT_SECRET == "" {
		return Config{}, fmt.Errorf("missing JWT")
	}
	return cfg, nil
}
