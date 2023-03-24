package storage

import (
	log "github.com/sirupsen/logrus"
	"github.com/surajn222/url-shortener/pkg/config"
)

func GetStorageObject(config config.Configurations) InterfaceStorage {
	// config := getStorageFromConfig()
	log.Infof("Storage is %+v", config.Storage.Storage_type)
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
