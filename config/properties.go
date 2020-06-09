package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// List of Constant Props
const (
	PORT = ":1220"
)

// Config : struct to declare the properties
type Config struct {
	DbUser   string
	DbHost   string
	DbKey    string
	DbSchema string
	DbPort   string
}

// GetProps : Returns Config struct with config values; calls getProp() internally
func GetProps() *Config {
	return &Config{
		DbUser:   getProp("DB_HOST", ""),
		DbHost:   getProp("DB_USER", ""),
		DbKey:    getProp("DB_KEY", ""),
		DbPort:   getProp("DB_PORT", "3306"),
		DbSchema: getProp("DB_SCHEMA", ""),
	}
}

var props map[string]string

// SetEnvProps : Sets environment variable for the svc
func SetEnvProps(env, path string) {
	if strings.ToLower(env) == "local" {
		if path == "" {
			err := godotenv.Load("./config/properties_local.env")
			if err != nil {
				log.Fatal("Error Loading Properties File")
			}
		} else {
			err := godotenv.Load(path)
			if err != nil {
				log.Fatal("Error Loading Properties File from Path: ", path)
			}
		}
	} else if strings.ToLower(env) == "prod" || strings.ToLower(env) == "production" {
		if path == "" {
			err := godotenv.Load("./config/properties_prod.env")
			if err != nil {
				log.Fatal("Error Loading Properties File")
			}
		} else {
			err := godotenv.Load(path)
			if err != nil {
				log.Fatal("Error Loading Properties File from Path: ", path)
			}
		}
	}
}

// getProp : Return the environment value of the required property if exists else returns the default value
func getProp(prop, defaultVal string) string {
	if val, exists := os.LookupEnv(prop); exists {
		return val
	}
	return defaultVal
}
