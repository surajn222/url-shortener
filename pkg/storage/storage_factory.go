package storage

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/surajn222/url-shortener/pkg/config"
)

func GetStorageObject() InterfaceStorage {
	config := getStorageFromConfig()

	if config.Storage.Storage_type == "memory" {
		return &MemStorage{}
	} else if config.Storage.Storage_type == "redis" {
		return &RedisStorage{
			DBName:     config.Database.DBName,
			DBPort:     config.Database.DBPort,
			DBUser:     config.Database.DBUser,
			DBPassword: config.Database.DBPassword,
		}
	}
	return nil
}

func getStorageFromConfig() config.Configurations {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return configuration

}
