package util

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadEnvVars(path string) (config Config, err error) {
	viper.AddConfigPath(path)  // path to look for the config file in
	viper.SetConfigName("app") // name of config file (without extension)
	viper.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// viper.AutomaticEnv() // allows to overwrite env variable from command line

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Unable to load ennvironment variables: ", err.Error())
		return
	}

	return config, err
}
