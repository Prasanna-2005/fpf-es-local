package invesmentservices

import (
	"quantrium-tech/form_pf_extractor_service/config"
	"quantrium-tech/form_pf_extractor_service/executors"
	"reflect"
)

type Fund struct {
	fundID        string
	series        string
	fundType      string
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

// GetFundType returns the fundID used internally
func (F *Fund) GetFundType() string {
	return F.fundType
}

// SetLiquidityData takes up a reference of LiquidityData and set upon the fund
func (F *Fund) SetLiquidityData(data *LiquidityData) {
	F.workPaperData.liquidityData = *data
}

func (F *Fund) SetUnfdCmtsData(data float64) {
	F.workPaperData.unfundedCommitments = data
}

// ExportWorkPaperData returns a mapped entityname to value mapper for further push
func (F *Fund) ExportWorkPaperData() map[string]string {
	entityValueMapper := make(map[string]string)
	executors.CollectEntities(reflect.ValueOf(F.workPaperData), entityValueMapper)
	return entityValueMapper
}

type WorkPaperData struct {
	totalAssets         float64 `entity:"tb_total_assets_bs"`
	unfundedCommitments float64 `entity:"tb_unfunded_commitments_nav_report"`
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
	yearGrossPercentage float64 `entity:"year_gross_percentage_nav_report" term:"year"   colName:"gross"`
	yearNetPercentage   float64 `entity:"year_net_percentage_nav_report"   term:"year"   colName:"net"`

	q1GrossPercentage float64 `entity:"q1_gross_percentage_nav_report" term:"first"  colName:"gross"`
	q2GrossPercentage float64 `entity:"q2_gross_percentage_nav_report" term:"second" colName:"gross"`
	q3GrossPercentage float64 `entity:"q3_gross_percentage_nav_report" term:"third"  colName:"gross"`
	q4GrossPercentage float64 `entity:"q4_gross_percentage_nav_report" term:"fourth" colName:"gross"`

	q1NetPercentage float64 `entity:"q1_net_percentage_nav_report" term:"first"  colName:"net"`
	q2NetPercentage float64 `entity:"q2_net_percentage_nav_report" term:"second" colName:"net"`
	q3NetPercentage float64 `entity:"q3_net_percentage_nav_report" term:"third"  colName:"net"`
	q4NetPercentage float64 `entity:"q4_net_percentage_nav_report" term:"fourth" colName:"net"`

	janGrossPercentage float64 `entity:"jan_gross_percentage_nav_report" term:"january"   colName:"gross"`
	febGrossPercentage float64 `entity:"feb_gross_percentage_nav_report" term:"february"  colName:"gross"`
	marGrossPercentage float64 `entity:"mar_gross_percentage_nav_report" term:"march"     colName:"gross"`
	aprGrossPercentage float64 `entity:"apr_gross_percentage_nav_report" term:"april"     colName:"gross"`
	mayGrossPercentage float64 `entity:"may_gross_percentage_nav_report" term:"may"       colName:"gross"`
	junGrossPercentage float64 `entity:"jun_gross_percentage_nav_report" term:"june"      colName:"gross"`
	julGrossPercentage float64 `entity:"jul_gross_percentage_nav_report" term:"july"      colName:"gross"`
	augGrossPercentage float64 `entity:"aug_gross_percentage_nav_report" term:"august"    colName:"gross"`
	sepGrossPercentage float64 `entity:"sep_gross_percentage_nav_report" term:"september" colName:"gross"`
	octGrossPercentage float64 `entity:"oct_gross_percentage_nav_report" term:"october"   colName:"gross"`
	novGrossPercentage float64 `entity:"nov_gross_percentage_nav_report" term:"november"  colName:"gross"`
	decGrossPercentage float64 `entity:"dec_gross_percentage_nav_report" term:"december"  colName:"gross"`

	janNetPercentage float64 `entity:"jan_net_percentage_nav_report" term:"january"   colName:"net"`
	febNetPercentage float64 `entity:"feb_net_percentage_nav_report" term:"february"  colName:"net"`
	marNetPercentage float64 `entity:"mar_net_percentage_nav_report" term:"march"     colName:"net"`
	aprNetPercentage float64 `entity:"apr_net_percentage_nav_report" term:"april"     colName:"net"`
	mayNetPercentage float64 `entity:"may_net_percentage_nav_report" term:"may"       colName:"net"`
	junNetPercentage float64 `entity:"jun_net_percentage_nav_report" term:"june"      colName:"net"`
	julNetPercentage float64 `entity:"jul_net_percentage_nav_report" term:"july"      colName:"net"`
	augNetPercentage float64 `entity:"aug_net_percentage_nav_report" term:"august"    colName:"net"`
	sepNetPercentage float64 `entity:"sep_net_percentage_nav_report" term:"september" colName:"net"`
	octNetPercentage float64 `entity:"oct_net_percentage_nav_report" term:"october"   colName:"net"`
	novNetPercentage float64 `entity:"nov_net_percentage_nav_report" term:"november"  colName:"net"`
	decNetPercentage float64 `entity:"dec_net_percentage_nav_report" term:"december"  colName:"net"`
}

// liquidityDataHolder represents the map that holds the liquidityData for the reference fund
type LiquidityDataHolder map[string]*LiquidityData
type UnfundedCommitmentsHolder map[string]float64
