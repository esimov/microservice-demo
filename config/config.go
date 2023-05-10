package config

import "errors"

type Config struct {
	MySQL
}

type MySQL struct {
	UserName string
	Password string
	HostName string
	Port     string
	DB       string
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
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
	c := &Config{}

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
