package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the applications
type Config struct {
	Loglevel string   `json:"loglevel" yaml:"loglevel"`
	Database Database `json:"database" yaml:"database"`
	Server   Server   `json:"server" yaml:"server"`
}

type Server struct {
	ServerAddress       string        `json:"serverAddress" yaml:"serverAddress"`
	TokenSymmetricKey   string        `json:"tokenSymmetricKey" yaml:"tokenSymmetricKey"`
	AccessTokenDuration time.Duration `json:"accessTokenDuration" yaml:"accessTokenDuration"`
}

type Database struct {
	DriverName     string `json:"driverName" yaml:"driverName"`
	DataSourceName string `json:"dataSourceName" yaml:"dataSourceName"`
	MigrationURL   string `json:"migrationURL" yaml:"migrationURL"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
