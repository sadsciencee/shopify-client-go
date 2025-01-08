package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ShopifyApiKey          string
	TestShopifyAccessToken string
	TestShopUrl            string
	ApiVersion             string
}

var cfg *Config

func Load(envPath string) {
	if os.Getenv("CI") != "true" {
		err := godotenv.Load(envPath)
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	cfg = &Config{
		ShopifyApiKey:          os.Getenv("SHOPIFY_API_KEY"),
		TestShopifyAccessToken: os.Getenv("TEST_SHOPIFY_ACCESS_TOKEN"),
		TestShopUrl:            os.Getenv("TEST_SHOP_URL"),
		ApiVersion:             os.Getenv("API_VERSION"),
	}
}

func Get() *Config {
	return cfg
}
