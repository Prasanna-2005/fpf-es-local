package privateinvestors

import (
	"database/sql"
	"quantrium-tech/form_pf_extractor_service/executors"
	"reflect"
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

	for i := 0; i < tDst.NumField(); i++ {
		f := tDst.Field(i)
		term := f.Tag.Get("term")
		col := f.Tag.Get("colName")
		if term == "" || col == "" {
			continue
		}

		src, ok := srcMap[keyT{term, col}]
		if !ok {
			continue
		}

		dstFld := vDst.Field(i)
		switch dstFld.Type() {
		case reflect.TypeOf(sql.NullFloat64{}):
			if src.Kind() == reflect.Float64 {
				dstFld.Set(reflect.ValueOf(sql.NullFloat64{
					Float64: src.Float(),
					Valid:   src.Float() != 0,
				}))
			}
		case reflect.TypeOf(sql.NullString{}):
			if src.Kind() == reflect.String {
				s := src.String()
				dstFld.Set(reflect.ValueOf(sql.NullString{
					String: s,
					Valid:  s != "",
				}))
			}
		}
	}
}
