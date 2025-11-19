package invesmentservices

import (
	"quantrium-tech/form_pf_extractor_service/config"
	"quantrium-tech/form_pf_extractor_service/executors"
	"reflect"
)

type Fund struct {
	fundID        string
	series        string
	workPaperData WorkPaperData
	holder        Holder
}

type Holder struct {
	tbFilePath          string
	reportFilePath      string
	capRegisterFilePath string
}

func NewFund(fundId string) *Fund {
	return &Fund{fundID: fundId}
}

func (F *Fund) SetHolder(FileType string, filePath string) {

	filePath = config.Cfg.StoragePath + filePath

	switch FileType {
	case "nav_tb":
		F.holder.tbFilePath = filePath
	case "nav_report":
		F.holder.reportFilePath = filePath
	case "capital_register":
		F.holder.capRegisterFilePath = filePath
	}

}

// GetFundID returns the fundID used internally
func (F *Fund) GetFundID() string {
	return F.fundID
}

// SetLiquidityData takes up a reference of LiquidityData and set upon the fund
func (F *Fund) SetLiquidityData(data *LiquidityData) {
	F.workPaperData.liquidityData = *data
}

// ExportWorkPaperData returns a mapped entityname to value mapper for further push
func (F *Fund) ExportWorkPaperData() map[string]string {

	entityValueMapper := make(map[string]string)
	executors.CollectEntities(reflect.ValueOf(F.workPaperData), entityValueMapper)
	return entityValueMapper

}

type WorkPaperData struct {
	totalAssets         float64 `entity:"tb_total_assets_bs"`
	unfundedCommitments float64 `entity:"tb_unfunded_commitments_captial_roll_qtd"`
	totalPartnerCaptial float64 `entity:"tb_total_partner_captial_bs"`
	totalLiabilities    float64 `entity:"tb_total_liabilities_bs"`
	// equityData          float64             `entity:"equity_otherfunds"`
	l1                  float64             `entity:"level1"`
	l2                  float64             `entity:"level2"`
	l3                  float64             `entity:"level3"`
	levelExtData        LevelingData        `entity:"level_ext_data"`
	sumLevelExtData     float64             `entity:"equity_otherfunds"`
	liquidityData       LiquidityData       `entity:"liquidity_data"`
	fundPerformanceData FundPerformanceData `entity:"fund_performance_data"`
}

type LevelingData struct {
	adrCost             float64 `entity:"adr_cost"`
	adrURGL             float64 `entity:"adr_ur_gl"`
	corpBondsCost       float64 `entity:"corp_bonds_cost"`
	corpBondsURGL       float64 `entity:"corp_bonds_ur_gl"`
	directHoldingCost   float64 `entity:"direct_holding_cost"`
	directHoldingURGL   float64 `entity:"direct_holding_ur_gl"`
	equitiesCost        float64 `entity:"equities_cost"`
	equitiesDividendRec float64 `entity:"equities_dividend_rec"`
	equitiesFtWhPayable float64 `entity:"equities_ft_wh_payable"`
	equitiesURFXGL      float64 `entity:"equities_ur_fx_gl"`
	equitiesURGL        float64 `entity:"equities_ur_gl"`
	gdrCost             float64 `entity:"gdr_cost"`
	gdrURGL             float64 `entity:"gdr_ur_gl"`
	govtBondsCost       float64 `entity:"govt_bonds_cost"`
	govtBondsURGL       float64 `entity:"govt_bonds_ur_gl"`
	invtHedgeFundsCost  float64 `entity:"invt_hedge_funds_cost"`
	invtInHedgeFundsUR  float64 `entity:"invt_in_hedge_funds_ur"`
	mortgageCmoCost     float64 `entity:"mortgage_cmo_cost"`
	mortgageCmoURGL     float64 `entity:"mortgage_cmo_ur_gl"`
	privateEqCost       float64 `entity:"private_eq_cost"`
	privateEqURGL       float64 `entity:"private_eq_ur_gl"`
	reitCost            float64 `entity:"reit_cost"`
	reitURGL            float64 `entity:"reit_ur_gl"`
	treasuriesCost      float64 `entity:"treasuries_cost"`
	treasuriesURGL      float64 `entity:"treasuries_ur_gl"`
	invtInMasterFund    float64 `entity:"invt_in_master_fund"`
}

// LiquidityData represents the structure that holds all the data retrieved from the investorLiquidity File
type LiquidityData struct {
	top5OwnedPercentage        float64 `entity:"top_5_owned_percentage"`
	nonProfitOwnedPercentage   float64 `entity:"non_profit_percentage"`
	pensionPlanOwnedPercentage float64 `entity:"pension_plan_percentage"`
	otherPlanOwnedPercentage   float64 `entity:"others_plan_percentage"`
	usPersonOwnedPercentage    float64 `entity:"us_person_percentage"`
	pvtFundPercentage          float64 `entity:"private_fund_percentage"`
	sovereignFundPercentage    float64 `entity:"sovereign_fund_percentage"`
}

type FundPerformanceData struct {
	q1GrossPercentage float64 `entity:"q1_gross_percentage_captial_roll_qtd"`
	q2GrossPercentage float64 `entity:"q2_gross_percentage_captial_roll_qtd"`
	q3GrossPercentage float64 `entity:"q3_gross_percentage_captial_roll_qtd"`
	q1NetPercentage   float64 `entity:"q1_net_percentage_captial_roll_qtd"`
	q2NetPercentage   float64 `entity:"q2_net_percentage_captial_roll_qtd"`
	q3NetPercentage   float64 `entity:"q3_net_percentage_captial_roll_qtd"`
}

// liquidityDataHolder represents the map that holds the liquidityData for the reference fund
type LiquidityDataHolder map[string]*LiquidityData
