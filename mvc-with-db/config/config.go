package config

import (
	"belajar-go-3/mvc-with-db/libraries"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./config/config.json")
	err := viper.ReadInConfig()
	libraries.CheckError(err)
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}
