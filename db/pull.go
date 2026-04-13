package db

import (
	"fmt"
	"quantrium-tech/form_pf_extractor_service/config"
	"strings"
)

// PullFiling pulls all the mappings attached to the filingID and returns the list of reduced FileMapper
func PullFiling(filingID string) (*[]FileMapper, error) {

	var mappings []FileMapper
	query := fmt.Sprintf("SELECT filingPath, fundShortCode, quarterNo,period from %s WHERE filingId= ?", config.TableCfg.MappingTable)

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

type AUM struct {
	FundType   string  `db:"fundType"`
	FieldNo    string  `db:"fieldNo"`
	FormLineNo string  `db:"formLineNo"`
	Asset      float64 `db:"assets"` // Matches the 'AS total_value' in SQL
}

func PullAssets(filingID string) (map[string]map[string]map[string]float64, error) {
	var rows []AUM
	// Map structure: [FundType][LineNo][FieldNo]AssetValue
	finalData := make(map[string]map[string]map[string]float64)

	query := `
        SELECT 
            funds.fundType, 
            legacy.fieldNo,
            legacy.formLineNo,
            SUM(legacy.value) AS assets
        FROM 
            form_data_legacy AS legacy
        LEFT JOIN 
            feg_funds AS funds
            ON legacy.fundSeriesId = funds.fegFundId
        WHERE 
            legacy.filingId = ?
            AND legacy.fieldNo IN ('8', '9') 
            AND legacy.formLineNo IN ('8', '9')
			AND isCurrentValue=1
        GROUP BY 
            funds.fundType, 
            legacy.formLineNo,
            legacy.fieldNo`

	err := db.Select(&rows, query, filingID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving detailed fund assets: %w", err)
	}

	for _, row := range rows {
		fType := strings.ToLower(row.FundType)
		if fType == "" {
			fType = "unknown"
		}

		// Level 1: Initialize FundType
		if _, ok := finalData[fType]; !ok {
			finalData[fType] = make(map[string]map[string]float64)
		}

		// Level 2: Initialize FormLineNo
		if _, ok := finalData[fType][row.FormLineNo]; !ok {
			finalData[fType][row.FormLineNo] = make(map[string]float64)
		}

		// Level 3: Set Asset value for the FieldNo
		finalData[fType][row.FormLineNo][row.FieldNo] = row.Asset
	}

	return finalData, nil
}
