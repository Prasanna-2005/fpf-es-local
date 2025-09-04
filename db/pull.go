package db

import (
	"fmt"
	"quantrium-tech/form_pf_extractor_service/config"
)

// PullFiling pulls all the mappings attached to the filingID and returns the list of reduced FileMapper
func PullFiling(filingID int) (*[]FileMapper, error) {

	var mappings []FileMapper
	query := fmt.Sprintf("SELECT filingPath, fundShortCode, quarterNo from %s WHERE filingId=%d", config.TableCfg.MappingTable, filingID)

	err := db.Select(&mappings, query)
	if err != nil {
		return &mappings, fmt.Errorf("retrieving mappings attached to the filingId: " + err.Error())
	}

	if len(mappings) == 0 {
		return &mappings, fmt.Errorf("no mappings attached to the filingId")
	}

	return &mappings, nil

}
