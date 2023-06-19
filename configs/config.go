package configs

import (
	"log"

	"github.com/spf13/viper"
)

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBConn        string `mapstructure:"DB_CONNECTION"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
}

func LoadConfig(path string) *conf {
	var cfg *conf
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
