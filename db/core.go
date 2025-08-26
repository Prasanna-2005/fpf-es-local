package db

import (
	"fmt"
	"quantrium-tech/form_pf_extractor_service/config"
)

// FileMapper represents the reduced mapper with respect to DB that only holds the neccessary information for execution
type FileMapper struct {
	FilingPath    string `db:"filingPath"`
	FundShortCode string `db:"fundShortCode"`
	QuarterNo     int    `db:"quarterNo"`
}

// PullFiling pulls all the mappings attached to the filingID and returns the list of reduced FileMapper
func PullFiling(filingID int) (*[]FileMapper, error) {

	var mappings []FileMapper
	query := fmt.Sprintf("SELECT * from %s WHERE filingId=%d", config.TableCfg.MappingTable, filingID)

	err := db.Select(&mappings, query)
	if err != nil {
		return &mappings, fmt.Errorf("retrieving mappings attached to the filingId: " + err.Error())
	}

	return &mappings, nil

}
