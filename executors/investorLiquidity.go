package executors

import (
	"fmt"
	"quantrium-tech/form_pf_extractor_service/logging"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func ExtractInvestorLiquidity(filePath string) (InvLiquidityData, error) {

	ldataHolder := make(InvLiquidityData)

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return ldataHolder, fmt.Errorf("opening file: " + err.Error())
	}
	defer func() {
		if err := f.Close(); err != nil {
			logging.Error("Error in closing the file: " + err.Error())
		}

	}()

	data, err := cleanExcel(f)
	if err != nil || data == nil {
		logging.Error(fmt.Sprintf("excel cleaning Error: %v", err.Error()))
	}

	processFundData(data, &ldataHolder)

	return ldataHolder, nil
}

// colnames : FundCode InvestorType TotalMVBase
func cleanExcel(f *excelize.File) ([][]string, error) {

	var fundLiqData [][]string
	rows, err := f.GetRows("InvestorLiquidity")
	if err != nil {
		return fundLiqData, fmt.Errorf("getting q15 rows: " + err.Error())
	}
	if len(rows) < 2 || len(rows[0]) < 2 {
		return fundLiqData, fmt.Errorf("Excel doesnt have proper number of rows or columns")
	}
	colFinder := func(col string) int {
		return slices.IndexFunc(rows[0], func(s string) bool { return strings.EqualFold(s, col) })
	}

	fundCodeIdx := colFinder("FundCode")
	invTypeIdx := colFinder("InvestorType")
	totMvBaseIdx := colFinder("TotalMVBase")
	if fundCodeIdx == -1 || invTypeIdx == -1 || totMvBaseIdx == -1 {
		return fundLiqData, fmt.Errorf("Excel doesnt have required column names")
	}

	for i, row := range rows[1:] {
		if len(row) < 3 {
			logging.Warn(fmt.Sprintf("InvestorLiqSheet: %v row doesn't have required number  of columns", i))
			continue
		}
		fundLiqData = append(fundLiqData, []string{strings.ToLower(row[fundCodeIdx]), strings.ToLower(row[invTypeIdx]), row[totMvBaseIdx]})
		// val, _ := strconv.ParseFloat(CleanNonNumeric(row[totMvBaseIdx]), 64)

	}

	return fundLiqData, nil
}

func processFundData(fundData [][]string, ldataHolder *InvLiquidityData) {
	grouped := make(map[string][]float64)
	nestedMap := make(map[string]map[string]float64)

	for _, row := range fundData {
		if len(row) < 3 {
			continue
		}
		fundCode := strings.ToLower(row[0])
		invType := strings.ToLower(row[1])
		val, _ := strconv.ParseFloat(row[2], 64)

		if nestedMap[fundCode] == nil {
			nestedMap[fundCode] = make(map[string]float64)
		}
		nestedMap[fundCode][invType] += val
		grouped[fundCode] = append(grouped[fundCode], val)
	}

	for fundCode, values := range grouped {
		var top5Sum, total float64
		sort.Slice(values, func(i, j int) bool { return values[i] > values[j] })

		for i, val := range values {
			if i < 5 {
				top5Sum += val
			}
			total += val
		}

		invMap := nestedMap[fundCode]

		switch {
		case total != 0:
			(*ldataHolder)[fundCode] = &LiquidityData{
				top5OwnedPercentage:             (top5Sum / total) * 100,
				usPersonOwnedPercentage:         (invMap["uspersons"] / total) * 100,
				notUsPersonOwnedPercentage:      (invMap["notuspersons"] / total) * 100,
				brokerOwnedPercentage:           (invMap["broker"] / total) * 100,
				insuranceOwnedPercentage:        (invMap["insurance"] / total) * 100,
				riaOwnedPercentage:              (invMap["ria"] / total) * 100,
				pvtFundPercentage:               ((invMap["private fund"] + invMap["privatefund"]) / total) * 100,
				nonProfitOwnedPercentage:        (invMap["non-profit"] / total) * 100,
				pensionPlanOwnedPercentage:      (invMap["pensionplans"] / total) * 100,
				propOwnedPercentage:             (invMap["proprietary"] / total) * 100,
				govtOwnedPercentage:             (invMap["government"] / total) * 100,
				govtPensionPlansOwnedPercentage: (invMap["governmentpensionplans"] / total) * 100,
				sovereignFundPercentage:         (invMap["sovereignfund"] / total) * 100,
				unknownInstiOwnedPercentage:     (invMap["unknowninstitution"] / total) * 100,
				otherPlanOwnedPercentage:        (invMap["other"] / total) * 100,
			}
		default:
			logging.Warn(fmt.Sprintf("TotalMVBase sum is zero for fund:%v", fundCode))
			(*ldataHolder)[fundCode] = &LiquidityData{
				top5OwnedPercentage:             0,
				usPersonOwnedPercentage:         0,
				notUsPersonOwnedPercentage:      0,
				brokerOwnedPercentage:           0,
				insuranceOwnedPercentage:        0,
				riaOwnedPercentage:              0,
				pvtFundPercentage:               0,
				nonProfitOwnedPercentage:        0,
				pensionPlanOwnedPercentage:      0,
				propOwnedPercentage:             0,
				govtOwnedPercentage:             0,
				govtPensionPlansOwnedPercentage: 0,
				sovereignFundPercentage:         0,
				unknownInstiOwnedPercentage:     0,
				otherPlanOwnedPercentage:        0,
			}
		}
	}
}
