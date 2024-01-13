package config

import (
	"backend/pkg/logger"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Port                 string
	LimitCountPerRequest int64
	Log_Level            string
	OpenAI_Secret        string
}

func SetupConfig() (Configuration, error) {

	var configuration *Configuration = new(Configuration)

	// Set Prefix for env variables so only the needed variables are being read
	viper.SetEnvPrefix("TRIPBUILDER_SERVICE")

	// Set Default
	viper.SetDefault("PORT", "80")
	viper.SetDefault("LOG_LEVEL", "DEBUG")
	viper.SetDefault("HOST", "0.0.0.0")

	// Load all variables with specified prefix and overwrite default values
	viper.AutomaticEnv()

	// check if obligatory envs are given and set them
	obligatoryEnvs := []string{"OPENAI_SECRET"}

	for _, s := range obligatoryEnvs {
		valueExists := viper.IsSet(s)
		if !valueExists {
			return *configuration, errors.New(fmt.Sprintf("Env %s not given", s))
		}
		viper.BindEnv(s, viper.GetString(s))
	}

	// unmarshall values to the config object and return the value
	err := viper.Unmarshal(configuration)
	if err != nil {
		logger.Error("error to decode, %v", err)
		return *configuration, err
	}

	// log information
	strConfig, _ := json.Marshal(*configuration)
	logger.Info(fmt.Sprintf("Used Configuration: %s", strConfig))

	return *configuration, nil
}

func ServerConfig() string {
	appServer := fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT"))
	logger.Log("Server Running at %s", appServer)
	return appServer
}
