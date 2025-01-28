package utils

import (
	"log"

	"github.com/spf13/viper"
)

var PrefixEnviromentVariables = "quill"

func InitEnviromentVariables() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(PrefixEnviromentVariables)
}

func GetEnviromentVariable(name string) string {
	envVar := viper.GetString(name)
	if envVar == "" {
		log.Fatalf("Failure with reading enviroment variable: %v", name)
	}
	return envVar
}
