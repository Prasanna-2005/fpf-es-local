package invesmentservices

import "quantrium-tech/form_pf_extractor_service/executors"

// Compute act over a private fund to calculate all the required compuations and returns a 1bData since PV only uses it in computations
func (F *Fund) Compute() *executors.Form1bData {

	wd := F.workPaperData
	ld := F.workPaperData.liquidityData
	// pd := F.workPaperData.fundPerformanceData

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
	data.UsPercent = ld.usPersonOwnedPercentage
	data.OtherPercent = ld.otherPlanOwnedPercentage
	data.PensionPlanPercent = ld.pensionPlanOwnedPercentage
	data.NonProfitPercent = ld.nonProfitOwnedPercentage
	data.Top5Percent = ld.top5OwnedPercentage
	data.SovereignPercent = ld.sovereignFundPercentage
	data.PrivateFundPercent = ld.pvtFundPercentage

	// // fund performance data
	// data.Q1GrossValue = pd.q1GrossPercentage
	// data.Q2GrossValue = pd.q2GrossPercentage
	// data.Q3GrossValue = pd.q3GrossPercentage
	// data.Q1NetValue = pd.q1NetPercentage
	// data.Q2NetValue = pd.q2NetPercentage
	// data.Q3NetValue = pd.q3NetPercentage

	return data

}

func calculateCBAsset(GrossAsset, LevelingSum float64) float64 {
	return GrossAsset - LevelingSum
}
