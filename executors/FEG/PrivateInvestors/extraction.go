package privateinvestors

import (
	"fmt"
	"math"
	"quantrium-tech/form_pf_extractor_service/executors"
	"quantrium-tech/form_pf_extractor_service/logging"
	"slices"
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
func extractQ16Data(f *excelize.File, ldataHolder *LiquidityDataHolder) error {

	dataHolder := *ldataHolder

	Q16rows, err := f.GetRows("Q16")
	if err != nil {
		return fmt.Errorf("getting q16 rows: " + err.Error())
	}

	markers := getQ16Markers(Q16rows)

	for _, m := range markers {
		i, j := m[0], m[1]
		fund, _ := f.GetCellValue("Q16", executors.IndexToExcel(i-3, j+2))

		if !strings.HasPrefix(fund, "FEG") {
			logging.Warn("Fund Code not present in expected cell. Skipping the current table")
			continue
		}

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
func ExtractInvestorLiquidity(filePath string) (LiquidityDataHolder, error) {

	ldataHolder := make(LiquidityDataHolder)

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return ldataHolder, fmt.Errorf("opening file: " + err.Error())
	}
	defer func() {
		if err := f.Close(); err != nil {
			logging.Error("Error in closing the file: " + err.Error())
		}

	}()

	q15Holder, err := extractQ15Data(f)
	if err != nil {
		return ldataHolder, fmt.Errorf("extracing Q15: " + err.Error())
	}

	for fund, q15Value := range *q15Holder {
		ldataHolder[fund] = &LiquidityData{top5OwnedPercentage: q15Value}
	}

	err = extractQ16Data(f, &ldataHolder)
	if err != nil {
		return ldataHolder, fmt.Errorf("extracting Q16: " + err.Error())
	}

	return ldataHolder, nil
}

// extractQTD retrieves the required cells (unfunded commitments, gross and net performance)
// from the QTD Sheet of the Q3 Workpaper
func extractQTDData(f *excelize.File, series string, sheet string) (float64, float64, float64, QTDval, error) {

	var unfunded, net, gross, begBal, cashDistrib, capCall, netIncome, mgmtFee float64
	var missing []string
	var qtdVal QTDval

	qtdRows, err := f.GetRows(sheet)
	if err != nil {
		return unfunded, net, gross, qtdVal, fmt.Errorf("getting QTD Rows: " + err.Error())
	}

	var start, begBalIndex, capCallIndex, cashDistribIndex, netIncomeIndex, mgmtFeeIndex int //rowIndex
	var unfundedIndex, allPartnerIndex int
Outer:
	for i, row := range qtdRows {
		if len(row) >= 2 {
			for j, cellValue := range row {
				if strings.Contains(cellValue, "Series "+series) && strings.Contains(row[j-1], "Entity") {
					start = i
					break Outer
				}
			}
		}

	}
Out:
	for i, row := range qtdRows[start+1:] {
		for j, cellValue := range row {
			switch {
			case cashDistribIndex != 0 && capCallIndex != 0 && begBalIndex != 0 && allPartnerIndex != 0:
				break Out
			case strings.Contains(cellValue, "Unfunded Commitment") && unfundedIndex == 0:
				unfundedIndex = j
				continue
			case strings.Contains(cellValue, "All Partner") && allPartnerIndex == 0:
				allPartnerIndex = i
				continue
			case strings.EqualFold(cellValue, "Beginning Balance") && begBalIndex == 0:
				begBalIndex = j
				// rowIndex = i
				continue
			case strings.EqualFold(cellValue, "Investor - Capital Call") && capCallIndex == 0:
				capCallIndex = j
				// rowIndex = i
				continue
			case strings.EqualFold(cellValue, "Investor - Cash Distribution") && cashDistribIndex == 0:
				cashDistribIndex = j
				// rowIndex = i
				continue
			case strings.EqualFold(cellValue, "Net Income") && netIncomeIndex == 0:
				netIncomeIndex = j
				// rowIndex = i
				continue
			case (strings.EqualFold(cellValue, "Management Fee Expense") || strings.EqualFold(cellValue, "Management Fee Expenses")) && mgmtFeeIndex == 0:
				mgmtFeeIndex = j
				// rowIndex = i
				continue
			}
		}
	}

	for name, idx := range map[string]int{"unfundedIndex": unfundedIndex, "allPartnerIndex": allPartnerIndex, "begBalIndex": begBalIndex, "capCallIndex": capCallIndex, "cashDistribIndex": cashDistribIndex, "netIncomeIndex": netIncomeIndex, "mgmtFeeIndex": mgmtFeeIndex} {
		if idx == 0 {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		logging.Warn(fmt.Sprintf("Missing marker indexes:%v", strings.Join(missing, ", ")))
	}

	// rowIndex += (start + 1)
	allPartnerIndex += (start + 1)

	netIncomeValue, _ := f.GetCellValue(sheet, executors.IndexToExcel(allPartnerIndex, netIncomeIndex))
	mgmtFeeValue, _ := f.GetCellValue(sheet, executors.IndexToExcel(allPartnerIndex, mgmtFeeIndex))
	begBalValue, _ := f.GetCellValue(sheet, executors.IndexToExcel(allPartnerIndex, begBalIndex))
	capCallValue, _ := f.GetCellValue(sheet, executors.IndexToExcel(allPartnerIndex, capCallIndex))
	cashDistribValue, _ := f.GetCellValue(sheet, executors.IndexToExcel(allPartnerIndex, cashDistribIndex))
	unfundedValue, _ := f.GetCellValue(sheet, executors.IndexToExcel(allPartnerIndex, unfundedIndex))

	begBalValue = executors.CleanNonNumeric(begBalValue)
	capCallValue = executors.CleanNonNumeric(capCallValue)
	cashDistribValue = executors.CleanNonNumeric(cashDistribValue)
	netIncomeValue = executors.CleanNonNumeric(netIncomeValue)
	mgmtFeeValue = executors.CleanNonNumeric(mgmtFeeValue)
	unfundedValue = executors.CleanNonNumeric(unfundedValue)

	begBal, _ = strconv.ParseFloat(begBalValue, 64)
	capCall, _ = strconv.ParseFloat(capCallValue, 64)
	cashDistrib, _ = strconv.ParseFloat(cashDistribValue, 64)
	mgmtFee, _ = strconv.ParseFloat(mgmtFeeValue, 64)
	netIncome, _ = strconv.ParseFloat(netIncomeValue, 64)
	unfunded, _ = strconv.ParseFloat(unfundedValue, 64)

	qtdVal = QTDval{
		BegBal:      begBal,
		CashDistrib: cashDistrib,
		CapCall:     capCall,
		NetIncome:   netIncome,
		MgmtFee:     mgmtFee,
	}

	switch begBal + capCall + cashDistrib {
	case 0:
		logging.Warn("Sum of begining balance ,capitalCall ,cashDistribution is zero")
		net = 0
		gross = 0
	default:
		logging.Info("Sum of begining balance ,capitalCall ,cashDistribution is not equal to zero")
		dr := begBal + capCall + cashDistrib
		net = math.Round((netIncome/dr)*10000) / 100
		gross = math.Round(((netIncome-mgmtFee)/dr)*10000) / 100
	}

	return unfunded, net, gross, qtdVal, nil
}

// extractBSData extracts the necessary data from the BS Datasheet and updates in-place for the Fund WorkPaperData
func extractBSData(f *excelize.File, wpdata *WorkPaperData, series string) error {

	var seriesIndex int

	entitytoMarker := make(map[string]int)
	entitytoValue := make(map[string]float64)

	bsRows, err := f.GetRows("BS")
	if err != nil {
		return fmt.Errorf("getting BS Rows: " + err.Error())
	}
	for i, row := range bsRows {
	In:
		for j, cellValue := range row {
			switch {
			case seriesIndex != 0 && j > seriesIndex:
				break In
			case strings.Contains(cellValue, "Series "+series) && seriesIndex != 0:
				seriesIndex = j
			case strings.Contains(cellValue, "1002 - Cash"):
				entitytoMarker["cash1002"] = i
			case strings.Contains(cellValue, "1005 - Cash"):
				entitytoMarker["cash1005"] = i
			case strings.Contains(cellValue, "1000 - Cash"):
				entitytoMarker["cash1000"] = i
			case strings.Contains(cellValue, "1900 - Investment"):
				entitytoMarker["investment1900"] = i
			case strings.Contains(cellValue, "1910 - Investment"):
				entitytoMarker["investment1910"] = i
			case strings.EqualFold(cellValue, "Total Assets"):
				entitytoMarker["assets"] = i
			case strings.EqualFold(cellValue, "Total Liabilities"):
				entitytoMarker["liabilities"] = i
			case strings.EqualFold(cellValue, "Total Partner's Capital"):
				entitytoMarker["partnersCapital"] = i
			case strings.Contains(cellValue, strings.ToLower("Comerica Line of Credit Payable")):
				entitytoMarker["commericaLineOfCreditPayable"] = i
			}

		}
	}
	for entity, marker := range entitytoMarker {

		rawValue, _ := f.GetCellValue("BS", executors.IndexToExcel(marker, seriesIndex))
		rawValue = executors.CleanNonNumeric(rawValue)
		entitytoValue[entity], _ = strconv.ParseFloat(rawValue, 64)

	}

	wpdata.totalAssets = entitytoValue["assets"]
	wpdata.cashGeneralAccount = entitytoValue["cash1005"]
	wpdata.cashSweep = entitytoValue["cash1002"]
	wpdata.cashChecking = entitytoValue["cash1000"]
	wpdata.totalLiabilities = entitytoValue["liabilities"]
	wpdata.totalPartnerCaptial = entitytoValue["partnersCapital"]
	wpdata.investments = entitytoValue["investment1900"]
	wpdata.investmentsUnrealisedGainLoss = entitytoValue["investment1910"]
	if val, ok := entitytoValue["commericaLineOfCreditPayable"]; ok {
		wpdata.commericaLineOfCreditPayable = val
	}

	return nil
}

// extractOtherQuartersData retrieves the fund performance (gross and net) for other quarters (q1,q2) and updates
// inplace in the Fund's FundPerformanceData
func extractOtherQuartersData(fd *FundPerformanceData, f1Path string, f2Path string, series string, sheet string) error {
	var q1ExtVal, q2ExtVal QTDval
	f1, err := excelize.OpenFile(f1Path)
	if err != nil {
		return fmt.Errorf("opening q1 workpaper: " + err.Error())
	}
	f2, err := excelize.OpenFile(f2Path)
	if err != nil {
		return fmt.Errorf("opening q2 workpaper: " + err.Error())
	}
	defer func() {
		if err := f1.Close(); err != nil {
			logging.Error("Error in closing the q1 file: " + err.Error())
		}

		if err := f2.Close(); err != nil {
			logging.Error("Error in closing the q2 file: " + err.Error())
		}

	}()

	_, fd.q1NetPercentage, fd.q1GrossPercentage, q1ExtVal, err = extractQTDData(f1, series, sheet)
	if err != nil {
		return fmt.Errorf("extracing QTD Data for quarter 1: " + err.Error())
	}
	_, fd.q2NetPercentage, fd.q2GrossPercentage, q2ExtVal, err = extractQTDData(f2, series, sheet)
	if err != nil {
		return fmt.Errorf("extracing QTD Data for quarter 2: " + err.Error())
	}

	fd.begBalq2 = q2ExtVal.BegBal
	fd.capCallq2 = q2ExtVal.CapCall
	fd.cashDistribq2 = q2ExtVal.CashDistrib
	fd.mgmtFeeq2 = q2ExtVal.MgmtFee
	fd.netIncomeq2 = q2ExtVal.NetIncome

	fd.begBalq1 = q1ExtVal.BegBal
	fd.capCallq1 = q1ExtVal.CapCall
	fd.cashDistribq1 = q1ExtVal.CashDistrib
	fd.mgmtFeeq1 = q1ExtVal.MgmtFee
	fd.netIncomeq1 = q1ExtVal.NetIncome

	return nil
}

// ExtractWorkPaperData extracts the necessary data from the Workpapers of three quarters and store them in the
// WorkPaperData allocated to that particular fund
func (F *PrivateFund) ExtractWorkPaperData() error {

	f3, err := excelize.OpenFile(F.holder.q3FilePath)
	if err != nil {
		return fmt.Errorf("opening q3 workpaper: " + err.Error())
	}
	defer func() {
		if err := f3.Close(); err != nil {
			logging.Error("Error in closing the file: " + err.Error())
		}

	}()

	qtdSheet := "Capital Roll QTD"

	// Accomodate Special Case of Fund 3
	if strings.HasPrefix(F.fundID, "FEGOF_3") {
		qtdSheet = "Capital Roll (QTD)"
	}

	err = extractBSData(f3, &F.workPaperData, F.series)
	if err != nil {
		return fmt.Errorf("extracting BS Data: " + err.Error())
	}

	unfunded, net, gross, qtdExtVal, err := extractQTDData(f3, F.series, qtdSheet)
	if err != nil {
		return fmt.Errorf("extracting QTD Sheet: " + err.Error())
	}

	F.workPaperData.unfundedCommitments = unfunded
	if slices.Contains(MaturedFunds, F.fundID) {
		F.workPaperData.unfundedCommitments = 0
	}
	F.workPaperData.fundPerformanceData.begBalq3 = qtdExtVal.BegBal
	F.workPaperData.fundPerformanceData.capCallq3 = qtdExtVal.CapCall
	F.workPaperData.fundPerformanceData.cashDistribq3 = qtdExtVal.CashDistrib
	F.workPaperData.fundPerformanceData.mgmtFeeq3 = qtdExtVal.MgmtFee
	F.workPaperData.fundPerformanceData.netIncomeq3 = qtdExtVal.NetIncome
	F.workPaperData.fundPerformanceData.q3GrossPercentage = gross
	F.workPaperData.fundPerformanceData.q3NetPercentage = net

	err = extractOtherQuartersData(&F.workPaperData.fundPerformanceData, F.holder.q1FilePath, F.holder.q2FilePath, F.series, qtdSheet)
	if err != nil {
		return fmt.Errorf("extracing other quarter data: " + err.Error())
	}

	return nil
}
