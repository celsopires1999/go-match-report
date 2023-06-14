package configs

import (
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBConn   string `mapstructure:"DB_CONNECTION"`
}

func LoadConfig(path string) *conf {
	var cfg *conf
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
