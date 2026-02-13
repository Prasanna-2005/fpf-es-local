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
type Form1cData struct {
	FundName string `fieldNo:"18" formLineNo:"18_a"`
	FundID   string `fieldNo:"18" formLineNo:"18_b"`

	Strategy string `fieldNo:"19" formLineNo:"19"`
	// Equity
	EquityMarketNeutralStrategy bool `fieldNo:"20" formLineNo:"20_(a)_Strategy"`
	// EquityMarketNeutralNAV      float64 `fieldNo:"20" formLineNo:"20_(a)_%_of_NAV"`
	// EquityMarketNeutralCapital  float64 `fieldNo:"20" formLineNo:"20_(a)_%_of_capital"`

	EquityLongShortStrategy bool `fieldNo:"20" formLineNo:"20_(b)_Strategy"`
	// EquityLongShortNAV      float64 `fieldNo:"20" formLineNo:"20_(b)_%_of_NAV"`
	// EquityLongShortCapital  float64 `fieldNo:"20" formLineNo:"20_(b)_%_of_capital"`

	EquityShortBiasStrategy bool `fieldNo:"20" formLineNo:"20_(c)_Strategy"`
	// EquityShortBiasNAV      float64 `fieldNo:"20" formLineNo:"20_(c)_%_of_NAV"`
	// EquityShortBiasCapital  float64 `fieldNo:"20" formLineNo:"20_(c)_%_of_capital"`

	EquityLongBiasStrategy bool   `fieldNo:"20" formLineNo:"20_(d)_Strategy"`
	EquityLongBiasNAV      string `fieldNo:"20" formLineNo:"20_(d)_%_of_NAV"`
	// EquityLongBiasCapital  float64 `fieldNo:"20" formLineNo:"20_(d)_%_of_capital"`

	// Macro
	MacroActiveTradingStrategy bool `fieldNo:"20" formLineNo:"20_(e)_Strategy"`
	// MacroActiveTradingNAV      float64 `fieldNo:"20" formLineNo:"20_(e)_%_of_NAV"`
	// MacroActiveTradingCapital  float64 `fieldNo:"20" formLineNo:"20_(e)_%_of_capital"`

	MacroCommodityStrategy bool `fieldNo:"20" formLineNo:"20_(f)_Strategy"`
	// MacroCommodityNAV      float64 `fieldNo:"20" formLineNo:"20_(f)_%_of_NAV"`
	// MacroCommodityCapital  float64 `fieldNo:"20" formLineNo:"20_(f)_%_of_capital"`

	MacroCurrencyStrategy bool `fieldNo:"20" formLineNo:"20_(g)_Strategy"`
	// MacroCurrencyNAV      float64 `fieldNo:"20" formLineNo:"20_(g)_%_of_NAV"`
	// MacroCurrencyCapital  float64 `fieldNo:"20" formLineNo:"20_(g)_%_of_capital"`

	MacroGlobalMacroStrategy bool `fieldNo:"20" formLineNo:"20_(h)_Strategy"`
	// MacroGlobalMacroNAV      float64 `fieldNo:"20" formLineNo:"20_(h)_%_of_NAV"`
	// MacroGlobalMacroCapital  float64 `fieldNo:"20" formLineNo:"20_(h)_%_of_capital"`

	// Relative Value
	RVFixedIncomeAssetBackedStrategy bool `fieldNo:"20" formLineNo:"20_(i)_Strategy"`
	// RVFixedIncomeAssetBackedNAV      float64 `fieldNo:"20" formLineNo:"20_(i)_%_of_NAV"`
	// RVFixedIncomeAssetBackedCapital  float64 `fieldNo:"20" formLineNo:"20_(i)_%_of_capital"`

	RVFixedIncomeConvArbStrategy bool `fieldNo:"20" formLineNo:"20_(j)_Strategy"`
	// RVFixedIncomeConvArbNAV      float64 `fieldNo:"20" formLineNo:"20_(j)_%_of_NAV"`
	// RVFixedIncomeConvArbCapital  float64 `fieldNo:"20" formLineNo:"20_(j)_%_of_capital"`

	RVFixedIncomeCorporateStrategy bool `fieldNo:"20" formLineNo:"20_(k)_Strategy"`
	// RVFixedIncomeCorporateNAV      float64 `fieldNo:"20" formLineNo:"20_(k)_%_of_NAV"`
	// RVFixedIncomeCorporateCapital  float64 `fieldNo:"20" formLineNo:"20_(k)_%_of_capital"`

	RVFixedIncomeSovereignStrategy bool `fieldNo:"20" formLineNo:"20_(l)_Strategy"`
	// RVFixedIncomeSovereignNAV      float64 `fieldNo:"20" formLineNo:"20_(l)_%_of_NAV"`
	// RVFixedIncomeSovereignCapital  float64 `fieldNo:"20" formLineNo:"20_(l)_%_of_capital"`

	RVVolatilityArbStrategy bool `fieldNo:"20" formLineNo:"20_(m)_Strategy"`
	// RVVolatilityArbNAV      float64 `fieldNo:"20" formLineNo:"20_(m)_%_of_NAV"`
	// RVVolatilityArbCapital  float64 `fieldNo:"20" formLineNo:"20_(m)_%_of_capital"`

	// Event Driven
	EventDrivenDistressedStrategy bool `fieldNo:"20" formLineNo:"20_(n)_Strategy"`
	// EventDrivenDistressedNAV      float64 `fieldNo:"20" formLineNo:"20_(n)_%_of_NAV"`
	// EventDrivenDistressedCapital  float64 `fieldNo:"20" formLineNo:"20_(n)_%_of_capital"`

	EventDrivenRiskArbStrategy bool `fieldNo:"20" formLineNo:"20_(o)_Strategy"`
	// EventDrivenRiskArbNAV      float64 `fieldNo:"20" formLineNo:"20_(o)_%_of_NAV"`
	// EventDrivenRiskArbCapital  float64 `fieldNo:"20" formLineNo:"20_(o)_%_of_capital"`

	EventDrivenEqSpecialSitStrategy bool `fieldNo:"20" formLineNo:"20_(p)_Strategy"`
	// EventDrivenEqSpecialSitNAV      float64 `fieldNo:"20" formLineNo:"20_(p)_%_of_NAV"`
	// EventDrivenEqSpecialSitCapital  float64 `fieldNo:"20" formLineNo:"20_(p)_%_of_capital"`

	// Credit
	CreditLongShortStrategy bool `fieldNo:"20" formLineNo:"20_(q)_Strategy"`
	// CreditLongShortNAV      float64 `fieldNo:"20" formLineNo:"20_(q)_%_of_NAV"`
	// CreditLongShortCapital  float64 `fieldNo:"20" formLineNo:"20_(q)_%_of_capital"`

	CreditAssetBasedLendingStrategy bool `fieldNo:"20" formLineNo:"20_(r)_Strategy"`
	// CreditAssetBasedLendingNAV      float64 `fieldNo:"20" formLineNo:"20_(r)_%_of_NAV"`
	// CreditAssetBasedLendingCapital  float64 `fieldNo:"20" formLineNo:"20_(r)_%_of_capital"`

	// Managed Futures / CTA
	ManagedFuturesFundamentalStrategy bool `fieldNo:"20" formLineNo:"20_(s)_Strategy"`
	// ManagedFuturesFundamentalNAV      float64 `fieldNo:"20" formLineNo:"20_(s)_%_of_NAV"`
	// ManagedFuturesFundamentalCapital  float64 `fieldNo:"20" formLineNo:"20_(s)_%_of_capital"`

	ManagedFuturesQuantitativeStrategy bool `fieldNo:"20" formLineNo:"20_(t)_Strategy"`
	// ManagedFuturesQuantitativeNAV      float64 `fieldNo:"20" formLineNo:"20_(t)_%_of_NAV"`
	// ManagedFuturesQuantitativeCapital  float64 `fieldNo:"20" formLineNo:"20_(t)_%_of_capital"`

	// Other
	InvestmentInOtherFundsStrategy bool `fieldNo:"20" formLineNo:"20_(u)_Strategy"`
	// InvestmentInOtherFundsNAV      float64 `fieldNo:"20" formLineNo:"20_(u)_%_of_NAV"`
	// InvestmentInOtherFundsCapital  float64 `fieldNo:"20" formLineNo:"20_(u)_%_of_capital"`

	OtherStrategy bool `fieldNo:"20" formLineNo:"20_(v)_Strategy"`
	// OtherNAV      float64 `fieldNo:"20" formLineNo:"20_(v)_%_of_NAV"`
	// OtherCapital  float64 `fieldNo:"20" formLineNo:"20_(v)_%_of_capital"`

	HFTPercentage string `fieldNo:"21" formLineNo:"21"`
	// Row A
	PercentageCounterpartyALegalName       string `fieldNo:"22" formLineNo:"22_(a)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	PercentageCounterpartyAAffiliatedInsti string `fieldNo:"22" formLineNo:"22_(a)_Indicate below if the counterparty is affiliated with a major financial institution"`
	PercentageCounterpartyAExposure        string `fieldNo:"22" formLineNo:"22_(a)_Exposure (% of reporting fund's net asset value)"`
	PercentageAInsticode                   string `fieldNo:"22" formLineNo:"22_(a)_Institution code"`

	// Row B
	PercentageCounterpartyBLegalName       string `fieldNo:"22" formLineNo:"22_(b)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	PercentageCounterpartyBAffiliatedInsti string `fieldNo:"22" formLineNo:"22_(b)_Indicate below if the counterparty is affiliated with a major financial institution"`
	PercentageCounterpartyBExposure        string `fieldNo:"22" formLineNo:"22_(b)_Exposure (% of reporting fund's net asset value)"`
	PercentageBInsticode                   string `fieldNo:"22" formLineNo:"22_(b)_Institution code"`
	// Row C
	PercentageCounterpartyCLegalName       string `fieldNo:"22" formLineNo:"22_(c)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	PercentageCounterpartyCAffiliatedInsti string `fieldNo:"22" formLineNo:"22_(c)_Indicate below if the counterparty is affiliated with a major financial institution"`
	PercentageCounterpartyCExposure        string `fieldNo:"22" formLineNo:"22_(c)_Exposure (% of reporting fund's net asset value)"`
	PercentageCInsticode                   string `fieldNo:"22" formLineNo:"22_(c)_Institution code"`
	// Row D
	PercentageCounterpartyDLegalName       string `fieldNo:"22" formLineNo:"22_(d)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	PercentageCounterpartyDAffiliatedInsti string `fieldNo:"22" formLineNo:"22_(d)_Indicate below if the counterparty is affiliated with a major financial institution"`
	PercentageCounterpartyDExposure        string `fieldNo:"22" formLineNo:"22_(d)_Exposure (% of reporting fund's net asset value)"`
	PercentageDInsticode                   string `fieldNo:"22" formLineNo:"22_(d)_Institution code"`
	// Row E
	PercentageCounterpartyELegalName       string `fieldNo:"22" formLineNo:"22_(e)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	PercentageCounterpartyEAffiliatedInsti string `fieldNo:"22" formLineNo:"22_(e)_Indicate below if the counterparty is affiliated with a major financial institution"`
	PercentageCounterpartyEExposure        string `fieldNo:"22" formLineNo:"22_(e)_Exposure (% of reporting fund's net asset value)"`
	PercentageEInsticode                   string `fieldNo:"22" formLineNo:"22_(e)_Institution code"`
	// Row A
	USDCounterpartyALegalName       string `fieldNo:"23" formLineNo:"23_(a)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	USDCounterpartyAAffiliatedInsti string `fieldNo:"23" formLineNo:"23_(a)_Indicate below if the counterparty is affiliated with a major financial institution"`
	USDCounterpartyAExposure        string `fieldNo:"23" formLineNo:"23_(a)_Exposure (in U.S. dollars))"`
	USDAInsticode                   string `fieldNo:"23" formLineNo:"23_(a)_Institution code"`
	// Row B
	USDCounterpartyBLegalName       string `fieldNo:"23" formLineNo:"23_(b)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	USDCounterpartyBAffiliatedInsti string `fieldNo:"23" formLineNo:"23_(b)_Indicate below if the counterparty is affiliated with a major financial institution"`
	USDCounterpartyBExposure        string `fieldNo:"23" formLineNo:"23_(b)_Exposure (in U.S. dollars))"`
	USDBInsticode                   string `fieldNo:"23" formLineNo:"23_(b)_Institution code"`
	// Row C
	USDCounterpartyCLegalName       string `fieldNo:"23" formLineNo:"23_(c)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	USDCounterpartyCAffiliatedInsti string `fieldNo:"23" formLineNo:"23_(c)_Indicate below if the counterparty is affiliated with a major financial institution"`
	USDCounterpartyCExposure        string `fieldNo:"23" formLineNo:"23_(c)_Exposure (in U.S. dollars))"`
	USDCInsticode                   string `fieldNo:"23" formLineNo:"23_(c)_Institution code"`
	// Row D
	USDCounterpartyDLegalName       string `fieldNo:"23" formLineNo:"23_(d)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	USDCounterpartyDAffiliatedInsti string `fieldNo:"23" formLineNo:"23_(d)_Indicate below if the counterparty is affiliated with a major financial institution"`
	USDCounterpartyDExposure        string `fieldNo:"23" formLineNo:"23_(d)_Exposure (in U.S. dollars))"`
	USDDInsticode                   string `fieldNo:"23" formLineNo:"23_(d)_Institution code"`
	// Row E
	USDCounterpartyELegalName       string `fieldNo:"23" formLineNo:"23_(e)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	USDCounterpartyEAffiliatedInsti string `fieldNo:"23" formLineNo:"23_(e)_Indicate below if the counterparty is affiliated with a major financial institution"`
	USDCounterpartyEExposure        string `fieldNo:"23" formLineNo:"23_(e)_Exposure (in U.S. dollars))"`
	USDEInsticode                   string `fieldNo:"23" formLineNo:"23_(e)_Institution code"`

	SecRegExchange         string `fieldNo:"24" formLineNo:"24_a_i"`
	SecOTC                 string `fieldNo:"24" formLineNo:"24_a_ii"`
	DerivativesRegExchange string `fieldNo:"24" formLineNo:"24_b_i"`
	DerivativesOTC         string `fieldNo:"24" formLineNo:"24_b_ii"`
	DerivativesCCP         string `fieldNo:"24" formLineNo:"24_c_i"`
	DerivativesBilateraly  string `fieldNo:"24" formLineNo:"24_c_ii"`
	RepoCCP                string `fieldNo:"24" formLineNo:"24_d_i"`
	RepoBilateraly         string `fieldNo:"24" formLineNo:"24_d_ii"`
	RepoTriParty           string `fieldNo:"24" formLineNo:"24_d_iii"`

	OtherTradePercent float64 `fieldNo:"25" formLineNo:"25"`
}

type Form1aData struct {
	LegalName           string `fieldNo:"1" formLineNo:"1_a_i"`
	CrdNumber           string `fieldNo:"1" formLineNo:"1_a_ii"`
	Sec801Number        string `fieldNo:"1" formLineNo:"1_a_iii"`
	NfaId               string `fieldNo:"1" formLineNo:"1_a_iv"`
	LargeTraderId       string `fieldNo:"1" formLineNo:"1_a_v"`
	LargeTraderIdSuffix string `fieldNo:"1" formLineNo:"1_a_vi"`

	FirmIndividualName      string `fieldNo:"2" formLineNo:"2_a_i"`
	FirmIndividualSignature string `fieldNo:"2" formLineNo:"2_a_ii"`
	FirmIndividualTitle     string `fieldNo:"2" formLineNo:"2_a_iii"`
	FirmIndividualEmail     string `fieldNo:"2" formLineNo:"2_a_iv"`
	FirmIndividualPhno      string `fieldNo:"2" formLineNo:"2_a_v"`

	FirmDependentsRelName string `fieldNo:"2" formLineNo:"2_b_i"`
	FirmDependentsName    string `fieldNo:"2" formLineNo:"2_b_ii"`
	FirmDependentsSign    string `fieldNo:"2" formLineNo:"2_b_iii"`
	FirmDependentsTitle   string `fieldNo:"2" formLineNo:"2_b_iv"`
	FirmDependentsEmail   string `fieldNo:"2" formLineNo:"2_b_v"`
	FirmDependentsPhno    string `fieldNo:"2" formLineNo:"2_b_vi"`

	HedgeFundsGAV     float64 `fieldNo:"3" formLineNo:"3_(a)_Regulatory_assets_under_management"`
	HedgeFundsNAV     float64 `fieldNo:"3" formLineNo:"3_(a)_Net_assets_under_management"`
	LiquidityGAV      float64 `fieldNo:"3" formLineNo:"3_(b)_Regulatory_assets_under_management"`
	LiquidityNAV      float64 `fieldNo:"3" formLineNo:"3_(b)_Net_assets_under_management"`
	PrivateGAV        float64 `fieldNo:"3" formLineNo:"3_(c)_Regulatory_assets_under_management"`
	PrivateNAV        float64 `fieldNo:"3" formLineNo:"3_(c)_Net_assets_under_management"`
	RealEstateGAV     float64 `fieldNo:"3" formLineNo:"3_(d)_Regulatory_assets_under_management"`
	RealEstateNAV     float64 `fieldNo:"3" formLineNo:"3_(d)_Net_assets_under_management"`
	SecuredAssetsGAV  float64 `fieldNo:"3" formLineNo:"3_(e)_Regulatory_assets_under_management"`
	SecuredAssetsNAV  float64 `fieldNo:"3" formLineNo:"3_(e)_Net_assets_under_management"`
	VentureCapitalGAV float64 `fieldNo:"3" formLineNo:"3_(f)_Regulatory_assets_under_management"`
	VentureCapitalNAV float64 `fieldNo:"3" formLineNo:"3_(f)_Net_assets_under_management"`
	OtherGAV          float64 `fieldNo:"3" formLineNo:"3_(g)_Regulatory_assets_under_management"`
	OtherNAV          float64 `fieldNo:"3" formLineNo:"3_(g)_Net_assets_under_management"`
	NOTAGAV           float64 `fieldNo:"3" formLineNo:"3_(h)_Regulatory_assets_under_management"`
	NOTANAV           float64 `fieldNo:"3" formLineNo:"3_(h)_Net_assets_under_management"`
}

func NewForm1bData() *Form1bData {
	// lastYear := time.Now().Year() - 1
	return &Form1bData{

		MasterFund:              false, // 6
		ParallelFundsAggregated: false, // 7
		HasDerivativesPositions: false, //13_b

	}
}

func NewForm1cData() *Form1cData {
	return &Form1cData{
		Strategy:                    "Single primary strategy",
		EquityLongBiasStrategy:      true,
		EquityLongBiasNAV:           "100",
		EquityMarketNeutralStrategy: false,
		EquityLongShortStrategy:     false,
		EquityShortBiasStrategy:     false,

		MacroActiveTradingStrategy:         false,
		MacroCommodityStrategy:             false,
		MacroCurrencyStrategy:              false,
		MacroGlobalMacroStrategy:           false,
		RVFixedIncomeAssetBackedStrategy:   false,
		RVFixedIncomeConvArbStrategy:       false,
		RVFixedIncomeCorporateStrategy:     false,
		RVFixedIncomeSovereignStrategy:     false,
		RVVolatilityArbStrategy:            false,
		EventDrivenDistressedStrategy:      false,
		EventDrivenRiskArbStrategy:         false,
		EventDrivenEqSpecialSitStrategy:    false,
		CreditLongShortStrategy:            false,
		CreditAssetBasedLendingStrategy:    false,
		ManagedFuturesFundamentalStrategy:  false,
		ManagedFuturesQuantitativeStrategy: false,
		InvestmentInOtherFundsStrategy:     false,
		OtherStrategy:                      false,

		HFTPercentage: "0%",

		PercentageCounterpartyALegalName:       "NA",
		PercentageCounterpartyAAffiliatedInsti: "NA",
		PercentageCounterpartyAExposure:        "NA",
		PercentageAInsticode:                   "32",
		PercentageCounterpartyBLegalName:       "NA",
		PercentageCounterpartyBAffiliatedInsti: "NA",
		PercentageCounterpartyBExposure:        "NA",
		PercentageBInsticode:                   "32",
		PercentageCounterpartyCLegalName:       "NA",
		PercentageCounterpartyCAffiliatedInsti: "NA",
		PercentageCounterpartyCExposure:        "NA",
		PercentageCInsticode:                   "32",
		PercentageCounterpartyDLegalName:       "NA",
		PercentageCounterpartyDAffiliatedInsti: "NA",
		PercentageCounterpartyDExposure:        "NA",
		PercentageDInsticode:                   "32",
		PercentageCounterpartyELegalName:       "NA",
		PercentageCounterpartyEAffiliatedInsti: "NA",
		PercentageCounterpartyEExposure:        "NA",
		PercentageEInsticode:                   "32",

		USDCounterpartyALegalName:       "NA",
		USDCounterpartyAAffiliatedInsti: "NA",
		USDCounterpartyAExposure:        "NA",
		USDAInsticode:                   "32",
		USDCounterpartyBLegalName:       "NA",
		USDCounterpartyBAffiliatedInsti: "NA",
		USDCounterpartyBExposure:        "NA",
		USDBInsticode:                   "32",
		USDCounterpartyCLegalName:       "NA",
		USDCounterpartyCAffiliatedInsti: "NA",
		USDCounterpartyCExposure:        "NA",
		USDCInsticode:                   "32",
		USDCounterpartyDLegalName:       "NA",
		USDCounterpartyDAffiliatedInsti: "NA",
		USDCounterpartyDExposure:        "NA",
		USDDInsticode:                   "32",
		USDCounterpartyELegalName:       "NA",
		USDCounterpartyEAffiliatedInsti: "NA",
		USDCounterpartyEExposure:        "NA",
		USDEInsticode:                   "32",

		SecRegExchange:        "NA", // 24_a_i
		DerivativesOTC:        "NA", // 24_b_ii
		DerivativesBilateraly: "NA", // 24_c_ii

		SecOTC:                 "NA", // 24_a_ii
		DerivativesRegExchange: "NA", // 24_b_i
		DerivativesCCP:         "NA", //24_c_i

		RepoCCP:        "NA", //24_d_i
		RepoBilateraly: "NA", // 24_d_ii
		RepoTriParty:   "NA", // 24_d_iii

		OtherTradePercent: 0,
	}
}

func NewForm1aData() *Form1aData {
	return &Form1aData{}
}
