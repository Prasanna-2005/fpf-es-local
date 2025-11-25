package invesmentservices

import (
	"quantrium-tech/form_pf_extractor_service/executors"
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
