package executors

import "database/sql"

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

	// Months
	JanEndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(a)_last_day_of_fiscal_period" colName:"date" term:"january"`
	JanGrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(a)_gross_performance"          colName:"gross" term:"january"`
	JanNetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(a)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"january"`

	FebEndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(b)_last_day_of_fiscal_period" colName:"date" term:"february"`
	FebGrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(b)_gross_performance"          colName:"gross" term:"february"`
	FebNetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(b)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"february"`

	MarEndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(c)_last_day_of_fiscal_period" colName:"date" term:"march"`
	MarGrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(c)_gross_performance"          colName:"gross" term:"march"`
	MarNetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(c)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"march"`

	// Quarters (PDF lines moved only)
	Q1EndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(d)_last_day_of_fiscal_period" colName:"date" term:"first"`
	Q1GrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(d)_gross_performance"          colName:"gross" term:"first"`
	Q1NetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(d)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"first"`

	AprEndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(e)_last_day_of_fiscal_period" colName:"date" term:"april"`
	AprGrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(e)_gross_performance"          colName:"gross" term:"april"`
	AprNetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(e)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"april"`

	MayEndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(f)_last_day_of_fiscal_period" colName:"date" term:"may"`
	MayGrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(f)_gross_performance"          colName:"gross" term:"may"`
	MayNetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(f)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"may"`

	JunEndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(g)_last_day_of_fiscal_period" colName:"date" term:"june"`
	JunGrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(g)_gross_performance"          colName:"gross" term:"june"`
	JunNetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(g)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"june"`

	Q2EndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(h)_last_day_of_fiscal_period" colName:"date" term:"second"`
	Q2GrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(h)_gross_performance"          colName:"gross" term:"second"`
	Q2NetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(h)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"second"`

	JulEndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(i)_last_day_of_fiscal_period" colName:"date" term:"july"`
	JulGrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(i)_gross_performance"          colName:"gross" term:"july"`
	JulNetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(i)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"july"`

	AugEndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(j)_last_day_of_fiscal_period" colName:"date" term:"august"`
	AugGrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(j)_gross_performance"          colName:"gross" term:"august"`
	AugNetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(j)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"august"`

	SepEndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(k)_last_day_of_fiscal_period" colName:"date" term:"september"`
	SepGrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(k)_gross_performance"          colName:"gross" term:"september"`
	SepNetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(k)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"september"`

	Q3EndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(l)_last_day_of_fiscal_period" colName:"date" term:"third"`
	Q3GrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(l)_gross_performance"          colName:"gross" term:"third"`
	Q3NetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(l)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"third"`

	OctEndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(m)_last_day_of_fiscal_period" colName:"date" term:"october"`
	OctGrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(m)_gross_performance"          colName:"gross" term:"october"`
	OctNetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(m)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"october"`

	NovEndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(n)_last_day_of_fiscal_period" colName:"date" term:"november"`
	NovGrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(n)_gross_performance"          colName:"gross" term:"november"`
	NovNetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(n)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"november"`

	DecEndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(o)_last_day_of_fiscal_period" colName:"date" term:"december"`
	DecGrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(o)_gross_performance"          colName:"gross" term:"december"`
	DecNetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(o)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"december"`

	Q4EndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(p)_last_day_of_fiscal_period" colName:"date" term:"fourth"`
	Q4GrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(p)_gross_performance"          colName:"gross" term:"fourth"`
	Q4NetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(p)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"fourth"`

	YearEndDate    sql.NullString  `fieldNo:"17" formLineNo:"17_(q)_last_day_of_fiscal_period" colName:"date" term:"year"`
	YearGrossValue sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(q)_gross_performance"          colName:"gross" term:"year"`
	YearNetValue   sql.NullFloat64 `fieldNo:"17" formLineNo:"17_(q)_Net_of_management_fees_and_incentives_fees_and_allocation" colName:"net" term:"year"`
}

func NewForm1bData() *Form1bData {
	// lastYear := time.Now().Year() - 1
	return &Form1bData{

		MasterFund:              false, // 6
		ParallelFundsAggregated: false, // 7
		HasDerivativesPositions: false, //13_b

	}
}
