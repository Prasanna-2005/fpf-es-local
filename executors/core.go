package executors

import (
	"fmt"
	"time"
)

// Form1bData holds the structure of the section 1b in Form PF.
// It will be dynamically increased based on the needs to answer other questions
type Form1bData struct {
	FundName                string `fieldNo:"5" formLineNo:"5_a"`
	FundID                  string `fieldNo:"5" formLineNo:"5_b"`
	FundNFA                 string `fieldNo:"5" formLineNo:"5_c"`
	FundLEI                 string `fieldNo:"5" formLineNo:"5_d"`
	MasterFund              bool   `fieldNo:"6" formLineNo:"6"`
	ParallelFundsAggregated bool   `fieldNo:"7" formLineNo:"7_a"`

	GrossAsset                  float64 `fieldNo:"8" formLineNo:"8"`
	NetAsset                    float64 `fieldNo:"9" formLineNo:"9"`
	EquityInOtherFunds          float64 `fieldNo:"10" formLineNo:"10"`
	ParallelManagedAccountValue float64 `fieldNo:"11" formLineNo:"11"`
	BorrowedTotalAmount         float64 `fieldNo:"12" formLineNo:"12_a"`
	BorrowedUSFinInstiAmt       float64 `fieldNo:"12" formLineNo:"12_b"`
	BorrowedNonUSFinInstiAmt    float64 `fieldNo:"12" formLineNo:"12_c"`
	BorrowedUSNonFinInstiAmt    float64 `fieldNo:"12" formLineNo:"12_d"`
	BorrowedNonUSNonFinInstiAmt float64 `fieldNo:"12" formLineNo:"12_e"`
	HasDerivativesPositions     bool    `fieldNo:"13" formLineNo:"13_a"`
	// DerivativesPositionsValue   float64 `fieldNo:"13" formLineNo:"13_b"`

	Level1Asset        float64 `fieldNo:"14" formLineNo:"14_assets_level1"`
	Level2Asset        float64 `fieldNo:"14" formLineNo:"14_assets_level2"`
	Level3Asset        float64 `fieldNo:"14" formLineNo:"14_assets_level3"`
	Level1Liabilities  float64 `fieldNo:"14" formLineNo:"14_liabilities_level1"`
	Level2Liabilities  float64 `fieldNo:"14" formLineNo:"14_liabilities_level2"`
	Level3Liabilities  float64 `fieldNo:"14" formLineNo:"14_liabilities_level3"`
	CostBasedAsset     float64 `fieldNo:"14" formLineNo:"14_assets_costbased"`
	Level1Liability    float64 `fieldNo:"14" formLineNo:"14_liabilities_level1"`
	Level2Liability    float64 `fieldNo:"14" formLineNo:"14_liabilities_level2"`
	Level3Liability    float64 `fieldNo:"14" formLineNo:"14_liabilities_level3"`
	CostBasedLiability float64 `fieldNo:"14" formLineNo:"14_liabilities_costbased"`

	Top5Percent                float64 `fieldNo:"15" formLineNo:"15"`
	UsPercent                  float64 `fieldNo:"16" formLineNo:"16_a"`
	NonUSPersonsPercent        float64 `fieldNo:"16" formLineNo:"16_b"`
	BrokerDealersPercent       float64 `fieldNo:"16" formLineNo:"16_c"`
	InsuranceCompaniesPercent  float64 `fieldNo:"16" formLineNo:"16_d"`
	InvestmentCompaniesPercent float64 `fieldNo:"16" formLineNo:"16_e"`
	PrivateFundsPercent        float64 `fieldNo:"16" formLineNo:"16_f"`
	NonProfitPercent           float64 `fieldNo:"16" formLineNo:"16_g"`
	PensionPlanPercent         float64 `fieldNo:"16" formLineNo:"16_h"`
	BankingInstitutionsPercent float64 `fieldNo:"16" formLineNo:"16_i"`
	GovernmentEntitiesPercent  float64 `fieldNo:"16" formLineNo:"16_j"`
	GovernmentPensionsPercent  float64 `fieldNo:"16" formLineNo:"16_k"`
	ForeignInstitutionsPercent float64 `fieldNo:"16" formLineNo:"16_l"`
	NonUSIntermediariesPercent float64 `fieldNo:"16" formLineNo:"16_m"`
	OtherPercent               float64 `fieldNo:"16" formLineNo:"16_n"`
	SovereignPercent           float64 `fieldNo:"16" formLineNo:"16_l"`
	PrivateFundPercent         float64 `fieldNo:"16" formLineNo:"16_f"`
	/*  MONTHLY  */
	// (a) 1st month of reporting fund’s fiscal year
	JanEndDate    string  `fieldNo:"17" formLineNo:"17_(a)_last_day_of_fiscal_period"`
	JanGrossValue float64 `fieldNo:"17" formLineNo:"17_(a)_gross_performance"`
	JanNetValue   float64 `fieldNo:"17" formLineNo:"17_(a)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// (b) 2nd month …
	FebEndDate    string  `fieldNo:"17" formLineNo:"17_(b)_last_day_of_fiscal_period"`
	FebGrossValue float64 `fieldNo:"17" formLineNo:"17_(b)_gross_performance"`
	FebNetValue   float64 `fieldNo:"17" formLineNo:"17_(b)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// (c) 3rd month …
	MarEndDate    string  `fieldNo:"17" formLineNo:"17_(c)_last_day_of_fiscal_period"`
	MarGrossValue float64 `fieldNo:"17" formLineNo:"17_(c)_gross_performance"`
	MarNetValue   float64 `fieldNo:"17" formLineNo:"17_(c)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// (e) 4th month …
	AprEndDate    string  `fieldNo:"17" formLineNo:"17_(e)_last_day_of_fiscal_period"`
	AprGrossValue float64 `fieldNo:"17" formLineNo:"17_(e)_gross_performance"`
	AprNetValue   float64 `fieldNo:"17" formLineNo:"17_(e)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// (f) 5th month …
	MayEndDate    string  `fieldNo:"17" formLineNo:"17_(f)_last_day_of_fiscal_period"`
	MayGrossValue float64 `fieldNo:"17" formLineNo:"17_(f)_gross_performance"`
	MayNetValue   float64 `fieldNo:"17" formLineNo:"17_(f)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// (g) 6th month …
	JunEndDate    string  `fieldNo:"17" formLineNo:"17_(g)_last_day_of_fiscal_period"`
	JunGrossValue float64 `fieldNo:"17" formLineNo:"17_(g)_gross_performance"`
	JunNetValue   float64 `fieldNo:"17" formLineNo:"17_(g)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// (i) 7th month …
	JulEndDate    string  `fieldNo:"17" formLineNo:"17_(i)_last_day_of_fiscal_period"`
	JulGrossValue float64 `fieldNo:"17" formLineNo:"17_(i)_gross_performance"`
	JulNetValue   float64 `fieldNo:"17" formLineNo:"17_(i)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// (j) 8th month …
	AugEndDate    string  `fieldNo:"17" formLineNo:"17_(j)_last_day_of_fiscal_period"`
	AugGrossValue float64 `fieldNo:"17" formLineNo:"17_(j)_gross_performance"`
	AugNetValue   float64 `fieldNo:"17" formLineNo:"17_(j)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// (k) 9th month …
	SepEndDate    string  `fieldNo:"17" formLineNo:"17_(k)_last_day_of_fiscal_period"`
	SepGrossValue float64 `fieldNo:"17" formLineNo:"17_(k)_gross_performance"`
	SepNetValue   float64 `fieldNo:"17" formLineNo:"17_(k)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// (m) 10th month …
	OctEndDate    string  `fieldNo:"17" formLineNo:"17_(m)_last_day_of_fiscal_period"`
	OctGrossValue float64 `fieldNo:"17" formLineNo:"17_(m)_gross_performance"`
	OctNetValue   float64 `fieldNo:"17" formLineNo:"17_(m)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// (n) 11th month …
	NovEndDate    string  `fieldNo:"17" formLineNo:"17_(n)_last_day_of_fiscal_period"`
	NovGrossValue float64 `fieldNo:"17" formLineNo:"17_(n)_gross_performance"`
	NovNetValue   float64 `fieldNo:"17" formLineNo:"17_(n)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// (o) 12th month …
	DecEndDate    string  `fieldNo:"17" formLineNo:"17_(o)_last_day_of_fiscal_period"`
	DecGrossValue float64 `fieldNo:"17" formLineNo:"17_(o)_gross_performance"`
	DecNetValue   float64 `fieldNo:"17" formLineNo:"17_(o)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	/*  QUARTERLY  */
	// (d) First quarter …
	Q1EndDate    string  `fieldNo:"17" formLineNo:"17_(d)_last_day_of_fiscal_period"`
	Q1GrossValue float64 `fieldNo:"17" formLineNo:"17_(d)_gross_performance"`
	Q1NetValue   float64 `fieldNo:"17" formLineNo:"17_(d)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// (h) Second quarter …
	Q2EndDate    string  `fieldNo:"17" formLineNo:"17_(h)_last_day_of_fiscal_period"`
	Q2GrossValue float64 `fieldNo:"17" formLineNo:"17_(h)_gross_performance"`
	Q2NetValue   float64 `fieldNo:"17" formLineNo:"17_(h)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// (l) Third quarter …
	Q3EndDate    string  `fieldNo:"17" formLineNo:"17_(l)_last_day_of_fiscal_period"`
	Q3GrossValue float64 `fieldNo:"17" formLineNo:"17_(l)_gross_performance"`
	Q3NetValue   float64 `fieldNo:"17" formLineNo:"17_(l)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// (p) Fourth quarter …
	Q4EndDate    string  `fieldNo:"17" formLineNo:"17_(p)_last_day_of_fiscal_period"`
	Q4GrossValue float64 `fieldNo:"17" formLineNo:"17_(p)_gross_performance"`
	Q4NetValue   float64 `fieldNo:"17" formLineNo:"17_(p)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	/*  ANNUAL  */
	// (q) Reporting fund’s most recently completed fiscal year …
	YearEndDate    string  `fieldNo:"17" formLineNo:"17_(q)_last_day_of_fiscal_period"`
	YearGrossValue float64 `fieldNo:"17" formLineNo:"17_(q)_gross_performance"`
	YearNetValue   float64 `fieldNo:"17" formLineNo:"17_(q)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// Q1EndDate    string  `fieldNo:"17" formLineNo:"17_(d)_last_day_of_fiscal_period"`
	// Q1GrossValue float64 `fieldNo:"17" formLineNo:"17_(d)_gross_performance"`
	// Q1NetValue   float64 `fieldNo:"17" formLineNo:"17_(d)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// Q2EndDate    string  `fieldNo:"17" formLineNo:"17_(h)_last_day_of_fiscal_period"`
	// Q2GrossValue float64 `fieldNo:"17" formLineNo:"17_(h)_gross_performance"`
	// Q2NetValue   float64 `fieldNo:"17" formLineNo:"17_(h)_Net_of_management_fees_and_incentives_fees_and_allocation"`

	// Q3EndDate    string  `fieldNo:"17" formLineNo:"17_(l)_last_day_of_fiscal_period"`
	// Q3GrossValue float64 `fieldNo:"17" formLineNo:"17_(l)_gross_performance"`
	// Q3NetValue   float64 `fieldNo:"17" formLineNo:"17_(l)_Net_of_management_fees_and_incentives_fees_and_allocation"`
}

func NewForm1bData() *Form1bData {
	lastYear := time.Now().Year() - 1
	return &Form1bData{

		MasterFund:              false, // 6
		ParallelFundsAggregated: false, // 7
		HasDerivativesPositions: false, //13_b

		// prev existing code
		Q1EndDate: fmt.Sprintf("%d/03/31", lastYear),
		Q2EndDate: fmt.Sprintf("%d/06/30", lastYear),
		Q3EndDate: fmt.Sprintf("%d/09/30", lastYear),
	}
}
