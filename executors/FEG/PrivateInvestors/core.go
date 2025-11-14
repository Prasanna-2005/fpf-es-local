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

// GetFundID returns the fundID used internally
func (F *PrivateFund) GetFundID() string {
	return F.fundID
}

// SetLiquidityData takes up a reference of LiquidityData and set upon the fund
func (F *PrivateFund) SetLiquidityData(data *LiquidityData) {
	F.workPaperData.liquidityData = *data
}

// ExportWorkPaperData returns a mapped entityname to value mapper for further push
func (F *PrivateFund) ExportWorkPaperData() map[string]string {

	entityValueMapper := make(map[string]string)
	executors.CollectEntities(reflect.ValueOf(F.workPaperData), entityValueMapper)
	return entityValueMapper

}

// WorkPaperData represents the core structure that holds all the data for a particular Fund-Series combination
type WorkPaperData struct {
	totalAssets                   float64             `entity:"q3_total_assets_bs"`
	unfundedCommitments           float64             `entity:"q3_unfunded_commitments_captial_roll_qtd"`
	totalPartnerCaptial           float64             `entity:"q3_total_partner_captial_bs"`
	totalLiabilities              float64             `entity:"q3_total_liabilities_bs"`
	investments                   float64             `entity:"q3_investments_bs"`
	investmentsUnrealisedGainLoss float64             `entity:"q3_investments_unrealised_gain_loss_bs"`
	cashSweep                     float64             `entity:"q3_cash_sweep_bs"`
	cashGeneralAccount            float64             `entity:"q3_cash_general_account_bs"`
	cashChecking                  float64             `entity: "q3_cash_checking_bs"`
	liquidityData                 LiquidityData       `entity:"liquidity_data"`
	fundPerformanceData           FundPerformanceData `entity:"fund_performance_data"`
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
	q1GrossPercentage float64 `entity:"q1_gross_percentage_captial_roll_qtd"`
	q2GrossPercentage float64 `entity:"q2_gross_percentage_captial_roll_qtd"`
	q3GrossPercentage float64 `entity:"q3_gross_percentage_captial_roll_qtd"`
	q1NetPercentage   float64 `entity:"q1_net_percentage_captial_roll_qtd"`
	q2NetPercentage   float64 `entity:"q2_net_percentage_captial_roll_qtd"`
	q3NetPercentage   float64 `entity:"q3_net_percentage_captial_roll_qtd"`
}

// Holder holds the path of the files needed for a extraction of a single Private Fund
type Holder struct {
	q1FilePath string
	q2FilePath string
	q3FilePath string
}

// liquidityDataHolder represents the map that holds the liquidityData for the reference fund
type LiquidityDataHolder map[string]*LiquidityData
