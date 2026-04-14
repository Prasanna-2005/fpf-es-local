package curio

import (
	"database/sql"
	"fmt"
	"quantrium-tech/form_pf_extractor_service/executors"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"
)

func CreateAllFundSeries(fundCode string) []*Fund {

	var funds []*Fund

	fundId := fundCode

	fund := NewFund(fundId)
	funds = append(funds, fund)

	return funds
}

func StaticFunds(FundShortCode string) (*executors.Form1bData, *Fund) {
	FundShortCodeUpper := strings.ToUpper(FundShortCode)
	fundData := executors.NewForm1bData()
	fundData.OtherPercent = 100
	fundObj := NewFund(FundShortCodeUpper)

	fundObj.fundID = FundShortCodeUpper
	fundObj.series = FundShortCodeUpper

	return fundData, fundObj

}

// normalizeCell cleans a raw cell string:
// - strips leading/trailing whitespace
// - removes all internal whitespace (spaces, tabs, non-breaking spaces, escape sequences)
// - lowercases everything
func normalizeCell(s string) string {
	// Remove escape sequences like \t, \n, \r, \x00, etc.
	s = strings.Map(func(r rune) rune {
		if unicode.IsControl(r) {
			return -1 // drop control characters
		}
		return r
	}, s)

	// Remove ALL whitespace (including mid-word spaces → "Unfunded Commitment" == "UnfundedCommitment")
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, "")

	// Lowercase
	return strings.ToLower(s)
}

func singularize(s string) string {
	if strings.HasSuffix(s, "ies") {
		return s[:len(s)-3] + "y" // committies → committy
	}
	if strings.HasSuffix(s, "es") {
		return s[:len(s)-2] // branches→ branch
	}
	if strings.HasSuffix(s, "s") {
		return s[:len(s)-1] // commitments → commitment
	}
	return s
}

// assignFundPerf copies every float64/string from src (FundPerformance) into dst
// (Form1bData) converting to sql.NullFloat64/sql.NullString when {term,colName} matches.
func assignFundPerf(dst *executors.Form1bData, irr map[string]*IRR) {
	t := reflect.TypeOf(executors.Form1bData{})
	vDst := reflect.ValueOf(dst).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		term := field.Tag.Get("term")
		irrType := field.Tag.Get("colName")

		if term == "" || colName == "" {
			continue
		}

		if fndPerf, ok := irr[term]; !ok {
			switch irrType {
			case "gross":
				vDst.Field(i).Set(reflect.ValueOf(sql.NullFloat64{
					Float64: fndPerf.grossReturns,
					Valid:   true,
				}))
			case "net":
				vDst.Field(i).Set(reflect.ValueOf(sql.NullFloat64{
					Float64: fndPerf.netReturns,
					Valid:   true,
				}))
			case "date":
				if fndPerf.grossReturns != 0 || fndPerf.netReturns != 0 {
					vDst.Field(i).Set(reflect.ValueOf(sql.NullString{
						String: lastDayOfTerm(term, time.Now().Year()-1),
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
