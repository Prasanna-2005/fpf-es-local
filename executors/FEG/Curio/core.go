package curio

import (
	"quantrium-tech/form_pf_extractor_service/config"
	"quantrium-tech/form_pf_extractor_service/executors"
	"reflect"
)

type UnfundedCommitmentsHolder map[string]float64

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
	l1              float64      `entity:"level1"`
	l2              float64      `entity:"level2"`
	l3              float64      `entity:"level3"`
	levelExtData    LevelingData `entity:"level_ext_data"`
	sumLevelExtData float64      `entity:"equity_otherfunds"`
	// liquidityData       LiquidityData       `entity:"liquidity_data"`
	fundPerformanceData FundPerformanceData `entity:"fund_performance_data"`
}

func NewFund(fundId string) *Fund {
	return &Fund{fundID: fundId, series: "Curio"}
}

type LevelingData struct {
	invtHedgeFundsCost float64 `entity:"invt_hedge_funds_cost"`
	invtInHedgeFundsUR float64 `entity:"invt_in_hedge_funds_ur"`
	privateEqCost      float64 `entity:"private_eq_cost"`
	privateEqURGL      float64 `entity:"private_eq_ur_gl"`
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

// GetFundID returns the fundID used internally
func (F *Fund) GetFundID() string {
	return F.fundID
}

// GetFundType returns the fundID used internally
func (F *Fund) GetFundType() string {
	return F.fundType
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
