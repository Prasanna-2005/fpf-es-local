package config

import (
	"fmt"
	"os"
	"quantrium-tech/form_pf_extractor_service/logging"

	"gopkg.in/ini.v1"
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

type InvLiqTypes struct {
	UsPerson               string `ini:"UsPerson"`
	NotUsPerson            string `ini:"NotUsPerson"`
	Broker                 string `ini:"Broker"`
	Insurance              string `ini:"Insurance"`
	RIA                    string `ini:"Ria"`
	PrivateFund            string `ini:"PrivateFund"`
	NonProfit              string `ini:"NonProfit"`
	PensionPlans           string `ini:"PensionPlans"`
	Proprietary            string `ini:"Proprietary"`
	Government             string `ini:"Government"`
	GovernmentPensionPlans string `ini:"GovernmentPensionPlans"`
	SovereignFund          string `ini:"SovereignFund"`
	UnknownInstitution     string `ini:"UnknownInstitution"`
	Other                  string `ini:"Other"`
}

var (
	Cfg              *config
	TableCfg         *tableConfig
	InvLiqTypeMarker *InvLiqTypes
)

// InitialiseConfig initialises the exported Cfg for other modules to use
// It Acquires the runtime env variables and assign them
func InitialiseConfig() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		logging.Error(fmt.Sprintf("Investor liq config load error : %v", err.Error()))
	}
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

	InvLiqTypeMarker = &InvLiqTypes{
		UsPerson:               cfg.Section("InvestorLiquidity").Key("UsPerson").String(),
		NotUsPerson:            cfg.Section("InvestorLiquidity").Key("NotUsPerson").String(),
		Broker:                 cfg.Section("InvestorLiquidity").Key("Broker").String(),
		Insurance:              cfg.Section("InvestorLiquidity").Key("Insurance").String(),
		RIA:                    cfg.Section("InvestorLiquidity").Key("RIA").String(),
		PrivateFund:            cfg.Section("InvestorLiquidity").Key("PrivateFund").String(),
		NonProfit:              cfg.Section("InvestorLiquidity").Key("NonProfit").String(),
		PensionPlans:           cfg.Section("InvestorLiquidity").Key("PensionPlans").String(),
		Proprietary:            cfg.Section("InvestorLiquidity").Key("Proprietary").String(),
		Government:             cfg.Section("InvestorLiquidity").Key("Government").String(),
		GovernmentPensionPlans: cfg.Section("InvestorLiquidity").Key("GovernmentPensionPlans").String(),
		SovereignFund:          cfg.Section("InvestorLiquidity").Key("SovereignFund").String(),
		UnknownInstitution:     cfg.Section("InvestorLiquidity").Key("UnknownInstitution").String(),
		Other:                  cfg.Section("InvestorLiquidity").Key("Other").String(),
	}

}
