package privateinvestors

import (
	"strings"
)

// CreateAllFundSeries is a helper function that takes up a fund and creates all of the fund-series combination and returns them
func CreateAllFundSeries(fundCode string) []*PrivateFund {

	var funds []*PrivateFund

	seriesList := FundtoSeriesMapper[fundCode]
	for _, series := range seriesList {

		fundId := fundCode

		if series != "" {
			fundId = fundId + "_" + series
		}

		series = strings.Replace(series, "_", " - ", 1)

		fund := NewPrivateFund(fundId, series)
		funds = append(funds, fund)
	}

	return funds
}
