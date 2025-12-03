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
type Form1cData struct {
	FundName string `fieldNo:"18" formLineNo:"18_a"`
	FundID   string `fieldNo:"18" formLineNo:"18_b"`

	Strategy               string  `fieldNo:"19" formLineNo:"19"`
	EquityLongBiasStrategy bool    `fieldNo:"20" formLineNo:"20_(d)_Strategy"`
	EquityLongBiasNAV      float64 `fieldNo:"20" formLineNo:"20_(d)_%_of_NAV"`
	// EquityLongBiasCapital  string `fieldNo:"20" formLineNo:"20_d_%_of_capital"`

	HFTPercentage string `fieldNo:"21" formLineNo:"21"`
	// Row A
	PercentageCounterpartyALegalName       string `fieldNo:"22" formLineNo:"22_(a)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	PercentageCounterpartyAAffiliatedInsti string `fieldNo:"22" formLineNo:"22_(a)_Indicate below if the counterparty is affiliated with a major financial institution"`
	PercentageCounterpartyAExposure        string `fieldNo:"22" formLineNo:"22_(a)_Exposure (% of reporting fund's net asset value)"`

	// Row B
	PercentageCounterpartyBLegalName       string `fieldNo:"22" formLineNo:"22_(b)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	PercentageCounterpartyBAffiliatedInsti string `fieldNo:"22" formLineNo:"22_(b)_Indicate below if the counterparty is affiliated with a major financial institution"`
	PercentageCounterpartyBExposure        string `fieldNo:"22" formLineNo:"22_(b)_Exposure (% of reporting fund's net asset value)"`

	// Row C
	PercentageCounterpartyCLegalName       string `fieldNo:"22" formLineNo:"22_(c)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	PercentageCounterpartyCAffiliatedInsti string `fieldNo:"22" formLineNo:"22_(c)_Indicate below if the counterparty is affiliated with a major financial institution"`
	PercentageCounterpartyCExposure        string `fieldNo:"22" formLineNo:"22_(c)_Exposure (% of reporting fund's net asset value)"`

	// Row D
	PercentageCounterpartyDLegalName       string `fieldNo:"22" formLineNo:"22_(d)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	PercentageCounterpartyDAffiliatedInsti string `fieldNo:"22" formLineNo:"22_(d)_Indicate below if the counterparty is affiliated with a major financial institution"`
	PercentageCounterpartyDExposure        string `fieldNo:"22" formLineNo:"22_(d)_Exposure (% of reporting fund's net asset value)"`

	// Row E
	PercentageCounterpartyELegalName       string `fieldNo:"22" formLineNo:"22_(e)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	PercentageCounterpartyEAffiliatedInsti string `fieldNo:"22" formLineNo:"22_(e)_Indicate below if the counterparty is affiliated with a major financial institution"`
	PercentageCounterpartyEExposure        string `fieldNo:"22" formLineNo:"22_(e)_Exposure (% of reporting fund's net asset value)"`

	// Row A
	USDCounterpartyALegalName       string `fieldNo:"23" formLineNo:"23_(a)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	USDCounterpartyAAffiliatedInsti string `fieldNo:"23" formLineNo:"23_(a)_Indicate below if the counterparty is affiliated with a major financial institution"`
	USDCounterpartyAExposure        string `fieldNo:"23" formLineNo:"23_(a)_Exposure (in U.S. dollars))"`

	// Row B
	USDCounterpartyBLegalName       string `fieldNo:"23" formLineNo:"23_(b)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	USDCounterpartyBAffiliatedInsti string `fieldNo:"23" formLineNo:"23_(b)_Indicate below if the counterparty is affiliated with a major financial institution"`
	USDCounterpartyBExposure        string `fieldNo:"23" formLineNo:"23_(b)_Exposure (in U.S. dollars))"`

	// Row C
	USDCounterpartyCLegalName       string `fieldNo:"23" formLineNo:"23_(c)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	USDCounterpartyCAffiliatedInsti string `fieldNo:"23" formLineNo:"23_(c)_Indicate below if the counterparty is affiliated with a major financial institution"`
	USDCounterpartyCExposure        string `fieldNo:"23" formLineNo:"23_(c)_Exposure (in U.S. dollars))"`

	// Row D
	USDCounterpartyDLegalName       string `fieldNo:"23" formLineNo:"23_(d)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	USDCounterpartyDAffiliatedInsti string `fieldNo:"23" formLineNo:"23_(d)_Indicate below if the counterparty is affiliated with a major financial institution"`
	USDCounterpartyDExposure        string `fieldNo:"23" formLineNo:"23_(d)_Exposure (in U.S. dollars))"`

	// Row E
	USDCounterpartyELegalName       string `fieldNo:"23" formLineNo:"23_(e)_Legal name of the counterparty (or, if multiple affiliated entities, counterparties)"`
	USDCounterpartyEAffiliatedInsti string `fieldNo:"23" formLineNo:"23_(e)_Indicate below if the counterparty is affiliated with a major financial institution"`
	USDCounterpartyEExposure        string `fieldNo:"23" formLineNo:"23_(e)_Exposure (in U.S. dollars))"`

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

func NewForm1cData() *Form1cData {
	return &Form1cData{
		Strategy:               "Single primary strategy",
		EquityLongBiasStrategy: true,
		EquityLongBiasNAV:      100,
		HFTPercentage:          "0%",

		PercentageCounterpartyALegalName:       "NA",
		PercentageCounterpartyAAffiliatedInsti: "NA",
		PercentageCounterpartyAExposure:        "NA",
		PercentageCounterpartyBLegalName:       "NA",
		PercentageCounterpartyBAffiliatedInsti: "NA",
		PercentageCounterpartyBExposure:        "NA",
		PercentageCounterpartyCLegalName:       "NA",
		PercentageCounterpartyCAffiliatedInsti: "NA",
		PercentageCounterpartyCExposure:        "NA",
		PercentageCounterpartyDLegalName:       "NA",
		PercentageCounterpartyDAffiliatedInsti: "NA",
		PercentageCounterpartyDExposure:        "NA",
		PercentageCounterpartyELegalName:       "NA",
		PercentageCounterpartyEAffiliatedInsti: "NA",
		PercentageCounterpartyEExposure:        "NA",

		USDCounterpartyALegalName:       "NA",
		USDCounterpartyAAffiliatedInsti: "NA",
		USDCounterpartyAExposure:        "NA",
		USDCounterpartyBLegalName:       "NA",
		USDCounterpartyBAffiliatedInsti: "NA",
		USDCounterpartyBExposure:        "NA",
		USDCounterpartyCLegalName:       "NA",
		USDCounterpartyCAffiliatedInsti: "NA",
		USDCounterpartyCExposure:        "NA",
		USDCounterpartyDLegalName:       "NA",
		USDCounterpartyDAffiliatedInsti: "NA",
		USDCounterpartyDExposure:        "NA",
		USDCounterpartyELegalName:       "NA",
		USDCounterpartyEAffiliatedInsti: "NA",
		USDCounterpartyEExposure:        "NA",

		SecRegExchange:        "100", // 24_a_i
		DerivativesOTC:        "100", // 24_b_ii
		DerivativesBilateraly: "100", // 24_c_ii

		SecOTC:                 "0", // 24_a_ii
		DerivativesRegExchange: "0", // 24_b_i
		DerivativesCCP:         "0", //24_c_i

		RepoCCP:        "NA", //24_d_i
		RepoBilateraly: "NA", // 24_d_ii
		RepoTriParty:   "NA", // 24_d_iii

		OtherTradePercent: 0,
	}
}
