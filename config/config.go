package config

import (
	"os"
)

// config represents the overall configuration of the service
type config struct {
	SockPath    string
	StoragePath string
	DBUser      string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
}

// tableConfig represents the DB Table configuration
type tableConfig struct {
	MappingTable    string
	ExtractionTable string
	FormDataTable   string
	FundTable       string
}

var (
	Cfg      *config
	TableCfg *tableConfig
)

// InitialiseConfig initialises the exported Cfg for other modules to use
// It Acquires the runtime env variables and assign them
func InitialiseConfig() {
	Cfg = &config{
		SockPath:    os.Getenv("SOCKET_PATH"),
		StoragePath: os.Getenv("STORAGE_PATH"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBName:      os.Getenv("DB_NAME"),
	}

	TableCfg = &tableConfig{
		MappingTable:    os.Getenv("MAPPING_TABLE"),
		ExtractionTable: os.Getenv("EXTRACTION_TABLE"),
		FormDataTable:   os.Getenv("FORM_DATA_TABLE"),
		FundTable:       os.Getenv("FUND_TABLE"),
	}

}
