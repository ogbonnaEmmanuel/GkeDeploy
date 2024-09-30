package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	RedisAddress  string `mapstructure:"REDIS_ADDRESS"`
	MONGO_URI     string `mapstructure:"MONGO_URI"`
	MONGO_DB_NAME string `mapstructure:"MONGO_DB_NAME"`
}

func LoadConfig(path string) (config Config, err error) {

	// Set config file options (optional, only if you are using a config file)
	viper.AddConfigPath(path)     // Config file path
	viper.SetConfigName("app")     // Config file name (without extension)
	viper.SetConfigType("env")     // Config file type
	viper.AutomaticEnv()           // Automatically read environment variables

	// Bind specific environment variables explicitly
	viper.BindEnv("REDIS_ADDRESS")
	viper.BindEnv("MONGO_URI")
	viper.BindEnv("MONGO_DB_NAME")

	// Try reading the config file if available
	err = viper.ReadInConfig()
	if err != nil {
		// Log the error but continue to read env variables
		log.Printf("Config file not found: %v, using environment variables if set", err)
	}

	// Unmarshal the environment variables into the Config struct
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
