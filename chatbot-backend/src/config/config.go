package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv     		string
	APIPort    		string
	DBHost     		string
	DBPort     		string
	DBUser     		string
	DBPassword 		string
	DBName     		string
	DBDriver   		string
	DBSource   		string 
	ServerAddress	string
	LogLevel   		string
	GeminiApiKey    string 
	PushoverToken   string 
	PushoverUser    string
	XApiKey			string
	
}

func LoadConfig(env string, isTesting bool) (*Config, error) {
	// Load the appropriate .env file based on the environment
	var prefix string = ""
	if isTesting {
		prefix = "../../../"
	}

	var envFile string
	switch env {
	case "staging":
		envFile = prefix + "config/stage.env"
	case "production":
		envFile = prefix + "config/prod.env"
	case "dev":
		envFile = prefix + "config/dev.env"
	default:
		log.Fatalf("Unknown environment: %s", env)
	}
	fmt.Println(envFile)

	if err := godotenv.Load(envFile); err != nil {
		return nil, err
	}

	return &Config{
		AppEnv:     	os.Getenv("APP_ENV"),
		APIPort:    	os.Getenv("API_PORT"),
		DBHost:     	os.Getenv("DB_HOST"),
		DBPort:     	os.Getenv("DB_PORT"),
		DBUser:     	os.Getenv("DB_USER"),
		DBPassword: 	os.Getenv("DB_PASSWORD"),
		DBName:     	os.Getenv("DB_NAME"),
		DBDriver: 		os.Getenv("DB_DRIVER"), 
		DBSource: 		os.Getenv("DB_SOURCE"), 
		ServerAddress: 	os.Getenv("SERVICE_ADDRESS"), 
		LogLevel:   	os.Getenv("LOG_LEVEL"),
		GeminiApiKey: 	os.Getenv("GEMINI_API_KEY"), 
		PushoverUser:   os.Getenv("PUSHOVER_USER"), 
		PushoverToken:  os.Getenv("PUSHOVER_TOKEN"),
		XApiKey: 		os.Getenv("X_API_KEY"),
	}, nil
}