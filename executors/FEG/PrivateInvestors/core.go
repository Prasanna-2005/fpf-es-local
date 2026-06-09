package privateinvestors

import (
	"quantrium-tech/form_pf_extractor_service/config"
	"quantrium-tech/form_pf_extractor_service/executors"
	"reflect"
)

// PrivateFund represents a single fund in private investment series. This act as the parent for underlying structures.
type PrivateFund struct {
	fundID        string
	series        string
	workPaperData WorkPaperData
	holder        Holder
}

// NewPrivateFund creates a new PrivateFund with the given fundId and series
func NewPrivateFund(fundId string, series string) *PrivateFund {
	return &PrivateFund{fundID: fundId, series: series, workPaperData: WorkPaperData{fundPerformanceData: FundPerformanceData{}}}
}

// SetHolder takes up a quarterNo and filePath and sets them in the file holder of the fund
func (F *PrivateFund) SetHolder(quarterNo int, filePath string) {

	filePath = config.Cfg.StoragePath + filePath

	switch quarterNo {
	case 1:
		F.holder.q1FilePath = filePath
	case 2:
		F.holder.q2FilePath = filePath
	case 3:
		F.holder.q3FilePath = filePath
	}

}

func (F *PrivateFund) CopyHolder(partnerFund *PrivateFund) {
	partnerFund.holder.q1FilePath = F.holder.q1FilePath
	partnerFund.holder.q2FilePath = F.holder.q2FilePath
	partnerFund.holder.q3FilePath = F.holder.q3FilePath
}

// GetFundID returns the fundID used internally
func (F *PrivateFund) GetFundID() string {
	return F.fundID
}

// SetLiquidityData takes up a reference of LiquidityData and set upon the fund
func (F *PrivateFund) SetLiquidityData(data *executors.LiquidityData) {
	F.workPaperData.liquidityData = *data
}

func (F *PrivateFund) CombineData(partnerFund *PrivateFund) {
	F.workPaperData.totalAssets += partnerFund.workPaperData.totalAssets
	F.workPaperData.unfundedCommitments += partnerFund.workPaperData.unfundedCommitments
	F.workPaperData.totalPartnerCaptial += partnerFund.workPaperData.totalPartnerCaptial
	F.workPaperData.totalLiabilities += partnerFund.workPaperData.totalLiabilities
	F.workPaperData.investments += partnerFund.workPaperData.investments
	F.workPaperData.investmentsUnrealisedGainLoss += partnerFund.workPaperData.investmentsUnrealisedGainLoss
	F.workPaperData.cashSweep += partnerFund.workPaperData.cashSweep
	F.workPaperData.cashGeneralAccount += partnerFund.workPaperData.cashGeneralAccount
	F.workPaperData.cashChecking += partnerFund.workPaperData.cashChecking
	F.CombineFundPerformance(partnerFund)
}

func (F *PrivateFund) CombineFundPerformance(partnerFund *PrivateFund) {
	pd := F.workPaperData.fundPerformanceData
	ppd := partnerFund.workPaperData.fundPerformanceData

	// Q1
	pd.netIncomeq1 += ppd.netIncomeq1
	pd.mgmtFeeq1 += ppd.mgmtFeeq1
	pd.begBalq1 += ppd.begBalq1
	pd.capCallq1 += ppd.capCallq1
	pd.cashDistribq1 += ppd.cashDistribq1

	// Q2
	pd.netIncomeq2 += ppd.netIncomeq2
	pd.mgmtFeeq2 += ppd.mgmtFeeq2
	pd.begBalq2 += ppd.begBalq2
	pd.capCallq2 += ppd.capCallq2
	pd.cashDistribq2 += ppd.cashDistribq2

	// Q3
	pd.netIncomeq3 += ppd.netIncomeq3
	pd.mgmtFeeq3 += ppd.mgmtFeeq3
	pd.begBalq3 += ppd.begBalq3
	pd.capCallq3 += ppd.capCallq3
	pd.cashDistribq3 += ppd.cashDistribq3

	F.workPaperData.fundPerformanceData = pd
}

// ExportWorkPaperData returns a mapped entityname to value mapper for further push
func (F *PrivateFund) ExportWorkPaperData() map[string]string {

	entityValueMapper := make(map[string]string)
	executors.CollectEntities(reflect.ValueOf(F.workPaperData), entityValueMapper)
	return entityValueMapper

}

// WorkPaperData represents the core structure that holds all the data for a particular Fund-Series combination
type WorkPaperData struct {
	totalAssets                   float64                 `entity:"q3_total_assets_bs"`
	unfundedCommitments           float64                 `entity:"q3_unfunded_commitments_captial_roll_qtd"`
	totalPartnerCaptial           float64                 `entity:"q3_total_partner_captial_bs"`
	totalLiabilities              float64                 `entity:"q3_total_liabilities_bs"`
	investments                   float64                 `entity:"q3_investments_bs"`
	investmentsUnrealisedGainLoss float64                 `entity:"q3_investments_unrealised_gain_loss_bs"`
	cashSweep                     float64                 `entity:"q3_cash_sweep_bs"`
	cashGeneralAccount            float64                 `entity:"q3_cash_general_account_bs"`
	cashChecking                  float64                 `entity:"q3_cash_checking_bs"`
	commericaLineOfCreditPayable  float64                 `entity:"q3_commerica_credit_payable"`
	liquidityData                 executors.LiquidityData `entity:"liquidity_data"`
	fundPerformanceData           FundPerformanceData     `entity:"fund_performance_data"`
}

// LiquidityData represents the structure that holds all the data retrieved from the investorLiquidity File
type LiquidityData struct {
	top5OwnedPercentage        float64 `entity:"top_5_owned_percentage"`
	nonProfitOwnedPercentage   float64 `entity:"non_profit_percentage"`
	pensionPlanOwnedPercentage float64 `entity:"pension_plan_percentage"`
	otherPlanOwnedPercentage   float64 `entity:"others_plan_percentage"`
	usPersonOwnedPercentage    float64 `entity:"us_person_percentage"`
}

// FundPerformanceData represents the strucure that holds the fund performance on three quarters
type FundPerformanceData struct {
	q1GrossPercentage float64 `entity:"q1_gross_percentage_captial_roll_qtd" term:"first"  colName:"gross"`
	q2GrossPercentage float64 `entity:"q2_gross_percentage_captial_roll_qtd" term:"second"  colName:"gross"`
	q3GrossPercentage float64 `entity:"q3_gross_percentage_captial_roll_qtd" term:"third"  colName:"gross"`
	q1NetPercentage   float64 `entity:"q1_net_percentage_captial_roll_qtd" term:"first"  colName:"net"`
	q2NetPercentage   float64 `entity:"q2_net_percentage_captial_roll_qtd" term:"second"  colName:"net"`
	q3NetPercentage   float64 `entity:"q3_net_percentage_captial_roll_qtd" term:"third"  colName:"net"`

	netIncomeq3   float64 `entity:"net_income_q3" term:"third"  colName:"netIncome"`
	mgmtFeeq3     float64 `entity:"mgmt_fee_q3" term:"third"  colName:"mgmtFee"`
	begBalq3      float64 `entity:"begining_balance_q3" term:"third"  colName:"beginBal"`
	capCallq3     float64 `entity:"investor_capital_call_q3" term:"third"  colName:"capitalCall"`
	cashDistribq3 float64 `entity:"cash_distribution_q3" term:"third"  colName:"cashDistibution"`

	netIncomeq2   float64 `entity:"net_income_q2" term:"second"  colName:"netIncome"`
	mgmtFeeq2     float64 `entity:"mgmt_fee_q2" term:"second"  colName:"mgmtFee"`
	begBalq2      float64 `entity:"begining_balance_q2" term:"second"  colName:"beginBal"`
	capCallq2     float64 `entity:"investor_capital_call_q2" term:"second"  colName:"capitalCall"`
	cashDistribq2 float64 `entity:"cash_distribution_q2" term:"second"  colName:"cashDistibution"`

	netIncomeq1   float64 `entity:"net_income_q1" term:"first"  colName:"netIncome"`
	mgmtFeeq1     float64 `entity:"mgmt_fee_q1" term:"first"  colName:"mgmtFee"`
	begBalq1      float64 `entity:"begining_balance_q1" term:"first"  colName:"beginBal"`
	capCallq1     float64 `entity:"investor_capital_call_q1" term:"first"  colName:"capitalCall"`
	cashDistribq1 float64 `entity:"cash_distribution_q1" term:"first"  colName:"cashDistibution"`
}
type QTDval struct {
	Net         float64 `entity:"net_percent"`
	Gross       float64 `entity:"gross_percent"`
	BegBal      float64 `entity:"beginning_balance"`
	CashDistrib float64 `entity:"cash_distribution"`
	CapCall     float64 `entity:"capital_call"`
	NetIncome   float64 `entity:"net_income"`
	MgmtFee     float64 `entity:"management_fee"`
}

// Holder holds the path of the files needed for a extraction of a single Private Fund
type Holder struct {
	q1FilePath string
	q2FilePath string
	q3FilePath string
}

// liquidityDataHolder represents the map that holds the liquidityData for the reference fund
type LiquidityDataHolder map[string]*LiquidityData
