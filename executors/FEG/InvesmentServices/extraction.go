package invesmentservices

import (
	"fmt"
	"os"
	"quantrium-tech/form_pf_extractor_service/executors"
	"quantrium-tech/form_pf_extractor_service/logging"
	"slices"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

var EntityErr = "error occured for enity:%v extraction err: %v"

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

func ExtractUnfundedCmts(filePath string) (UnfundedCommitmentsHolder, error) {
	UfdataHolder := make(UnfundedCommitmentsHolder)

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return UfdataHolder, fmt.Errorf("opening file: " + err.Error())
	}
	defer func() {
		if err := f.Close(); err != nil {
			logging.Error("Error in closing the file: " + err.Error())
		}

	}()

	unfdCmtsHolder, err := extUnfndCmts(f)
	fmt.Println("UF cmmts")
	fmt.Println(unfdCmtsHolder)
	if err != nil {
		return unfdCmtsHolder, fmt.Errorf("extracing unfndcmts: " + err.Error())
	}
	UfdataHolder = unfdCmtsHolder
	return UfdataHolder, nil
}

func extUnfndCmts(f *excelize.File) (map[string]float64, error) {
	rows, err := f.GetRows(os.Getenv("UNFND_CMTS_SHEET"))
	fmt.Println(rows)
	if err != nil {
		return nil, fmt.Errorf("getting unfunded commitments rows: %w", err)
	}

	ymarkers := make(map[string]int)

	// Find column indices for BillingID and Unfunded Commitment
	for i, row := range rows {
		for j, cellValue := range row {
			lowerCellValue := strings.ToLower(cellValue)
			if strings.Contains(lowerCellValue, "billingid") {
				ymarkers["billid"] = j
				ymarkers["startx"] = i
			}
			if strings.Contains(lowerCellValue, "unfunded commitment") {
				ymarkers["unfdCmts"] = j
				ymarkers["startx"] = i
			}
		}
	}

	// Validate that we found both columns
	billIDCol, billIDOk := ymarkers["billid"]
	unfdCmtsCol, unfdCmtsOk := ymarkers["unfdCmts"]
	startRow, startRowOk := ymarkers["startx"]

	if !billIDOk || !unfdCmtsOk || !startRowOk {
		return nil, fmt.Errorf("could not find BillingID or Unfunded Commitment columns")
	}

	// Group unfunded commitments by billing ID
	result := make(map[string]float64)

	for i := startRow + 1; i < len(rows); i++ {
		row := rows[i]

		// Validate row has enough columns
		if len(row) <= billIDCol || len(row) <= unfdCmtsCol {
			continue
		}

		billID := strings.TrimSpace(row[billIDCol])
		unfdCmtStr := strings.TrimSpace(row[unfdCmtsCol])

		// Skip empty rows
		if billID == "" || unfdCmtStr == "" {
			break
		}

		// Parse unfunded commitment as float64
		unfdCmtStr = executors.CleanNonNumeric(unfdCmtStr)
		amount, err := strconv.ParseFloat(unfdCmtStr, 64)
		if err != nil {
			// Log warning and skip malformed entries
			logging.Warn(fmt.Sprintf("warning: skipping row %d - invalid amount: %s\n", i, unfdCmtStr))
			continue
		}

		// Add to existing sum or create new entry
		result[ShortCodeConvertor[billID]] += amount
	}

	return result, nil
}

// extractQ15Data extracts all the funds Q15 answer and returns them with fund mapped values
func extractQ15Data(f *excelize.File) (*map[string]float64, error) {

	q15Holder := make(map[string]float64)
	Q15rows, err := f.GetRows("Q15, 16")
	if err != nil {
		return &q15Holder, fmt.Errorf("getting q15 rows: " + err.Error())
	}

	markers := make(map[string]string)
	for i, row := range Q15rows {
		for j, cellValue := range row {
			if strings.Contains(cellValue, "Q15") {
				fundCol := executors.IndexToExcel(i-3, j-2)
				markers[fundCol] = executors.IndexToExcel(i, j-1)
			}
		}
	}
	for fundCol, location := range markers {
		fund, _ := f.GetCellValue("Q15, 16", fundCol)
		fund = ShortCodeConvertor[strings.TrimSpace(fund)]
		value, _ := f.GetCellValue("Q15, 16", location)
		fmt.Printf("\nfund:%vhivalue:%vhi\n", fund, value)
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
			if strings.Contains(cellValue, "Q16") {
				markers = append(markers, []int{i, j})
			}
		}
	}
	return markers
}

// extractQ16Data extracts all the funds contributions and write them in the liquidityDataHolder.
// Same fund code is followed across Q15 and Q16
func extractQ16Data(f *excelize.File, ldataHolder *LiquidityDataHolder) error {

	dataHolder := *ldataHolder

	Q16rows, err := f.GetRows("Q15, 16")
	if err != nil {
		return fmt.Errorf("getting q16 rows: " + err.Error())
	}

	markers := getQ16Markers(Q16rows)

	for _, m := range markers {
		i, j := m[0], m[1]
		fundMarker, _ := f.GetCellValue("Q15, 16", executors.IndexToExcel(i-3, j-4))

		if strings.Compare(fundMarker, "FundCode") != 0 {
			logging.Warn("Fund Code not present in expected cell. Skipping the current table")
			continue
		}
		fund, _ := f.GetCellValue("Q15, 16", executors.IndexToExcel(i-3, j-3))
		fund = ShortCodeConvertor[strings.TrimSpace(fund)]
		fundData := dataHolder[fund]
		for {
			cellValue, _ := f.GetCellValue("Q15, 16", executors.IndexToExcel(i, j-4))
			if strings.Contains(cellValue, "Grand Total") {
				break
			}

			itype, _ := f.GetCellValue("Q15, 16", executors.IndexToExcel(i, j-4))
			ivalue, _ := f.GetCellValue("Q15, 16", executors.IndexToExcel(i, j-2))

			itype = strings.TrimSpace(strings.ToLower(itype))
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
			case strings.Contains(itype, "private"):
				fundData.pvtFundPercentage = percentage
			case strings.Contains(itype, "sovereign"):
				fundData.sovereignFundPercentage = percentage
			}

			i += 1
		}

	}
	return nil

}
func (F *Fund) ExtractQ17() error {
	fNAVReport, err := excelize.OpenFile(F.holder.reportFilePath)
	if err != nil {
		return fmt.Errorf("opening  NAV reports for q17 data: " + err.Error())
	}
	defer func() {
		if err := fNAVReport.Close(); err != nil {
			logging.Error("Error in closing the file: " + err.Error())
		}
	}()
	err = extracNAVRepData(fNAVReport, &F.workPaperData)
	if err != nil {
		return fmt.Errorf("extracting Q17 data: " + err.Error())
	}
	return nil
}

// ExtractWorkPaperData extracts the necessary data from the Workpapers of three quarters and store them in the
// WorkPaperData allocated to that particular fund
func (F *Fund) ExtractTrialBalanceData() error {

	fTrialBalance, err := excelize.OpenFile(F.holder.tbFilePath)
	if err != nil {
		return fmt.Errorf("opening  Trial balance: " + err.Error())
	}
	defer func() {
		if err := fTrialBalance.Close(); err != nil {
			logging.Error("Error in closing the file: " + err.Error())
		}

	}()

	ymarker, err := extGrossAndNetAssests(fTrialBalance, &F.workPaperData)
	if err != nil {
		return fmt.Errorf("extracting TB,Unfunded cmts Data for GA NA: " + err.Error())
	}

	err = extLevelingData(ymarker, fTrialBalance, &F.workPaperData, F)
	if err != nil {
		return fmt.Errorf("extracting TB Data for leveling: " + err.Error())
	}

	return nil
}

func extLevelingData(ordinate int, f *excelize.File, tbdata *WorkPaperData, F *Fund) error {
	tbSheet := os.Getenv("TB_SHEET")
	markerl1 := make(map[string]int)
	markerl2 := make(map[string]int)
	markerl3 := make(map[string]int)
	l1, _ := LevelMapper[F.fundID]["l1"]
	l2, _ := LevelMapper[F.fundID]["l2"]
	l3, _ := LevelMapper[F.fundID]["l3"]

	tbRows, err := f.GetRows(tbSheet)
	if err != nil {
		return fmt.Errorf("getting TB Rows: " + err.Error())
	}
	for i, row := range tbRows {
		for _, cellValue := range row {
			cellval := strings.ToLower(cellValue)
			if slices.Contains(l1, cellval) {
				markerl1[LevelingEntityMapper[cellval]] = i
			}
			if slices.Contains(l2, cellval) {
				markerl2[LevelingEntityMapper[cellval]] = i
			}
			if slices.Contains(l3, cellval) {
				markerl3[LevelingEntityMapper[cellval]] = i
			}

		}
	}

	for marker, abscissa := range markerl1 {
		rawValue, err := f.GetCellValue(tbSheet, executors.IndexToExcel(abscissa, ordinate))
		if err != nil {
			logging.Error(fmt.Sprintf(EntityErr, marker, err.Error()))
		}
		rawValue = executors.CleanNonNumeric(rawValue)
		l1, _ := strconv.ParseFloat(rawValue, 64)
		if fp := FieldPtr(&tbdata.levelExtData, marker); fp != nil {
			*fp = l1
		}
		tbdata.l1 += l1
		tbdata.sumLevelExtData += l1
	}
	for marker, abscissa := range markerl2 {
		rawValue, err := f.GetCellValue(tbSheet, executors.IndexToExcel(abscissa, ordinate))
		if err != nil {
			logging.Error(fmt.Sprintf(EntityErr, marker, err.Error()))
		}
		rawValue = executors.CleanNonNumeric(rawValue)
		l2, _ := strconv.ParseFloat(rawValue, 64)
		if fp := FieldPtr(&tbdata.levelExtData, marker); fp != nil {
			*fp = l2
		}

		fmt.Println("Hi")
		tbdata.l2 += l2
		tbdata.sumLevelExtData += l2
	}
	for marker, abscissa := range markerl3 {
		rawValue, err := f.GetCellValue(tbSheet, executors.IndexToExcel(abscissa, ordinate))
		if err != nil {
			logging.Error(fmt.Sprintf(EntityErr, marker, err.Error()))
		}
		rawValue = executors.CleanNonNumeric(rawValue)
		l3, _ := strconv.ParseFloat(rawValue, 64)
		if fp := FieldPtr(&tbdata.levelExtData, marker); fp != nil {
			*fp = l3
		}
		tbdata.l3 += l3
		tbdata.sumLevelExtData += l3
	}
	return nil
}

func extGrossAndNetAssests(f *excelize.File, tbdata *WorkPaperData) (int, error) {
	yEntitytoMarker := make(map[string]int)
	tbSheet := os.Getenv("TB_SHEET")

	tbRows, err := f.GetRows(tbSheet)
	if err != nil {
		return -1, fmt.Errorf("getting TB Rows: " + err.Error())
	}
	for i, row := range tbRows {
		for j, cellValue := range row {
			cellval := strings.ToLower(cellValue)
			switch {
			case strings.Contains(cellval, "ending balance"):
				yEntitytoMarker["ed"] = j
			case strings.Contains(cellval, "total liabilities"):
				rawValue, err := f.GetCellValue(tbSheet, executors.IndexToExcel(i, yEntitytoMarker["ed"]))
				if err != nil {
					logging.Error(fmt.Sprintf(EntityErr, cellval, err.Error()))
				}
				rawValue = executors.CleanNonNumeric(rawValue)
				tbdata.totalLiabilities, _ = strconv.ParseFloat(rawValue, 64)
			case strings.Contains(cellval, "total assets"):
				rawValue, err := f.GetCellValue(tbSheet, executors.IndexToExcel(i, yEntitytoMarker["ed"]))
				if err != nil {
					logging.Error(fmt.Sprintf(EntityErr, cellval, err.Error()))
				}
				rawValue = executors.CleanNonNumeric(rawValue)
				tbdata.totalAssets, _ = strconv.ParseFloat(rawValue, 64)
			case strings.Contains(cellval, "total capital"):
				rawValue, err := f.GetCellValue(tbSheet, executors.IndexToExcel(i, yEntitytoMarker["ed"]))
				if err != nil {
					logging.Error(fmt.Sprintf(EntityErr, cellval, err.Error()))
				}
				rawValue = executors.CleanNonNumeric(rawValue)
				tbdata.totalPartnerCaptial, _ = strconv.ParseFloat(rawValue, 64)
				if tbdata.totalPartnerCaptial < 0 {
					tbdata.totalPartnerCaptial *= (-1)
				}
			}
			tbdata.unfundedCommitments = 0
		}
	}
	return yEntitytoMarker["ed"], nil
}

func extracNAVRepData(f *excelize.File, Navdata *WorkPaperData) error {
	netMap := make(map[string]float64)
	grossMap := make(map[string]float64)
	absicca := make(map[string]int)
	ordinate := make(map[string]int)
	xmarkers := []string{
		"january", "february", "march", "first",
		"april", "may", "june", "second",
		"july", "august", "september", "third",
		"october", "november", "december", "fourth",
		"year",
	}
	ymarkers := []string{
		"net of mgmt", "gross of fees",
	}
	reportSheet := "Page 1"
	reportRows, err := f.GetRows(reportSheet)
	if err != nil {
		return fmt.Errorf("getting NAV report rows: " + err.Error())
	}

	for i, row := range reportRows {
		for j, cellValue := range row {
			cellval := strings.ToLower(cellValue)
			switch {
			case slices.Contains(xmarkers, cellval):
				absicca[cellval] = i
			case slices.Contains(ymarkers, cellval):
				ordinate[cellval] = j
			}
		}
	}

	for term, xcoor := range absicca {

		rawValue, err := f.GetCellValue(reportSheet, executors.IndexToExcel(xcoor, ordinate["net of mgmt"]))
		if err != nil {
			logging.Error(fmt.Sprintf("excel error: %v", err.Error()))
		}
		rawValue = executors.CleanNonNumeric(rawValue)
		netPercent, _ := strconv.ParseFloat(rawValue, 64)
		if netPercent != 0 {
			netMap[term] = netPercent
		}
		rawValue, err = f.GetCellValue(reportSheet, executors.IndexToExcel(xcoor, ordinate["gross of fees"]))
		if err != nil {
			logging.Error(fmt.Sprintf("excel error: %v", err.Error()))
		}
		rawValue = executors.CleanNonNumeric(rawValue)
		grossPercent, _ := strconv.ParseFloat(rawValue, 64)
		if grossPercent != 0 {
			grossMap[term] = grossPercent
		}

	}
	fmt.Println("inside extraction")
	fmt.Printf("\ngmap:%v\n", grossMap)
	assignPerformance(grossMap, Navdata, true)
	fmt.Println("\nnetmap:%v\n", netMap)
	assignPerformance(netMap, Navdata, false)
	fmt.Printf("\nnavdata:%v\n", Navdata.fundPerformanceData)
	return nil
}

func assignPerformance(src map[string]float64, Navdata *WorkPaperData, isGross bool) {
	fp := &Navdata.fundPerformanceData
	xmarkers := []string{
		"january", "february", "march", "first",
		"april", "may", "june", "second",
		"july", "august", "september", "third",
		"october", "november", "december", "fourth",
		"year",
	}
	// iterate in the exact order you defined in xmarkers
	for _, key := range xmarkers {
		val, ok := src[key]
		if !ok || val == 0 { // skip missing / zero
			continue
		}

		switch key {
		// months
		case "january":
			if isGross {
				fp.janGrossPercentage = val
			} else {
				fp.janNetPercentage = val
			}
		case "february":
			if isGross {
				fp.febGrossPercentage = val
			} else {
				fp.febNetPercentage = val
			}
		case "march":
			if isGross {
				fp.marGrossPercentage = val
			} else {
				fp.marNetPercentage = val
			}
		case "april":
			if isGross {
				fp.aprGrossPercentage = val
			} else {
				fp.aprNetPercentage = val
			}
		case "may":
			if isGross {
				fp.mayGrossPercentage = val
			} else {
				fp.mayNetPercentage = val
			}
		case "june":
			if isGross {
				fp.junGrossPercentage = val
			} else {
				fp.junNetPercentage = val
			}
		case "july":
			if isGross {
				fp.julGrossPercentage = val
			} else {
				fp.julNetPercentage = val
			}
		case "august":
			if isGross {
				fp.augGrossPercentage = val
			} else {
				fp.augNetPercentage = val
			}
		case "september":
			if isGross {
				fp.sepGrossPercentage = val
			} else {
				fp.sepNetPercentage = val
			}
		case "october":
			if isGross {
				fp.octGrossPercentage = val
			} else {
				fp.octNetPercentage = val
			}
		case "november":
			if isGross {
				fp.novGrossPercentage = val
			} else {
				fp.novNetPercentage = val
			}
		case "december":
			if isGross {
				fp.decGrossPercentage = val
			} else {
				fp.decNetPercentage = val
			}

		// quarters
		case "first":
			if isGross {
				fp.q1GrossPercentage = val
			} else {
				fp.q1NetPercentage = val
			}
		case "second":
			if isGross {
				fp.q2GrossPercentage = val
			} else {
				fp.q2NetPercentage = val
			}
		case "third":
			if isGross {
				fp.q3GrossPercentage = val
			} else {
				fp.q3NetPercentage = val
			}
		case "fourth":
			if isGross {
				fp.q4GrossPercentage = val
			} else {
				fp.q4NetPercentage = val
			}
		case "year":
			if isGross {
				fp.yearGrossPercentage = val
			} else {
				fp.yearNetPercentage = val
			}

		}
	}
}

func FieldPtr(ld *LevelingData, marker string) *float64 {
	switch marker {
	case "adr @ cost":
		return &ld.adrCost
	case "adr u/r g/l":
		return &ld.adrURGL
	case "corp bonds @ cost":
		return &ld.corpBondsCost
	case "corp bonds u/r g/l":
		return &ld.corpBondsURGL
	case "direct holding @ cost":
		return &ld.directHoldingCost
	case "direct holding u/r g/l":
		return &ld.directHoldingURGL
	case "equities @ cost":
		return &ld.equitiesCost
	case "equities dividend rec":
		return &ld.equitiesDividendRec
	case "equities ft w/h payable":
		return &ld.equitiesFtWhPayable
	case "equities u/r fx g/l":
		return &ld.equitiesURFXGL
	case "equities u/r g/l":
		return &ld.equitiesURGL
	case "gdr @ cost":
		return &ld.gdrCost
	case "gdr u/r g/l":
		return &ld.gdrURGL
	case "govt bonds @ cost":
		return &ld.govtBondsCost
	case "govt bonds u/r g/l":
		return &ld.govtBondsURGL
	case "invt hedge funds @ cost":
		return &ld.invtHedgeFundsCost
	case "invt in hedge funds u/r":
		return &ld.invtInHedgeFundsUR
	case "mortgage & cmo @ cost":
		return &ld.mortgageCmoCost
	case "mortgage & cmo u/r g/l":
		return &ld.mortgageCmoURGL
	case "private eq @ cost":
		return &ld.privateEqCost
	case "private eq u/r g/l":
		return &ld.privateEqURGL
	case "reit @ cost":
		return &ld.reitCost
	case "reit u/r g/l":
		return &ld.reitURGL
	case "treasuries @ cost":
		return &ld.treasuriesCost
	case "treasuries u/r g/l":
		return &ld.treasuriesURGL
	case "invt in master fund":
		return &ld.invtInMasterFund
	}
	return nil
}
