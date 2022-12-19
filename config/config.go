package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	App         string
	AppVersion  string
	Environment string // development, staging, production

	GRPCPort string

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
}

// Load ...
func Load() Config {
	// .env file bor yo'qligini tekshirilvotti
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("APP", "article"))
	config.AppVersion = cast.ToString(getOrReturnDefaultValue("APP_VERSION", "1.0.0"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))

	config.GRPCPort = cast.ToString(getOrReturnDefaultValue("GRPC_PORT", ":9001"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "test"))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "test"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "test"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
