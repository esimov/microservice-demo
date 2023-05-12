package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	ApiSecret string `mapstructure:"API_SECRET"`
	UserName  string `mapstructure:"MYSQL_USERNAME"`
	Password  string `mapstructure:"MYSQL_PASSWORD"`
	HostName  string `mapstructure:"MYSQL_HOSTNAME"`
	Port      string `mapstructure:"MYSQL_PORT"`
	DB        string `mapstructure:"MYSQL_DATABASE"`
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.SetConfigType("env")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.Join(err, errors.New("config file was not found"))
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	c := new(Config)

	err := v.Unmarshal(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Get config
func GetConfig(configFile string) (*Config, error) {
	cfgFile, err := LoadConfig(configFile)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
