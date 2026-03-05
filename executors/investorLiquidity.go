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
	cleanHeader := func(s string) string {
		s = strings.TrimSpace(s)                // trim spaces
		s = strings.TrimLeft(s, "\uFEFF")       // trim BOM (common on first cell)
		s = strings.ReplaceAll(s, "\r", "")     // trim carriage returns
		s = strings.ReplaceAll(s, "\u00a0", "") // trim non-breaking spaces
		return s
	}
	colFinder := func(col string) int {
		return slices.IndexFunc(rows[0], func(s string) bool { return strings.EqualFold(cleanHeader(s), col) })
	}

	fundCodeIdx := colFinder("FundCode")
	invTypeIdx := colFinder("InvestorType")
	totMvBaseIdx := colFinder("TotalMVBase")
	fmt.Println(rows[0])
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
		fundCode := strings.ToUpper(row[0])
		invType := strings.ToLower(row[1])
		val, _ := strconv.ParseFloat(CleanNonNumeric(row[2]), 64)
		if conv, ok := ShortCodeConvertor[strings.TrimSpace(fundCode)]; ok {
			fundCode = conv
		}
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
				Top5OwnedPercentage:             (top5Sum / total) * 100,
				UsPersonOwnedPercentage:         (invMap["uspersons"] / total) * 100,
				NotUsPersonOwnedPercentage:      (invMap["notuspersons"] / total) * 100,
				BrokerOwnedPercentage:           (invMap["broker"] / total) * 100,
				InsuranceOwnedPercentage:        (invMap["insurance"] / total) * 100,
				RiaOwnedPercentage:              (invMap["ria"] / total) * 100,
				PvtFundPercentage:               ((invMap["private fund"] + invMap["privatefund"]) / total) * 100,
				NonProfitOwnedPercentage:        (invMap["non-profit"] / total) * 100,
				PensionPlanOwnedPercentage:      (invMap["pensionplans"] / total) * 100,
				PropOwnedPercentage:             (invMap["proprietary"] / total) * 100,
				GovtOwnedPercentage:             (invMap["government"] / total) * 100,
				GovtPensionPlansOwnedPercentage: (invMap["governmentpensionplans"] / total) * 100,
				SovereignFundPercentage:         (invMap["sovereignfund"] / total) * 100,
				UnknownInstiOwnedPercentage:     (invMap["unknowninstitution"] / total) * 100,
				OtherPlanOwnedPercentage:        (invMap["other"] / total) * 100,
			}
		default:
			logging.Warn(fmt.Sprintf("TotalMVBase sum is zero for fund:%v", fundCode))
			(*ldataHolder)[fundCode] = &LiquidityData{
				Top5OwnedPercentage:             0,
				UsPersonOwnedPercentage:         0,
				NotUsPersonOwnedPercentage:      0,
				BrokerOwnedPercentage:           0,
				InsuranceOwnedPercentage:        0,
				RiaOwnedPercentage:              0,
				PvtFundPercentage:               0,
				NonProfitOwnedPercentage:        0,
				PensionPlanOwnedPercentage:      0,
				PropOwnedPercentage:             0,
				GovtOwnedPercentage:             0,
				GovtPensionPlansOwnedPercentage: 0,
				SovereignFundPercentage:         0,
				UnknownInstiOwnedPercentage:     0,
				OtherPlanOwnedPercentage:        0,
			}
		}
	}
}
