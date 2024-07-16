package config

import (
	"log"
	"os"

	"github.com/spf13/cast"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP_PORT             string
	AUTH_SERVICE_PORT     string
	PRODUCT_SERVICE       string
	PAYMENTS_SERVICE_PORT string
	SIGNING_KEY           string
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load .env ", err)
	}

	cfg := Config{}

	cfg.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":8888"))
	cfg.AUTH_SERVICE_PORT = cast.ToString(coalesce("AUTH_SERVICE_PORT", ":7777"))
	cfg.RESERVATION_SERVICE = cast.ToString(coalesce("RESERVATION_SERVICE", ":9999"))
	cfg.PAYMENTS_SERVICE_PORT = cast.ToString(coalesce("PAYMENTS_SERVICE_PORT", "GARD"))

	return &cfg
}

func coalesce(key string, value interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return value
}
