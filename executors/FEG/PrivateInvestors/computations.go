package privateinvestors

import "quantrium-tech/form_pf_extractor_service/executors"

// POF Specific Formulae

// calculateLevel3Asset calculates the Level3 Asset for PV based on backtracked formula
func calculateLevel3Asset(investment1900 float64, unrealisedGainLoss1910 float64) float64 {
	return investment1900 + unrealisedGainLoss1910
}

// below one is Depriciated formula to calculate CB assets . This Version 1 formula
// calculateCBAsset calculates the CB Asset for PV based on backtracked formula
// func calculateCBAsset(cashChecking1000 float64, cashSweep1002 float64, cashGeneral1005 float64) float64 {
// 	return cashChecking1000 + cashSweep1002 + cashGeneral1005
// }

func calculateCBAsset(GAV, UnfundedCommitments, FairValueAssets float64) float64 {
	return GAV - UnfundedCommitments - FairValueAssets
}

// Compute act over a private fund to calculate all the required compuations and returns a 1bData since PV only uses it in computations
func (F *PrivateFund) Compute() *executors.Form1bData {

	wd := F.workPaperData
	ld := F.workPaperData.liquidityData
	pd := F.workPaperData.fundPerformanceData

	data := executors.NewForm1bData()

	// workpaper data
	data.GrossAsset = executors.CalculateGrossAssets(wd.totalAssets, wd.unfundedCommitments)
	data.NetAsset = executors.CalculateNetAsset(wd.totalPartnerCaptial, wd.unfundedCommitments)
	data.EquityInOtherFunds = calculateLevel3Asset(wd.investments, wd.investmentsUnrealisedGainLoss)
	data.Level3Asset = data.EquityInOtherFunds
	data.CostBasedAsset = calculateCBAsset(data.GrossAsset, wd.unfundedCommitments, data.Level3Asset)
	data.CostBasedLiability = executors.CalcCBLiabilities(data.GrossAsset, data.NetAsset, 0)

	// investor liquidity data
	data.UsPercent = ld.UsPersonOwnedPercentage
	data.OtherPercent = ld.OtherPlanOwnedPercentage
	data.PensionPlanPercent = ld.PensionPlanOwnedPercentage
	data.NonProfitPercent = ld.NonProfitOwnedPercentage
	data.Top5Percent = ld.Top5OwnedPercentage

	// fund performance data
	assignFundPerf(data, pd)

	return data

}
