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
	AccessTokenTime      int
	RefrashTokenTime     int
	MailPort		int
	MailHost		string
	MailUser		string
	MailPass		string
	MailFrom		string
}

func LoadConfig() *Config { 
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, reading from environment")
	}

	rateInt, _ := strconv.Atoi(os.Getenv("RATE_NUMBER"))
	mailPort, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))

	accessToken, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_TIME"))
	refrashToken, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_TIME"))

	return &Config{
		MySQLURI:       os.Getenv("MYSQL_URI"),
		MongoURI:       os.Getenv("MONGO_URI"),
		DBName:         os.Getenv("DB_NAME"),
		CollectionName: os.Getenv("COLLECTION_NAME"),
		Port:           getOrDefault("PORT", ":9999"),
		RedisHost:      os.Getenv("REDIS_HOST"),
		AccessTokenTime:       accessToken,
		RefrashTokenTime:      refrashToken,
		RateNumber:     rateInt,

		MailPort:     mailPort,
		MailHost:     os.Getenv("MAIL_HOST"),
		MailUser:     os.Getenv("MAIL_USER"),
		MailPass:     os.Getenv("MAIL_PASS"),
		MailFrom:     os.Getenv("MAIL_FROM"),
	}
}

func getOrDefault(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
