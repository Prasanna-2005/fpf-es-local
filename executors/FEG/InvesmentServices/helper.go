package invesmentservices

import (
	"database/sql"
	"fmt"
	"quantrium-tech/form_pf_extractor_service/executors"
	"reflect"
	"time"
)

// CreateAllFundSeries is a helper function that takes up a fund and creates all of the fund-series combination and returns them
func CreateAllFundSeries(fundCode string) []*Fund {

	var funds []*Fund

	fundId := fundCode

	fund := NewFund(fundId)
	funds = append(funds, fund)

	return funds
}

// func IndexToValue(tbSheet, yEntitytoMarker,  i int,cellval) (float64, error) {
// 	rawValue, err := f.GetCellValue(tbSheet, executors.IndexToExcel(i, yEntitytoMarker["ed"]))
// 	if err != nil {
// 		logging.Error(fmt.Sprintf(EntityErr, cellval, err.Error()))
// 	}
// 	rawValue = executors.CleanNonNumeric(rawValue)
// 	realValue, err := strconv.ParseFloat(rawValue, 64)
// 	if err != nil {
// 		logging.Error(fmt.Sprintf("string to floeat conversion err for entity %v: %v",cellval,err.Error()))
// 	}
// 	return realValue, nil
// }

func StaticFunds() (*executors.Form1bData, *Fund) {
	fundData := executors.NewForm1bData()
	fundData.OtherPercent = 100
	fundObj := NewFund("FEGSLLC")

	fundObj.fundID = "FSSLC"
	fundObj.series = "FSSLC"

	return fundData, fundObj

}

// assignFundPerf copies every float64/string from src (FundPerformance) into dst
// (Form1bData) converting to sql.NullFloat64/sql.NullString when {term,colName} matches.
func assignFundPerf(dst *executors.Form1bData, src FundPerformanceData) {
	vDst := reflect.ValueOf(dst).Elem()
	tDst := vDst.Type()
	vSrc := reflect.ValueOf(src)

	type keyT [2]string
	srcMap := make(map[keyT]reflect.Value)
	for i := 0; i < vSrc.NumField(); i++ {
		f := vSrc.Type().Field(i)
		term := f.Tag.Get("term")
		col := f.Tag.Get("colName")
		if term != "" && col != "" {
			srcMap[keyT{term, col}] = vSrc.Field(i)
		}
	}

	// 1. map  term → index of date field
	dateIdx := make(map[string]int)
	for i := 0; i < tDst.NumField(); i++ {
		f := tDst.Field(i)
		if f.Tag.Get("colName") == "date" {
			dateIdx[f.Tag.Get("term")] = i
		}
	}

	nonZero := func(v reflect.Value) bool {
		return v.IsValid() && v.Kind() == reflect.Float64 && v.Float() != 0
	}

	for i := 0; i < tDst.NumField(); i++ {
		f := tDst.Field(i)
		term := f.Tag.Get("term")
		col := f.Tag.Get("colName")
		if term == "" || col == "" {
			continue
		}

		dstFld := vDst.Field(i)

		switch dstFld.Type() {
		case reflect.TypeOf(sql.NullFloat64{}):
			src, ok := srcMap[keyT{term, col}]
			if !ok {
				continue
			}
			dstFld.Set(reflect.ValueOf(sql.NullFloat64{
				Float64: src.Float(),
				Valid:   src.Float() != 0,
			}))

			// 2. single-shot date assignment
			fmt.Println(dateIdx)
			if idx, ok := dateIdx[term]; ok {
				otherCol := "gross"
				if col == "gross" {
					otherCol = "net"
				}
				other, _ := srcMap[keyT{term, otherCol}]
				if nonZero(src) || nonZero(other) {
					date := lastDayOfTerm(term, time.Now().Year()-1)
					vDst.Field(idx).Set(reflect.ValueOf(sql.NullString{
						String: date,
						Valid:  true,
					}))
				}
			}
		}
	}
}

// lastDayOfTerm returns the last day of the month / quarter / year
// for the *current* calendar year (or the just-ended year when term=="year").
// You can change `reportYear` if your fiscal year is different.
func lastDayOfTerm(term string, year int) string {
	switch term {
	// months
	case "january":
		return fmt.Sprintf("%04d-01-31", year)
	case "february":
		// last day of Feb = 1-Mar minus 1 day
		return time.Date(year, 3, 0, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
	case "march":
		return fmt.Sprintf("%04d-03-31", year)
	case "april":
		return fmt.Sprintf("%04d-04-30", year)
	case "may":
		return fmt.Sprintf("%04d-05-31", year)
	case "june":
		return fmt.Sprintf("%04d-06-30", year)
	case "july":
		return fmt.Sprintf("%04d-07-31", year)
	case "august":
		return fmt.Sprintf("%04d-08-31", year)
	case "september":
		return fmt.Sprintf("%04d-09-30", year)
	case "october":
		return fmt.Sprintf("%04d-10-31", year)
	case "november":
		return fmt.Sprintf("%04d-11-30", year)
	case "december":
		return fmt.Sprintf("%04d-12-31", year)

	// quarters
	case "first":
		return fmt.Sprintf("%04d-03-31", year)
	case "second":
		return fmt.Sprintf("%04d-06-30", year)
	case "third":
		return fmt.Sprintf("%04d-09-30", year)
	case "fourth":
		return fmt.Sprintf("%04d-12-31", year)

	// full year
	case "year":
		return fmt.Sprintf("%04d-12-31", year)

	default:
		return "" // safeguard
	}
}
