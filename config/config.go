package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MySQLURI       string
	MongoURI       string
	DBName         string
	CollectionName string
	Port           string
	RedisHost      string
	RateNumber     int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, reading from environment")
	}

	rateInt, err := strconv.Atoi(os.Getenv("RATE_NUMBER"))
	if err != nil {
		rateInt = 10
	}

	return &Config{
		MySQLURI:       os.Getenv("MYSQL_URI"),
		MongoURI:       os.Getenv("MONGO_URI"),
		DBName:         os.Getenv("DB_NAME"),
		CollectionName: os.Getenv("COLLECTION_NAME"),
		Port:           getOrDefault("PORT", ":9999"),
		RedisHost:      os.Getenv("REDIS_HOST"),
		RateNumber:     rateInt,
	}
}

func getOrDefault(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
