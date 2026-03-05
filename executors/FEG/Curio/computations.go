package curio

import (
	"quantrium-tech/form_pf_extractor_service/executors"
)

// Compute act over a private fund to calculate all the required compuations and returns a 1bData since PV only uses it in computations
func (F *Fund) Compute() *executors.Form1bData {

	wd := F.workPaperData
	ld := F.workPaperData.liquidityData
	pd := F.workPaperData.fundPerformanceData

	data := executors.NewForm1bData()

	// workpaper data
	data.GrossAsset = executors.CalculateGrossAssets(wd.totalAssets, wd.unfundedCommitments)
	data.NetAsset = executors.CalculateNetAsset(wd.totalPartnerCaptial, wd.unfundedCommitments)
	data.EquityInOtherFunds = wd.sumLevelExtData
	data.Level1Asset = wd.l1
	data.Level2Asset = wd.l2
	data.Level3Asset = wd.l3
	data.CostBasedAsset = calculateCBAsset(data.GrossAsset, wd.sumLevelExtData)
	data.CostBasedLiability = wd.totalLiabilities

	// investor liquidity data
	data.UsPercent = ld.UsPersonOwnedPercentage
	data.OtherPercent = ld.OtherPlanOwnedPercentage
	data.PensionPlanPercent = ld.PensionPlanOwnedPercentage
	data.NonProfitPercent = ld.NonProfitOwnedPercentage
	data.Top5Percent = ld.Top5OwnedPercentage
	data.SovereignPercent = ld.SovereignFundPercentage
	data.PrivateFundPercent = ld.PvtFundPercentage
	assignFundPerf(data, pd)

	return data

}

func calculateCBAsset(GrossAsset, LevelingSum float64) float64 {
	return GrossAsset - LevelingSum
}
