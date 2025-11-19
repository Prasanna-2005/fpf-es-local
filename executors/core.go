package executors

import (
	"fmt"
	"time"
)

// Form1bData holds the structure of the section 1b in Form PF.
// It will be dynamically increased based on the needs to answer other questions
type Form1bData struct {
	FundName string `fieldNo:"5" formLineNo:"5_a"`
	FundID   string `fieldNo:"5" formLineNo:"5_b"`
	FundNFA  string `fieldNo:"5" formLineNo:"5_c"`
	FundLEI  string `fieldNo:"5" formLineNo:"5_d"`

	GrossAsset         float64 `fieldNo:"8" formLineNo:"8"`
	NetAsset           float64 `fieldNo:"9" formLineNo:"9"`
	EquityInOtherFunds float64 `fieldNo:"10" formLineNo:"10"`
	Level3Asset        float64 `fieldNo:"14" formLineNo:"14_assets_level3"`
	CostBasedAsset     float64 `fieldNo:"14" formLineNo:"14_assets_costbased"`
	CostBasedLiability float64 `fieldNo:"14" formLineNo:"14_liabilities_costbased"`
	Top5Percent        float64 `fieldNo:"15" formLineNo:"15"`
	UsPercent          float64 `fieldNo:"16" formLineNo:"16_a"`
	NonProfitPercent   float64 `fieldNo:"16" formLineNo:"16_g"`
	PensionPlanPercent float64 `fieldNo:"16" formLineNo:"16_h"`
	OtherPercent       float64 `fieldNo:"16" formLineNo:"16_n"`
	SovereignPercent   float64 `fieldNo:"16" formLineNo:"16_l"`
	PrivateFundPercent float64 `fieldNo:"16" formLineNo:"16_f"`

	Q1EndDate    string  `fieldNo:"17" formLineNo:"17_(d)_last_day_of_fiscal_period"`
	Q1GrossValue float64 `fieldNo:"17" formLineNo:"17_(d)_gross_performance"`
	Q1NetValue   float64 `fieldNo:"17" formLineNo:"17_(d)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	Q2EndDate    string  `fieldNo:"17" formLineNo:"17_(h)_last_day_of_fiscal_period"`
	Q2GrossValue float64 `fieldNo:"17" formLineNo:"17_(h)_gross_performance"`
	Q2NetValue   float64 `fieldNo:"17" formLineNo:"17_(h)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	Q3EndDate    string  `fieldNo:"17" formLineNo:"17_(l)_last_day_of_fiscal_period"`
	Q3GrossValue float64 `fieldNo:"17" formLineNo:"17_(l)_gross_performance"`
	Q3NetValue   float64 `fieldNo:"17" formLineNo:"17_(l)_Net_of_management_fees_and_incentives_fees_and_allocation"`
}

func NewForm1bData() *Form1bData {

	lastYear := time.Now().Year() - 1
	return &Form1bData{
		Q1EndDate: fmt.Sprintf("03/31/%d", lastYear),
		Q2EndDate: fmt.Sprintf("06/30/%d", lastYear),
		Q3EndDate: fmt.Sprintf("09/30/%d", lastYear),
	}

}
