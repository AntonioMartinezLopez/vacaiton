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
	Secret               string
	LimitCountPerRequest int64
	Log_Level            string
	Db_name              string
	Db_User              string
	Db_Password          string `json:"-"`
	Db_Host              string
	Db_Port              string
	Db_LogMode           bool
}

func SetupConfig() (Configuration, error) {

	var configuration *Configuration = new(Configuration)

	// Set Prefix for env variables so only the needed variables are being read
	viper.SetEnvPrefix("USERSERVICE")

	// Set Default
	viper.SetDefault("PORT", "80")
	viper.SetDefault("LOG_LEVEL", "DEBUG")
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("URL", "localhost:8080/tripservice")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_LOGMODE", true)
	viper.SetDefault("DB_SSLMODE", "disable")

	// Load all variables with specified prefix and overwrite default values
	viper.AutomaticEnv()

	// check if obligatory envs are given and set them
	obligatoryEnvs := []string{"SECRET", "DB_HOST", "DB_NAME", "DB_USER", "DB_PASSWORD"}

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

func GetDSNConfig() string {
	DbName := viper.GetString("DB_NAME")
	DbUser := viper.GetString("DB_USER")
	DbPassword := viper.GetString("DB_PASSWORD")
	DbHost := viper.GetString("DB_HOST")
	DbPort := viper.GetString("DB_PORT")
	DbSslMode := viper.GetString("SSLMODE")

	masterDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		DbHost, DbUser, DbPassword, DbName, DbPort, DbSslMode,
	)

	return masterDSN
}
