package db

import (
	"fmt"
	"quantrium-tech/form_pf_extractor_service/config"
)

// PullFiling pulls all the mappings attached to the filingID and returns the list of reduced FileMapper
func PullFiling(filingID string) (*[]FileMapper, error) {

	var mappings []FileMapper
	query := fmt.Sprintf("SELECT filingPath, fundShortCode, quarterNo from %s WHERE filingId= ?", config.TableCfg.MappingTable)

	err := db.Select(&mappings, query, filingID)
	if err != nil {
		return &mappings, fmt.Errorf("retrieving mappings attached to the filingId: " + err.Error())
	}

	if len(mappings) == 0 {
		return &mappings, fmt.Errorf("no mappings attached to the filingId")
	}

	return &mappings, nil

}

func GetExisitingFunds(filingID string) (map[string]struct{}, error) {

	var data []FormData
	ExisitingFunds := make(map[string]struct{})

	query := fmt.Sprintf("SELECT fundSeriesId from %s where filingId=?", config.TableCfg.FormDataTable)

	err := db.Select(&data, query, filingID)
	if err != nil {
		return ExisitingFunds, fmt.Errorf("retrieiving previous form data attached to the filingId: " + err.Error())
	}

	// create a temp reverse map
	mapper := make(map[string]string)
	for k, v := range FundCodeMapper {
		mapper[v] = k
	}

	for _, d := range data {
		ExisitingFunds[mapper[d.FegFundId]] = struct{}{}
	}

	return ExisitingFunds, nil

}

// SetFundInfo sets a local storer with the fund data present in the db
func SetFundInfo() error {

	var funds []FundInfo
	query := fmt.Sprintf("SELECT fundName, fegFundId, nfaId, leiId,fundType from %s", config.TableCfg.FundTable)

	err := db.Select(&funds, query)
	if err != nil {
		return fmt.Errorf("retrieving funds info: " + err.Error())
	}

	if len(funds) == 0 {
		return fmt.Errorf("no funds retrieved")
	}

	// create a temp reverse map
	mapper := make(map[string]string)
	for k, v := range FundCodeMapper {
		mapper[v] = k
	}

	for _, fund := range funds {
		FundStorer[mapper[fund.FundID]] = fund
	}

	return nil
}
