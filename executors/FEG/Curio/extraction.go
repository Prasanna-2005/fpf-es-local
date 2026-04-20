package curio

import (
	"fmt"
	"math"
	"os"
	"quantrium-tech/form_pf_extractor_service/executors"
	"quantrium-tech/form_pf_extractor_service/logging"
	"slices"
	"strconv"
	"strings"

	"github.com/shakinm/xlsReader/xls"
	"github.com/xuri/excelize/v2"
)

var EntityErr = "error occured for enity:%v extraction err: %v"

var singular = map[string]string{

	"distributions":   "distribution",
	"contributions":   "contribution",
	"management fees": "management fee",
}

func (F *Fund) ExtractFundPerformance() error {
	F.workPaperData.irr = make(map[string]*IRR)
	for period, path := range F.holder.capRegisterFilePath {
		wb, err := xls.OpenFile(path)
		if err != nil {
			return fmt.Errorf("opening  NAV reports for capital register data: %v" + err.Error())
		}
		fNAVReport, err := wb.GetSheet(0)
		if err != nil {
			return fmt.Errorf("Sheet error in capital register:%v", err.Error())
		}
		err = F.extracNAVRepData(fNAVReport, period)
		if err != nil {
			return fmt.Errorf("extracting fund performance data: " + err.Error())
		}
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

func (F *Fund) ExtractCapReg() error {
	path, ok := F.holder.capRegisterFilePath["september"]
	if !ok {
		path, ok = F.holder.capRegisterFilePath["third"]
	}

	if !ok {
		logging.Error("File not found: September month capital register file not found.")
		return fmt.Errorf("Capital register not found for september")
	}

	wb, err := xls.OpenFile(path)
	if err != nil {
		return fmt.Errorf("opening capital register: %w", err)
	}
	err = extractUnfundedCmts(wb, &F.workPaperData, F.fundID)
	if err != nil {
		return fmt.Errorf("extracting TB,Unfunded cmts Data for GA NA: " + err.Error())
	}
	return nil
}

func extractUnfundedCmts(wb xls.Workbook, tbdata *WorkPaperData, shortcode string) error {
	shortcode = strings.ToUpper(shortcode)
	ordinates := make(map[string]int)
	absicca := make(map[string]int)

	sheet, err := wb.GetSheet(0)
	if err != nil {
		return fmt.Errorf("no sheet found in capital register: %w", err)
	}

	for i := 0; i < sheet.GetNumberRows(); i++ {
		row, err := sheet.GetRow(i)
		if err != nil {
			continue
		}
		for j := 0; j < len(row.GetCols()); j++ {
			cell, err := row.GetCol(j)
			if err != nil {
				continue
			}
			cellVal := cell.GetString()
			switch {
			case strings.EqualFold("Unfunded commitment", cellVal):
				ordinates["unfdCmts"] = j
			case strings.EqualFold("id", cellVal):
				ordinates["id"] = j
			case strings.EqualFold(shortcode, cellVal) && j == ordinates["id"]:
				absicca[shortcode] = i
			}
		}
	}

	targetRow, rowFound := absicca[shortcode]
	targetCol, colFound := ordinates["unfdCmts"]
	if !rowFound || !colFound {
		logging.Error(fmt.Sprintf("could not find (row,col)(%v,%v) for shortcode:%v", shortcode, rowFound, colFound))
		return nil
	}

	reportRows := GetRows(sheet)

	cell := reportRows[targetRow][targetCol]
	cell = executors.CleanNonNumeric(cell)
	// GetFloat64() preserves decimals correctly
	tbdata.unfundedCommitments, _ = strconv.ParseFloat(cell, 64)
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
			case strings.Contains(cellval, strings.ToLower("Comerica Line of Credit Payable")):
				rawValue, err := f.GetCellValue(tbSheet, executors.IndexToExcel(i, yEntitytoMarker["ed"]))
				if err != nil {
					logging.Error(fmt.Sprintf(EntityErr, cellval, err.Error()))
				}
				rawValue = executors.CleanNonNumeric(rawValue)
				tbdata.commericaLineOfCreditPayable, _ = strconv.ParseFloat(rawValue, 64)

			}
			tbdata.unfundedCommitments = 0
		}
	}
	return yEntitytoMarker["ed"], nil
}

func (F *Fund) extracNAVRepData(f *xls.Sheet, period string) error {
	var totalCapCall float64
	absicca := make(map[string]int)
	ordinate := make(map[string]int)
	F.workPaperData.irr[period] = &IRR{
		capReg: make(map[string]float64),
	}
	ymarkers := []string{
		"open net balance", "contributions", "distributions",
		"grosspnl", "management fee", "management fees", "contribution", "distribution",
	}
	reportRows := GetRows(f)
	if len(reportRows) == 0 {
		return fmt.Errorf("empty capital register found in NAV extraction")
	}

	for i, row := range reportRows {
		for j, cellValue := range row {
			cellval := strings.ToLower(cellValue)
			switch {
			case strings.EqualFold("id", cellval):
				absicca["id"] = i
				ordinate["id"] = j
			case strings.EqualFold(F.fundID, cellval):
				if ordinate["id"] != j {
					// logging.Warn(fmt.Sprintf("'id' and %s (%s)  extracted are not in same column", cellval, F.fundID))
					continue
				}
				absicca[F.fundID] = i
			case slices.Contains(ymarkers, cellval):
				if absicca["id"] != i {
					// logging.Warn(fmt.Sprintf("'id' and %s  extracted are not in same row", cellval))
					continue
				}
				if m, ok := singular[cellval]; ok {
					cellval = m
				}
				ordinate[cellval] = j
			case strings.Contains(cellval, "capital call"):
				if F.fundID == "CURHP" && !containsKey(ordinate, cellval) && absicca["id"] == i {
					ordinate[cellval] = j
				}
			}
		}
	}
	fundRowIdx, ok := absicca[F.fundID]
	if !ok {
		return fmt.Errorf("fundID %s not found in sheet", F.fundID)
	}
	for marker, ycoor := range ordinate {
		if marker == "id" {
			continue
		}
		cellVal := executors.CleanNonNumeric(reportRows[fundRowIdx][ycoor])
		cell, _ := strconv.ParseFloat(cellVal, 64)
		if _, ok := F.workPaperData.irr[period].capReg[marker]; !ok {
			if strings.Contains(marker, "capital call") {
				totalCapCall += math.Abs(cell)
			}
			F.workPaperData.irr[period].capReg[marker] = cell
		}
	}
	dr := F.workPaperData.irr[period].capReg["open net balance"] + F.workPaperData.irr[period].capReg["contributions"] + F.workPaperData.irr[period].capReg["distributions"] + totalCapCall
	if dr == 0 {
		logging.Warn(fmt.Sprintf("Fund performance denominator is zero for period:%v", period))
		F.workPaperData.irr[period].grossReturns = 0
		F.workPaperData.irr[period].netReturns = 0
	} else {
		grossRet := F.workPaperData.irr[period].capReg["grosspnl"] * 100 / dr
		netRet := (F.workPaperData.irr[period].capReg["grosspnl"] - math.Abs(F.workPaperData.irr[period].capReg["management fee"])) * 100 / dr
		F.workPaperData.irr[period].grossReturns = math.Round(grossRet*100) / 100
		F.workPaperData.irr[period].netReturns = math.Round(netRet*100) / 100
	}
	return nil
}

func assignPerformance(src map[string]float64, Navdata *WorkPaperData, isGross bool) {
	fp := &Navdata.fundPerformanceData
	xmarkers := []string{
		"january", "february", "march", "first",
		"april", "may", "june", "second",
		"july", "august", "september", "third",
		"october", "november", "december", "fourth",
		"year to",
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
		case "year to":
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
	case "invt hedge funds @ cost":
		return &ld.invtHedgeFundsCost
	case "invt in hedge funds u/r":
		return &ld.invtInHedgeFundsUR
	case "invt in hedge funds u/r g/l":
		return &ld.invtInHedgeFundsURGL
	case "private eq @ cost":
		return &ld.privateEqCost
	case "private eq u/r g/l":
		return &ld.privateEqURGL
	}
	return nil
}
