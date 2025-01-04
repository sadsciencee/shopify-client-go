package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	SourceShopifyKey         string
	SourceShopifySecret      string
	SourceShopifyAccessToken string
	SourceShopUrl            string
	DestShopifyKey           string
	DestShopifySecret        string
	DestShopifyAccessToken   string
	DestShopUrl              string
	ApiVersion               string
}

var cfg *Config

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg = &Config{
		SourceShopifyKey:         os.Getenv("SOURCE_SHOPIFY_KEY"),
		SourceShopifySecret:      os.Getenv("SOURCE_SHOPIFY_SECRET"),
		SourceShopifyAccessToken: os.Getenv("SOURCE_SHOPIFY_ACCESS_TOKEN"),
		SourceShopUrl:            os.Getenv("SOURCE_SHOP_URL"),
		DestShopifyKey:           os.Getenv("DEST_SHOPIFY_KEY"),
		DestShopifySecret:        os.Getenv("DEST_SHOPIFY_SECRET"),
		DestShopifyAccessToken:   os.Getenv("DEST_SHOPIFY_ACCESS_TOKEN"),
		DestShopUrl:              os.Getenv("DEST_SHOP_URL"),
		ApiVersion:               os.Getenv("API_VERSION"),
	}
}

func Get() *Config {
	return cfg
}
