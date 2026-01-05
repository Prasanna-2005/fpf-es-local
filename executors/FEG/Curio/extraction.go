package curio

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
		"year to",
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
	fmt.Printf("\nnetmap:%v\n", netMap)
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
	case "private eq @ cost":
		return &ld.privateEqCost
	case "private eq u/r g/l":
		return &ld.privateEqURGL
	}
	return nil
}
