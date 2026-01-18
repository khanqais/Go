package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoUrl   string
	MongoDB    string
	ServerPort string
}

func Load() (Config, error) {
	//godotenv.Load() read
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("Failed to load .env")

	}
	mongoURL, err := ExtractEnv("MONGO_URI")
	if err != nil {
		return Config{}, err
	}
	mongoDB, err := ExtractEnv("MONGO_DB_NAME")
	if err != nil {
		return Config{}, err
	}
	port, err := ExtractEnv("PORT")
	if err != nil {
		return Config{}, err
	}
	return Config{
		MongoUrl:   mongoURL,
		MongoDB:    mongoDB,
		ServerPort: port,
	}, nil

}

func ExtractEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("missing req env")
	}
	return val, nil

}
