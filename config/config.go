package config

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Database   *Database `json:"database"`
		Port       Port
		Collection Collection
	}

	Database struct {
		Uri      string `json:"uri"`
		Name     string `json:"databasename"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Port struct {
		HTTPServer string
	}

	Collection struct {
		Offer    string `json:"offer"`
		Customer string `json:"customer"`
	}
)

const (
	osEnv = "FSHENV"

	// environment development, staging and production
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"

	AppsURLDevelopment = "http://localhost:3001"
)

var (
	thisEnv    string
	thisConfig Config
)

func GetEnv() (env string) {
	if thisEnv != "" {
		return thisEnv
	}

	env = os.Getenv(osEnv)
	if env == "" {
		env = EnvDevelopment
	}

	thisEnv = env
	return
}

//  Main for config
func New() (config Config, err error) {
	log.Println("Config starting...")

	env := GetEnv()

	viper.SetConfigType("yml")
	// Read yml file for switching env
	if env == "production" {
		viper.SetConfigName("env_production")
	} else {
		viper.SetConfigName("env_developmnet")
	}
	viper.AddConfigPath("./env/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config, %s", err)
	}

	err = viper.Unmarshal(&config)
	if err == nil {
		thisConfig = config
		return
	}

	return config, errors.Wrapf(err, "failed to unmarshall config")
}

// Setter and Getter Config
func Set(cfg Config) {
	thisConfig = cfg
}

var Get = func() (cfg Config) {
	cfg, _ = New()
	thisConfig = cfg

	return cfg
}
