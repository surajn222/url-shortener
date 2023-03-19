package config

// Configurations exported
type Configurations struct {
	Storage  StorageConfigurations
	Database DatabaseConfigurations
}

// ServerConfigurations exported
type StorageConfigurations struct {
	Storage_type string
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBName     string
	DBPort     int
	DBUser     string
	DBPassword string
}
