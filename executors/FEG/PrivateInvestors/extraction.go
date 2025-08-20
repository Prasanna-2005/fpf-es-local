package privateinvestors

import (
	"fmt"
	"quantrium-tech/form_pf_extractor_service/executors"
	"quantrium-tech/form_pf_extractor_service/logging"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// extractQ15Data extracts all the funds Q15 answer and returns them with fund mapped values
func extractQ15Data(f *excelize.File) (*map[string]float64, error) {

	q15Holder := make(map[string]float64)

	Q15rows, err := f.GetRows("Q15")
	if err != nil {
		return &q15Holder, fmt.Errorf("getting q15 rows: " + err.Error())
	}

	markers := make(map[string]string)

	for i, row := range Q15rows {
		for j, cellValue := range row {
			if strings.Contains(cellValue, "FEG") {
				markers[cellValue] = executors.IndexToExcel(i+2, j+3)
			}
		}
	}
	for fund, location := range markers {
		value, _ := f.GetCellValue("Q15", location)
		value = executors.CleanNonNumeric(value)

		percentage, _ := strconv.ParseFloat(value, 64)

		q15Holder[fund] = percentage
	}

	return &q15Holder, nil
}

// getQ16Markers retrieves all the markers used to get the contributions of investors for a particular fund from Q16 sheet
func getQ16Markers(rows [][]string) [][]int {

	var markers [][]int
	for i, row := range rows {
		for j, cellValue := range row {
			if strings.Contains(cellValue, "Row") {
				markers = append(markers, []int{i + 1, j - 1})
			}
		}
	}
	return markers
}

// extractQ16Data extracts all the funds contributions and write them in the liquidityDataHolder.
// Same fund code is followed across Q15 and Q16
func extractQ16Data(f *excelize.File, ldataHolder *liquidityDataHolder) error {

	dataHolder := *ldataHolder

	Q16rows, err := f.GetRows("Q16")
	if err != nil {
		return fmt.Errorf("getting q16 rows: " + err.Error())
	}

	markers := getQ16Markers(Q16rows)

	for _, m := range markers {
		i, j := m[0], m[1]
		fund, _ := f.GetCellValue("Q16", executors.IndexToExcel(i-3, j+2))
		fundData := dataHolder[fund]
		for {
			cellValue, _ := f.GetCellValue("Q16", executors.IndexToExcel(i, j))
			if !strings.HasPrefix(cellValue, "16") {
				break
			}

			itype, _ := f.GetCellValue("Q16", executors.IndexToExcel(i, j+1))
			ivalue, _ := f.GetCellValue("Q16", executors.IndexToExcel(i, j+3))

			itype = strings.ToLower(itype)
			ivalue = executors.CleanNonNumeric(ivalue)

			percentage, _ := strconv.ParseFloat(ivalue, 64)

			switch {
			case strings.Contains(itype, "other"):
				fundData.otherPlanOwnedPercentage = percentage
			case strings.Contains(itype, "us"):
				fundData.usPersonOwnedPercentage = percentage
			case strings.Contains(itype, "pension"):
				fundData.pensionPlanOwnedPercentage = percentage
			case strings.Contains(itype, "non"):
				fundData.nonProfitOwnedPercentage = percentage
			}

			i += 1
		}

	}
	return nil

}

// ExtractInvestorLiquidity acts as fund agnostic extraction with the input investorLiquidity file.
// Stores all funds data and returns the reference of the holder
func ExtractInvestorLiquidity(filePath string) (*liquidityDataHolder, error) {

	ldataHolder := make(liquidityDataHolder)

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return &ldataHolder, fmt.Errorf("opening file: " + err.Error())
	}
	defer func() {
		if err := f.Close(); err != nil {
			logging.Error("Error in closing the file: " + err.Error())
		}

	}()

	q15Holder, err := extractQ15Data(f)
	if err != nil {
		return &ldataHolder, fmt.Errorf("extracing Q15: " + err.Error())
	}

	for fund, q15Value := range *q15Holder {
		ldataHolder[fund] = &LiquidityData{top5OwnedPercentage: q15Value}
	}

	err = extractQ16Data(f, &ldataHolder)
	if err != nil {
		return &ldataHolder, fmt.Errorf("extracting Q16: " + err.Error())
	}

	return &ldataHolder, nil
}
