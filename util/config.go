package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config contains the env settings used for program
// configuration
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	LisAddr             string        `mapstructure:"LIS_ADDR"`
	TokenSymKey         string        `mapstructure:"TOKEN_SYM_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DUR"`
}

// LoadConfig reads a config file given the path
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
